package provider

import (
	"context"
	"time"

	"github.com/boz/go-lifecycle"
	"github.com/pkg/errors"
	tpubsub "github.com/troian/pubsub"

	provider "github.com/akash-network/akash-api/go/provider/v1"

	"github.com/akash-network/node/pubsub"

	"github.com/akash-network/provider/bidengine"
	"github.com/akash-network/provider/cluster"
	ctypes "github.com/akash-network/provider/cluster/types/v1beta3"
	"github.com/akash-network/provider/manifest"
	"github.com/akash-network/provider/operator/waiter"
	"github.com/akash-network/provider/session"
	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/entities"
	"github.com/akash-network/provider/tools/fromctx"
	ptypes "github.com/akash-network/provider/types"
)

// StatusClient is the interface which includes status of service
//
//go:generate mockery --name StatusClient
type StatusClient interface {
	Status(context.Context) (*Status, error)
	StatusV1(ctx context.Context) (*provider.Status, error)
}

//go:generate mockery --name Client
type Client interface {
	StatusClient
	Manifest() manifest.Client
	Cluster() cluster.Client
	Hostname() ctypes.HostnameServiceClient
	ClusterService() cluster.Service
	Bus() pubsub.Bus
}

// Service is the interface that includes StatusClient interface.
// It also wraps ManifestHandler, Close and Done methods.
type Service interface {
	Client

	Close() error
	Done() <-chan struct{}
}

// NewService creates and returns new Service instance
// Simple wrapper around various services needed for running a provider.
func NewService(ctx context.Context,
	spClient *spheron.Client,
	accAddr string,
	session session.Session,
	bus pubsub.Bus,
	cclient cluster.Client,
	waiter waiter.OperatorWaiter,
	cfg Config) (Service, error) {
	ctx, cancel := context.WithCancel(ctx)

	session = session.ForModule("provider-service")

	clusterConfig := cluster.NewDefaultConfig()
	clusterConfig.InventoryResourcePollPeriod = cfg.InventoryResourcePollPeriod
	clusterConfig.InventoryResourceDebugFrequency = cfg.InventoryResourceDebugFrequency
	clusterConfig.InventoryExternalPortQuantity = cfg.ClusterExternalPortQuantity
	clusterConfig.CPUCommitLevel = cfg.CPUCommitLevel
	clusterConfig.MemoryCommitLevel = cfg.MemoryCommitLevel
	clusterConfig.StorageCommitLevel = cfg.StorageCommitLevel
	clusterConfig.BlockedHostnames = cfg.BlockedHostnames
	clusterConfig.DeploymentIngressStaticHosts = cfg.DeploymentIngressStaticHosts
	clusterConfig.DeploymentIngressDomain = cfg.DeploymentIngressDomain
	clusterConfig.ClusterSettings = cfg.ClusterSettings

	cluster, err := cluster.NewService(ctx, session, bus, cclient, spClient, waiter, clusterConfig)
	if err != nil {
		cancel()
		return nil, err
	}

	bidengine, err := bidengine.NewService(ctx, session, spClient, cluster, bus, waiter, bidengine.Config{
		PricingStrategy: cfg.BidPricingStrategy,
		Deposit:         cfg.BidDeposit,
		BidTimeout:      cfg.BidTimeout,
		Attributes:      cfg.Attributes,
		MaxGroupVolumes: cfg.MaxGroupVolumes,
	})
	if err != nil {
		errmsg := "creating bidengine service"
		session.Log().Error(errmsg, "err", err)
		cancel()
		<-cluster.Done()
		return nil, errors.Wrap(err, errmsg)
	}

	manifestConfig := manifest.ServiceConfig{
		HTTPServicesRequireAtLeastOneHost: !cfg.DeploymentIngressStaticHosts,
		ManifestTimeout:                   cfg.ManifestTimeout,
		RPCQueryTimeout:                   cfg.RPCQueryTimeout,
		CachedResultMaxAge:                cfg.CachedResultMaxAge,
	}

	manifest, err := manifest.NewService(ctx, session, bus, cluster.HostnameService(), spClient, manifestConfig)
	if err != nil {
		session.Log().Error("creating manifest handler", "err", err)
		cancel()
		<-cluster.Done()
		<-bidengine.Done()
		return nil, err
	}

	svc := &service{
		session:   session,
		bus:       bus,
		cluster:   cluster,
		cclient:   cclient,
		bidengine: bidengine,
		manifest:  manifest,
		ctx:       ctx,
		cancel:    cancel,
		lc:        lifecycle.New(),
		config:    cfg,
	}

	go svc.lc.WatchContext(ctx)
	go svc.run()
	go svc.statusRun()

	return svc, nil
}

