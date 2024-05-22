package bidengine

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/boz/go-lifecycle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/akash-network/node/pubsub"
	metricsutils "github.com/akash-network/node/util/metrics"
	"github.com/akash-network/node/util/runner"

	"github.com/akash-network/provider/cluster"
	ctypes "github.com/akash-network/provider/cluster/types/v1beta3"
	"github.com/akash-network/provider/event"
	"github.com/akash-network/provider/session"
	"github.com/akash-network/provider/spheron"
	"github.com/akash-network/provider/spheron/entities"
)

// order manages bidding and general lifecycle handling of an order.
type order struct {
	orderID mtypes.OrderID
	cfg     Config

	session                    session.Session
	cluster                    cluster.Cluster
	bus                        pubsub.Bus
	sub                        pubsub.Subscriber
	reservationFulfilledNotify chan<- int

	log      log.Logger
	lc       lifecycle.Lifecycle
	pass     ProviderAttrSignatureService
	spClient *spheron.Client
}

var (
	pricingDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:        "provider_bid_pricing_duration",
		Help:        "",
		ConstLabels: nil,
		Buckets:     prometheus.ExponentialBuckets(150000.0, 2.0, 10.0),
	})

	bidCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "provider_bid",
		Help: "The total number of bids created",
	}, []string{"action", "result"})

	reservationDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:        "provider_reservation_duration",
		Help:        "",
		ConstLabels: nil,
		Buckets:     prometheus.ExponentialBuckets(150000.0, 2.0, 10.0),
	})

	reservationCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "provider_reservation",
		Help: "",
	}, []string{"action", "result"})

	shouldBidCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "provider_should_bid",
		Help: "",
	}, []string{"result"})

	orderCompleteCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "provider_order_complete",
		Help: "",
	}, []string{"result"})
)

func newOrder(svc *service, oid mtypes.OrderID, cfg Config, pass ProviderAttrSignatureService, checkForExistingBid bool, spClient *spheron.Client) (*order, error) {
	return newOrderInternal(svc, oid, cfg, pass, checkForExistingBid, nil, spClient)
}

func newOrderInternal(svc *service, oid mtypes.OrderID, cfg Config, pass ProviderAttrSignatureService, checkForExistingBid bool, reservationFulfilledNotify chan<- int, spClient *spheron.Client) (*order, error) {
	// Create a subscription that will see all events that have not been read from e.sub.Events()
	sub, err := svc.sub.Clone()
	if err != nil {
		return nil, err
	}

	session := svc.session.ForModule("bidengine-order")

	log := session.Log().With("order", oid)

	order := &order{
		cfg:                        cfg,
		orderID:                    oid,
		session:                    session,
		cluster:                    svc.cluster,
		bus:                        svc.bus,
		sub:                        sub,
		log:                        log,
		lc:                         lifecycle.New(),
		reservationFulfilledNotify: reservationFulfilledNotify, // Normally nil in production
		pass:                       pass,
		spClient:                   spClient,
	}

	// Shut down when parent begins shutting down
	go order.lc.WatchChannel(svc.lc.ShuttingDown())

	// Run main loop in separate thread.
	go order.run(checkForExistingBid)

	// Notify parent of completion (allows drain).
	go func() {
		<-order.lc.Done()
		svc.drainch <- order
	}()

	return order, nil
}

var matchBidNotFound = regexp.MustCompile("^.+bid not found.+$")

func (o *order) bidTimeoutEnabled() bool {
	return o.cfg.BidTimeout > time.Duration(0)
}

func (o *order) getBidTimeout() <-chan time.Time {
	if o.bidTimeoutEnabled() {
		return time.After(o.cfg.BidTimeout)
	}

	return nil
}