type service struct {
	config  Config
	session session.Session
	bus     pubsub.Bus
	cclient cluster.Client

	cluster   cluster.Service
	bidengine bidengine.Service
	manifest  manifest.Service

	ctx    context.Context
	cancel context.CancelFunc
	lc     lifecycle.Lifecycle
}

func (s *service) Hostname() ctypes.HostnameServiceClient {
	return s.cluster.HostnameService()
}

func (s *service) ClusterService() cluster.Service {
	return s.cluster
}

func (s *service) Close() error {
	s.lc.Shutdown(nil)
	return s.lc.Error()
}

func (s *service) Done() <-chan struct{} {
	return s.lc.Done()
}

func (s *service) Manifest() manifest.Client {
	return s.manifest
}

func (s *service) Cluster() cluster.Client {
	return s.cclient
}

func (s *service) Bus() pubsub.Bus {
	return s.bus
}

func (s *service) Status(ctx context.Context) (*Status, error) {
	cluster, err := s.cluster.Status(ctx)
	if err != nil {
		return nil, err
	}
	bidengine, err := s.bidengine.Status(ctx)
	if err != nil {
		return nil, err
	}
	manifest, err := s.manifest.Status(ctx)
	if err != nil {
		return nil, err
	}
	return &Status{
		Cluster:               cluster,
		Bidengine:             bidengine,
		Manifest:              manifest,
		ClusterPublicHostname: s.config.ClusterPublicHostname,
	}, nil
}

func (s *service) StatusV1(ctx context.Context) (*provider.Status, error) {
	cluster, err := s.cluster.StatusV1(ctx)
	if err != nil {
		return nil, err
	}
	bidengine, err := s.bidengine.StatusV1(ctx)
	if err != nil {
		return nil, err
	}
	manifest, err := s.manifest.StatusV1(ctx)
	if err != nil {
		return nil, err
	}
	return &provider.Status{
		Cluster:   cluster,
		BidEngine: bidengine,
		Manifest:  manifest,
		PublicHostnames: []string{
			s.config.ClusterPublicHostname,
		},
		Timestamp: time.Now().UTC(),
	}, nil
}

func (s *service) run() {
	defer s.lc.ShutdownCompleted()

	// Wait for any service to finish
	select {
	case <-s.lc.ShutdownRequest():
	case <-s.cluster.Done():
	case <-s.bidengine.Done():
	case <-s.manifest.Done():
	}

	// Shut down all services
	s.lc.ShutdownInitiated(nil)
	s.cancel()

	// Wait for all services to finish
	<-s.cluster.Done()
	<-s.bidengine.Done()
	<-s.manifest.Done()
	// <-s.bc.lc.Done()

	s.session.Log().Info("shutdown complete")
}

func (s *service) statusRun() {
	bus := fromctx.MustPubSubFromCtx(s.ctx)

	events := bus.Sub(
		ptypes.PubSubTopicClusterStatus,
		ptypes.PubSubTopicBidengineStatus,
		ptypes.PubSubTopicManifestStatus,
	)

	defer bus.Unsub(events)

	status := provider.Status{
		PublicHostnames: []string{
			s.config.ClusterPublicHostname,
		},
	}

loop:
	for {
		select {
		case <-s.cluster.Done():
			return
		case evt := <-events:
			switch obj := evt.(type) {
			case provider.ClusterStatus:
				status.Timestamp = time.Now().UTC()
				status.Cluster = &obj
			case provider.BidEngineStatus:
				status.Timestamp = time.Now().UTC()
				status.BidEngine = &obj
			case provider.ManifestStatus:
				status.Timestamp = time.Now().UTC()
				status.Manifest = &obj
			default:
				continue loop
			}

			bus.Pub(status, []string{ptypes.PubSubTopicProviderStatus}, tpubsub.WithRetain())
		}
	}
}

type reservation struct {
	resources         entities.ResourceGroup
	adjustedResources entities.ResourceUnits
	clusterParams     interface{}
}

var _ ctypes.ReservationGroup = (*reservation)(nil)

func (r *reservation) Resources() entities.ResourceGroup {
	return r.resources
}

func (r *reservation) SetAllocatedResources(val entities.ResourceUnits) {
	r.adjustedResources = val
}

func (r *reservation) GetAllocatedResources() entities.ResourceUnits {
	return r.adjustedResources
}

func (r *reservation) SetClusterParams(val interface{}) {
	r.clusterParams = val
}

func (r *reservation) ClusterParams() interface{} {
	return r.clusterParams
}