func (o *order) isStaleBid(bid mtypes.Bid) bool {
	if !o.bidTimeoutEnabled() {
		return false
	}

	// This bid could be very old, compute the minimum age of the bid
	// do not try anything clever here like asking the RPC node for the current height
	// just use the height from when the session is created
	createdAtBlock := bid.GetCreatedAt()
	blockAge := createdAtBlock - o.session.CreatedAtBlockHeight()
	const minTimePerBlock = 5 * time.Second
	atLeastThisOld := time.Duration(blockAge) * minTimePerBlock
	return atLeastThisOld > o.cfg.BidTimeout
}

func (o *order) run(checkForExistingBid bool) {
	defer o.lc.ShutdownCompleted()
	ctx, cancel := context.WithCancel(context.Background())

	var (
		// channels for async operations.
		groupch       <-chan runner.Result
		storedGroupCh <-chan runner.Result
		clusterch     <-chan runner.Result
		bidch         <-chan runner.Result
		pricech       <-chan runner.Result
		queryBidCh    <-chan runner.Result
		shouldBidCh   <-chan runner.Result
		bidTimeout    <-chan time.Time

		group       *dtypes.Group
		reservation ctypes.Reservation

		won bool
		// msg *mtypes.MsgCreateBid
	)

	groupch = runner.Do(func() runner.Result {
		res, err := o.spClient.GetGroup(ctx, o.orderID.GroupID().DSeq)
		return runner.NewResult(res, err)
	})

	// Load existing bid if needed
	if checkForExistingBid {
		queryBidCh = runner.Do(func() runner.Result {
			res, err := o.spClient.GetBid(ctx, o.orderID.GroupID().DSeq)
			return runner.NewResult(res, err)
		})
		// Hide the group details result for later
		storedGroupCh = groupch
		groupch = nil
	}

	bidPlaced := false
loop:
	for {
		select {
		case <-o.lc.ShutdownRequest():
			break loop

		case queryBid := <-queryBidCh:

			o.log.Info("queryBid", "group", queryBid)

			err := queryBid.Error()
			bidFound := true
			if err != nil {
				// Use super-advanced technique for detecting if bid is not on blockchain
				if matchBidNotFound.MatchString(err.Error()) {
					bidFound = false
				} else {
					o.session.Log().Error("could not get existing bid", "err", err, "errtype", fmt.Sprintf("%T", err))
					break loop
				}
			}

			if bidFound {
				o.session.Log().Info("found existing bid")
				bidResponse := queryBid.Value().(*mtypes.QueryBidResponse)
				bid := bidResponse.GetBid()
				bidState := bid.GetState()
				if bidState != mtypes.BidOpen {
					o.session.Log().Error("bid in unexpected state", "bid-state", bidState)
					break loop
				}
				bidPlaced = true

				if o.isStaleBid(bid) {
					o.session.Log().Info("found expired bid", "block-height", bid.GetCreatedAt())
					break loop
				}

				bidTimeout = o.getBidTimeout()
			}
			groupch = storedGroupCh // Allow getting the group details result now
			storedGroupCh = nil

		case ev := <-o.sub.Events():
			o.log.Info("ORDER", "group", ev)

			switch ev := ev.(type) {
			case mtypes.EventLeaseCreated:

				// different group
				if !o.orderID.GroupID().Equals(ev.ID.GroupID()) {
					o.log.Debug("ignoring group", "group", ev.ID.GroupID())
					break
				}

				// check winning provider
				if ev.ID.Provider != o.spClient.Context.Key.Address.Hex() {
					orderCompleteCounter.WithLabelValues("lease-lost").Inc()
					o.log.Info("lease lost", "lease", ev.ID)
					bidPlaced = false // Lease lost, network closes bid
					break loop
				}

				// TODO(spheron): compare event.ID.Provider to o.session.Provider.Address so that provider knows if he lost the bid
				orderCompleteCounter.WithLabelValues("lease-won").Inc()

				// TODO: sanity check (price, state, etc...)
				o.log.Info("lease won", "lease", ev.ID)

				if err := o.bus.Publish(event.LeaseWon{
					LeaseID: ev.ID,
					Group:   group,
				}); err != nil {
					o.log.Error("failed to publish to event queue", err)
				}
				won = true

				break loop

			case mtypes.EventOrderClosed:
				// different deployment
				if !ev.ID.Equals(o.orderID) {
					break
				}

				o.log.Info("order closed")
				orderCompleteCounter.WithLabelValues("order-closed").Inc()
				break loop

			case mtypes.EventBidClosed:
				if won {
					// Ignore any event after LeaseCreated
					continue
				}

				// Ignore bid closed not for this group
				if !o.orderID.GroupID().Equals(ev.ID.GroupID()) {
					break
				}

				// Ignore bid closed not for this provider
				if ev.ID.GetProvider() != o.session.Provider().String() {
					break
				}

				// Bid has been closed (possibly by someone manually closing it on the CLI)
				bidPlaced = false // bid already not on the blockchain
				orderCompleteCounter.WithLabelValues("bid-closed-external").Inc()
				break loop
			}

		case result := <-groupch:
			// Group details fetched.

			groupch = nil
			o.log.Info("group fetched")

			if result.Error() != nil {
				o.log.Error("fetching group", "err", result.Error())
				break loop
			}

			res := result.Value().(dtypes.Group)
			group = &res
			o.log.Info("GROUP", "group", group)

			shouldBidCh = runner.Do(func() runner.Result {
				return runner.NewResult(o.shouldBid(group))
			})

		case result := <-shouldBidCh:
			shouldBidCh = nil

			if result.Error() != nil {
				shouldBidCounter.WithLabelValues(metricsutils.FailLabel).Inc()
				o.log.Error("failure during checking should bid", "err", result.Error())
				break loop
			}

			shouldBid := result.Value().(bool)
			if !shouldBid {
				shouldBidCounter.WithLabelValues("decline").Inc()
				o.log.Debug("declined to bid")
				break loop
			}

			shouldBidCounter.WithLabelValues("accept").Inc()
			o.log.Info("requesting reservation")
			// Begin reserving resources from cluster.
			clusterch = runner.Do(metricsutils.ObserveRunner(func() runner.Result {
				return runner.NewResult(o.cluster.Reserve(o.orderID, group))
			}, reservationDuration))

		case result := <-clusterch:
			clusterch = nil

			if result.Error() != nil {
				reservationCounter.WithLabelValues(metricsutils.OpenLabel, metricsutils.FailLabel)
				o.log.Error("reserving resources", "err", result.Error())
				break loop
			}

			reservationCounter.WithLabelValues(metricsutils.OpenLabel, metricsutils.SuccessLabel)

			o.log.Info("Reservation fulfilled")

			// If the channel is assigned and there is capacity, write into the channel
			if o.reservationFulfilledNotify != nil {
				select {
				case o.reservationFulfilledNotify <- 0:
				default:
				}
			}

			// Resources reserved
			reservation = result.Value().(ctypes.Reservation)
			if bidPlaced {
				o.log.Info("Fulfillment already exists")
				// fulfillment already created (state recovered via queryExistingOrders)
				break
			}
			pricech = runner.Do(metricsutils.ObserveRunner(func() runner.Result {
				// Calculate price & bid
				priceReq := Request{
					Owner:          group.GroupID.Owner,
					GSpec:          &group.GroupSpec,
					PricePrecision: DefaultPricePrecision,
				}
				return runner.NewResult(o.cfg.PricingStrategy.CalculatePrice(ctx, priceReq))
			}, pricingDuration))
		case result := <-pricech:
			pricech = nil
			if result.Error() != nil {
				o.log.Error("error calculating price", "err", result.Error())
				break loop
			}

			price := result.Value().(sdk.DecCoin)
			maxPrice := group.GroupSpec.Price()

			if maxPrice.GetDenom() != price.GetDenom() {
				o.log.Error("Unsupported Denomination", "calculated", price.String(), "max-price", maxPrice.String())
				break loop
			}

			if maxPrice.IsLT(price) {
				o.log.Info("Price too high, not bidding", "price", price.String(), "max-price", maxPrice.String())
				break loop
			}

			o.log.Debug("submitting fulfillment", "price", price)

			// TODO(spheron): make provider address dynamic
			msg := entities.Bid{
				OrderID:  o.orderID.DSeq,
				Bidder:   o.spClient.Context.Key.Address.Hex(),
				BidPrice: 1, // price.Amount.BigInt().Uint64(),
			}

			bidch = runner.Do(func() runner.Result {
				return runner.NewResult(o.spClient.BcClient.CreateBid(ctx, &msg))
			})

		case result := <-bidch:
			bidch = nil
			if result.Error() != nil {
				bidCounter.WithLabelValues(metricsutils.OpenLabel, metricsutils.FailLabel).Inc()
				o.log.Error("bid failed", "err", result.Error())
				break loop
			}

			o.log.Info("bid complete")
			bidCounter.WithLabelValues(metricsutils.OpenLabel, metricsutils.SuccessLabel).Inc()

			// Fulfillment placed.
			bidPlaced = true

			bidTimeout = o.getBidTimeout()
		case <-bidTimeout:
			// The bid was not acted upon (e.g. lease created or deployment closed) so close it now
			o.log.Info("bid timeout, closing bid")
			orderCompleteCounter.WithLabelValues("bid-timeout").Inc()
			break loop
		}
	}

	o.log.Info("shutting down")
	o.lc.ShutdownInitiated(nil)
	o.sub.Close()

	// cancel reservation
	if !won {
		if clusterch != nil {
			result := <-clusterch
			clusterch = nil
			if result.Error() == nil {
				reservation = result.Value().(ctypes.Reservation)
			}
		}
		if reservation != nil {
			o.log.Debug("unreserving reservation")
			if err := o.cluster.Unreserve(reservation.OrderID()); err != nil {
				o.log.Error("error unreserving reservation", "err", err)
				reservationCounter.WithLabelValues("close", metricsutils.FailLabel)
			} else {
				reservationCounter.WithLabelValues("close", metricsutils.SuccessLabel)
			}
		}
	}
	cancel()

	// Wait for all runners to complete.
	if groupch != nil {
		<-groupch
	}
	if clusterch != nil {
		<-clusterch
	}
	if bidch != nil {
		<-bidch
	}
	if pricech != nil {
		<-pricech
	}
}

func (o *order) shouldBid(group *dtypes.Group) (bool, error) {
	// does provider have required attributes?
	if !group.GroupSpec.MatchAttributes(o.session.Provider().Attributes) {
		o.log.Debug("unable to fulfill: incompatible provider attributes")
		return false, nil
	}

	// does order have required attributes?
	if !o.cfg.Attributes.SubsetOf(group.GroupSpec.Requirements.Attributes) {
		o.log.Debug("unable to fulfill: incompatible order attributes")
		return false, nil
	}

	if !spheron.ArrayContainsString(o.session.AcceptedTokens(), group.GetGroupSpec().Price().Denom) {
		o.log.Debug("unable to fulfill: incompatible order token")
		return false, nil
	}

	attr, err := o.pass.GetAttributes()
	if err != nil {
		return false, err
	}

	// does provider have required capabilities?
	if !group.GroupSpec.MatchResourcesRequirements(attr) {
		o.log.Debug("unable to fulfill: incompatible attributes for resources requirements", "wanted", group.GroupSpec, "have", attr)
		return false, nil
	}

	for _, resources := range group.GroupSpec.GetResourceUnits() {
		if len(resources.Resources.Storage) > o.cfg.MaxGroupVolumes {
			o.log.Info(fmt.Sprintf("unable to fulfill: group volumes count exceeds (%d > %d)", len(resources.Resources.Storage), o.cfg.MaxGroupVolumes))
			return false, nil
		}
	}
	// TODO(spheron): check wallet requirements
	if err := group.GroupSpec.ValidateBasic(); err != nil {
		o.log.Error("unable to fulfill: group validation error",
			"err", err)
		return false, nil
	}
	return true, nil
}
