package rest

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	kubeErrors "k8s.io/apimachinery/pkg/api/errors"
	kubeVersion "k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/tools/remotecommand"

	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	manifest "github.com/akash-network/akash-api/go/manifest/v2beta2"
	manifestValidation "github.com/akash-network/akash-api/go/manifest/v2beta2"
	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/akash-network/akash-api/go/sdkutil"

	"github.com/akash-network/node/pubsub"
	"github.com/akash-network/node/util/wsutil"

	"github.com/akash-network/provider"
	"github.com/akash-network/provider/cluster"
	"github.com/akash-network/provider/cluster/kube/builder"
	kubeclienterrors "github.com/akash-network/provider/cluster/kube/errors"
	cltypes "github.com/akash-network/provider/cluster/types/v1beta3"
	cip "github.com/akash-network/provider/cluster/types/v1beta3/clients/ip"
	clfromctx "github.com/akash-network/provider/cluster/types/v1beta3/fromctx"
	pmanifest "github.com/akash-network/provider/manifest"
	"github.com/akash-network/provider/tools/fromctx"
	"github.com/akash-network/provider/version"
)

type CtxAuthKey string

const (
	contentTypeJSON = "application/json; charset=UTF-8"

	// Time allowed to write the file to the client.
	pingWait = 15 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 15 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = 10 * time.Second
)

const (
	// as per RFC https://www.iana.org/assignments/websocket/websocket.xhtml#close-code-number
	// errors from private use staring
	websocketInternalServerErrorCode = 4000
	websocketLeaseNotFound           = 4001
	manifestSubmitTimeout            = 120 * time.Second
)

type wsStreamConfig struct {
	lid       mtypes.LeaseID
	services  string
	follow    bool
	tailLines *int64
	log       log.Logger
	client    cluster.ReadClient
}

func newRouter(log log.Logger, addr sdk.Address, pclient provider.Client, ctxConfig map[interface{}]interface{}, middlewares ...mux.MiddlewareFunc) *mux.Router {
	router := mux.NewRouter()

	// store provider address in context as lease endpoints below need it
	//ILIJA FIX 1
	// router.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		gcontext.Set(r, providerContextKey, addr)

	// 		next.ServeHTTP(w, r)
	// 	})
	// })
	//ILIJA FIX 2
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gcontext.Set(r, providerContextKey, "provider")

			next.ServeHTTP(w, r)
		})
	})

	router.Use(middlewares...)

	// GET /version
	// provider version endpoint does not require authentication
	router.HandleFunc("/version",
		createVersionHandler(log, pclient)).
		Methods(http.MethodGet)

	// GET /address
	// provider status endpoint does not require authentication
	router.HandleFunc("/address",
		createAddressHandler(log, addr)).
		Methods("GET")

	// GET /status
	// provider status endpoint does not require authentication
	router.HandleFunc("/status",
		createStatusHandler(log, pclient, addr)).
		Methods("GET")

	vrouter := router.NewRoute().Subrouter()
	vrouter.Use(requireOwner())

	// GET /validate
	// validate endpoint checks if provider will bid on given groupspec
	vrouter.HandleFunc("/validate",
		validateHandler(log, pclient)).
		Methods("GET")

	// GET /wiboy (aka would I bid on you)
	// validate endpoint checks if provider will bid on given groupspec
	vrouter.HandleFunc("/wiboy",
		validateHandler(log, pclient)).
		Methods("GET")

	// PUT
	shperonRouter := router.PathPrefix(spheronPathPrefix).Subrouter()
	shperonRouter.HandleFunc("/bid",
		createDeploymentCreateHandler(log, pclient.Bus())).
		Methods(http.MethodPut)
	shperonRouter.HandleFunc("/lease",
		createLeasetCreateHandler(log, pclient.Bus())).
		Methods(http.MethodPut)
	shperonRouter.HandleFunc("/deployment/close",
		createDeploymentCloseHandler(log, pclient.Bus())).
		Methods(http.MethodDelete)
	shperonRouter.HandleFunc("/deployment/update",
		createDeploymentUpdateHandler(log, pclient.Bus())).
		Methods(http.MethodPut)

	hostnameRouter := router.PathPrefix(hostnamePrefix).Subrouter()
	hostnameRouter.Use(requireOwner())
	hostnameRouter.HandleFunc(migratePathPrefix, migrateHandler(log, pclient.Hostname(), pclient.ClusterService())).
		Methods(http.MethodPost)

	endpointRouter := router.PathPrefix(endpointPrefix).Subrouter()
	endpointRouter.Use(requireOwner())
	endpointRouter.HandleFunc(migratePathPrefix, migrateEndpointHandler(log, pclient.ClusterService(), pclient.Cluster())).
		Methods(http.MethodPost)

	// PUT /deployment/manifest
	drouter := router.PathPrefix(deploymentPathPrefix).Subrouter()
	drouter.Use(
		requireOwner(),
		requireDeploymentID(),
	)

	drouter.HandleFunc("/manifest",
		createManifestHandler(log, pclient.Manifest())).
		Methods(http.MethodPut)

	lrouter := router.PathPrefix(leasePathPrefix).Subrouter()
	lrouter.Use(
		requireOwner(),
		requireLeaseID(),
	)

	// GET /lease/<lease-id>/status
	lrouter.HandleFunc("/status",
		leaseStatusHandler(log, pclient.Cluster(), ctxConfig)).
		Methods(http.MethodGet)

	// GET /lease/<lease-id>/kubeevents
	eventsRouter := lrouter.PathPrefix("/kubeevents").Subrouter()
	eventsRouter.Use(
		requestStreamParams(),
	)
	eventsRouter.HandleFunc("",
		leaseKubeEventsHandler(log, pclient.Cluster())).
		Methods("GET")

	logRouter := lrouter.PathPrefix("/logs").Subrouter()
	logRouter.Use(
		requestStreamParams(),
	)

	// GET /lease/<lease-id>/logs
	logRouter.HandleFunc("",
		leaseLogsHandler(log, pclient.Cluster())).
		Methods("GET")

	srouter := lrouter.PathPrefix("/service/{serviceName}").Subrouter()
	srouter.Use(
		requireService(),
	)

	// GET /lease/<lease-id>/service/<service-name>/status
	srouter.HandleFunc("/status",
		leaseServiceStatusHandler(log, pclient.Cluster())).
		Methods("GET")

	// POST /lease/<lease-id>/shell
	lrouter.HandleFunc("/shell",
		leaseShellHandler(log, pclient.Manifest(), pclient.Cluster()))

	return router
}

func newJwtServerRouter(addr sdk.Address, privateKey interface{}, jwtExpiresAfter time.Duration, certSerialNumber string) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/jwt",
		jwtServiceHandler(addr, privateKey, jwtExpiresAfter, certSerialNumber)).
		Methods("GET")

	return router
}

func newResourceServerRouter(log log.Logger, providerAddr sdk.Address, publicKey *ecdsa.PublicKey, lokiGwAddr string) *mux.Router {
	router := mux.NewRouter()

	// add a middleware to verify the JWT provided in Authorization header
	router.Use(resourceServerAuth(log, providerAddr, publicKey))

	lrouter := router.PathPrefix(leasePathPrefix).Subrouter()
	lrouter.Use(requireLeaseID())

	lokiServiceRouter := lrouter.PathPrefix("/loki-service").Subrouter()
	lokiServiceRouter.NewRoute().Handler(lokiServiceHandler(log, lokiGwAddr))

	return router
}

// lokiServiceHandler forwards all requests to the loki instance running in provider's cluster.
// Example:
//
//	Incoming Request: http://localhost:8445/lease/1/1/1/loki-service/loki/api/v1/query?query={app=".+"}
//	Outgoing Request: http://{lokiGwAddr}/loki/api/v1/query?query={app=".+"}
func lokiServiceHandler(log log.Logger, lokiGwAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set the X-Scope-OrgID header for fetching logs for the right tenant
		r.Header.Set("X-Scope-OrgID", builder.LidNS(requestLeaseID(r)))

		// build target url for the reverse proxy
		scheme := "http" // for http & https
		if strings.HasPrefix(r.URL.Scheme, "ws") {
			scheme = "ws" // for ws & wss
		}
		lokiURL, err := url.Parse(fmt.Sprintf("%s://%s", scheme, lokiGwAddr))
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reverseProxy := httputil.NewSingleHostReverseProxy(lokiURL)

		// remove the "/lease/{dseq}/{gseq}/{oseq}/loki-service" path prefix from the request url
		// before it is sent to the reverse proxy.
		pathSplits := strings.SplitN(r.URL.Path, "/", 7)
		if len(pathSplits) < 7 || pathSplits[6] == "" {
			log.Error("loki api not provided in url")
			http.Error(w, "loki api not provided in url", http.StatusBadRequest)
			return
		}
		r.URL.Path = pathSplits[6]

		// serve the request using the reverse proxy
		log.Info("Forwarding request to loki", "HTTP_API", pathSplits[6])
		reverseProxy.ServeHTTP(w, r)
	}
}

func jwtServiceHandler(paddr sdk.Address, privateKey interface{}, jwtExpiresAfter time.Duration, certSerialNumber string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		now := time.Now()
		claim := ClientCustomClaims{
			AkashNamespace: &AkashNamespace{
				V1: &ClaimsV1{
					CertSerialNumber: certSerialNumber,
				},
			},
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(now.Add(jwtExpiresAfter)),
				IssuedAt:  jwt.NewNumericDate(now),
				// account address of the tenant: trustable as it has already been verified by mTLS
				Subject: request.TLS.PeerCertificates[0].Subject.CommonName,
				Issuer:  paddr.String(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, &claim)
		jwtString, err := token.SignedString(privateKey)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = io.WriteString(writer, jwtString)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

type channelToTerminalSizeQueue <-chan remotecommand.TerminalSize

func (sq channelToTerminalSizeQueue) Next() *remotecommand.TerminalSize {
	v, ok := <-sq
	if !ok {
		return nil
	}

	return &v // Interface is dumb and use a pointer
}

type leaseShellResponse struct {
	ExitCode int    `json:"exit_code"`
	Message  string `json:"message,omitempty"`
}

func leaseShellHandler(log log.Logger, mclient pmanifest.Client, cclient cluster.Client) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		leaseID := requestLeaseID(req)

		//  check if deployment actually exists in the first place before querying kubernetes
		active, err := mclient.IsActive(req.Context(), leaseID.DeploymentID())
		if err != nil {
			log.Error("failed checking deployment activity", "err", err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !active {
			log.Info("no active deployment", "lease", leaseID)
			rw.WriteHeader(http.StatusNotFound)
			return
		}

		localLog := log.With("lease", leaseID.String(), "action", "shell")

		vars := req.URL.Query()
		var cmd []string

		for i := 0; true; i++ {
			v := vars.Get(fmt.Sprintf("cmd%d", i))
			if 0 == len(v) {
				break
			}
			cmd = append(cmd, v)
		}

		if len(cmd) == 0 {
			localLog.Error("missing cmd parameter")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		tty := vars.Get("tty")
		if 0 == len(tty) {
			localLog.Error("missing parameter tty")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		isTty := tty == "1"

		service := vars.Get("service")
		if 0 == len(service) {
			localLog.Error("missing parameter service")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		stdin := vars.Get("stdin")
		if 0 == len(stdin) {
			localLog.Error("missing parameter stdin")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		connectStdin := stdin == "1"

		podIndexStr := vars.Get("podIndex")
		if len(podIndexStr) == 0 {
			localLog.Error("missing parameter podIndex")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		podIndex64, err := strconv.ParseUint(podIndexStr, 0, 31)
		if err != nil {
			localLog.Error("parameter podIndex invalid", "err", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		podIndex := uint(podIndex64)

		upgrader := websocket.Upgrader{
			ReadBufferSize:  0,
			WriteBufferSize: 0,
		}

		shellWs, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			// At this point the connection either has a response sent already
			// or it has been closed
			localLog.Error("failed handshake", "err", err)
			return
		}

		var stdinPipeOut *io.PipeWriter
		var stdinPipeIn *io.PipeReader
		wg := &sync.WaitGroup{}

		var tsq remotecommand.TerminalSizeQueue
		var terminalSizeUpdate chan remotecommand.TerminalSize
		if isTty {
			terminalSizeUpdate = make(chan remotecommand.TerminalSize, 1)
			tsq = channelToTerminalSizeQueue(terminalSizeUpdate)
		}

		if connectStdin {
			stdinPipeIn, stdinPipeOut = io.Pipe()

			wg.Add(1)
			go leaseShellWebsocketHandler(localLog, wg, shellWs, stdinPipeOut, terminalSizeUpdate)
		}

		l := &sync.Mutex{}
		stdout := wsutil.NewWsWriterWrapper(shellWs, LeaseShellCodeStdout, l)
		stderr := wsutil.NewWsWriterWrapper(shellWs, LeaseShellCodeStderr, l)

		subctx, subcancel := context.WithCancel(req.Context())
		wg.Add(1)
		go leaseShellPingHandler(subctx, wg, shellWs)

		var stdinForExec io.Reader
		if connectStdin {
			stdinForExec = stdinPipeIn
		}
		result, err := cclient.Exec(subctx, leaseID, service, podIndex, cmd, stdinForExec, stdout, stderr, isTty, tsq)
		subcancel()

		responseData := leaseShellResponse{}
		var resultWriter io.Writer
		encodeData := true
		resultWriter = wsutil.NewWsWriterWrapper(shellWs, LeaseShellCodeResult, l)

		if result != nil {
			responseData.ExitCode = result.ExitCode()

			localLog.Info("lease shell completed", "exitcode", result.ExitCode())
		} else {
			if cluster.ErrorIsOkToSendToClient(err) {
				responseData.Message = err.Error()
			} else {
				resultWriter = wsutil.NewWsWriterWrapper(shellWs, LeaseShellCodeFailure, l)
				// Don't return errors like this to the client, they could contain information
				// that should not be let out
				encodeData = false

				localLog.Error("lease exec failed", "err", err)
			}
		}

		if encodeData {
			encoder := json.NewEncoder(resultWriter)
			err = encoder.Encode(responseData)
		} else {
			// Just send an empty message so the remote knows things are over
			_, err = resultWriter.Write([]byte{})
		}

		_ = shellWs.Close()

		if err != nil {
			localLog.Error("failed writing response to client after exec", "err", err)
		}

		wg.Wait()

		if stdinPipeOut != nil {
			_ = stdinPipeOut.Close()
		}
		if stdinPipeIn != nil {
			_ = stdinPipeIn.Close()
		}

		if terminalSizeUpdate != nil {
			close(terminalSizeUpdate)
		}
	}
}

func createAddressHandler(log log.Logger, providerAddr sdk.Address) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := struct {
			Address string `json:"address"`
		}{
			Address: providerAddr.String(),
		}
		writeJSON(log, w, data)
	}
}

type versionInfo struct {
	Akash version.Info      `json:"akash"`
	Kube  *kubeVersion.Info `json:"kube"`
}

func createVersionHandler(log log.Logger, pclient provider.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		kube, err := pclient.Cluster().KubeVersion()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(log, w, versionInfo{
			Akash: version.NewInfo(),
			Kube:  kube,
		})
	}
}

func createStatusHandler(log log.Logger, sclient provider.StatusClient, providerAddr sdk.Address) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		status, err := sclient.Status(req.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := struct {
			provider.Status
			Address string `json:"address"`
		}{
			Status:  *status,
			Address: providerAddr.String(),
		}
		writeJSON(log, w, data)
	}
}

func validateHandler(log log.Logger, cl provider.ValidateClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(data) == 0 {
			http.Error(w, "empty payload", http.StatusBadRequest)
			return
		}

		owner := requestOwner(req)

		var gspec dtypes.GroupSpec

		if err := json.Unmarshal(data, &gspec); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		validate, err := cl.Validate(req.Context(), owner, gspec)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(log, w, validate)
	}
}

func createManifestHandler(log log.Logger, mclient pmanifest.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var mani manifest.Manifest
		decoder := json.NewDecoder(req.Body)
		defer func() {
			_ = req.Body.Close()
		}()

		if err := decoder.Decode(&mani); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		subctx, cancel := context.WithTimeout(req.Context(), manifestSubmitTimeout)
		defer cancel()
		if err := mclient.Submit(subctx, requestDeploymentID(req), mani); err != nil {
			if errors.Is(err, manifestValidation.ErrInvalidManifest) {
				http.Error(w, err.Error(), http.StatusUnprocessableEntity)
				return
			}
			if errors.Is(err, pmanifest.ErrNoLeaseForDeployment) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			log.Error("manifest submit failed", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func createDeploymentCreateHandler(log log.Logger, bus pubsub.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		defer func() {
			_ = req.Body.Close()
		}()

		bus.Publish(mtypes.EventOrderCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "bid-created"}, ID: mtypes.OrderID{
			Owner: "owner",
			DSeq:  15,
			GSeq:  1,
			OSeq:  1,
		}})
	}
}

func createDeploymentUpdateHandler(log log.Logger, bus pubsub.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		defer func() {
			_ = req.Body.Close()
		}()

		v, _ := hex.DecodeString("1")

		bus.Publish(dtypes.EventDeploymentUpdated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-updated"}, ID: dtypes.DeploymentID{
			Owner: "owner",
			DSeq:  15,
		}, Version: v,
		})
	}
}

func createLeasetCreateHandler(log log.Logger, bus pubsub.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var request struct {
			Dseq uint64 `json:"dseq"`
		}

		decoder := json.NewDecoder(req.Body)
		defer func() {
			_ = req.Body.Close()
		}()

		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		bus.Publish(mtypes.EventLeaseCreated{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-created"}, ID: mtypes.LeaseID{
			Owner:    "owner",
			DSeq:     request.Dseq,
			GSeq:     1,
			OSeq:     1,
			Provider: "provider",
		}, Price: sdk.DecCoin{
			Denom:  "uakt",
			Amount: types.OneDec(),
		}})
	}
}

func createDeploymentCloseHandler(log log.Logger, bus pubsub.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var request struct {
			Dseq uint64 `json:"dseq"`
		}

		decoder := json.NewDecoder(req.Body)
		defer func() {
			_ = req.Body.Close()
		}()

		if err := decoder.Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		bus.Publish(dtypes.EventDeploymentClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-close"}, ID: dtypes.DeploymentID{
			Owner: "owner",
			DSeq:  request.Dseq,
		}})
		bus.Publish(mtypes.EventLeaseClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "lease-close"}, ID: mtypes.LeaseID{
			Owner:    "owner",
			DSeq:     request.Dseq,
			GSeq:     1,
			OSeq:     1,
			Provider: "provider",
		}, Price: sdk.DecCoin{
			Denom:  "uakt",
			Amount: types.OneDec(),
		}})
		bus.Publish(dtypes.EventGroupClosed{Context: sdkutil.BaseModuleEvent{Module: "market", Action: "deployment-close"}, ID: dtypes.GroupID{
			Owner: "owner",
			DSeq:  request.Dseq,
			GSeq:  1,
		}})
	}
}

func leaseKubeEventsHandler(log log.Logger, cclient cluster.ReadClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			// At this point the connection either has a response sent already
			// or it has been closed
			return
		}

		wsEventWriter(r.Context(), ws, wsStreamConfig{
			lid:      requestLeaseID(r),
			follow:   requestLogFollow(r),
			services: requestServices(r),
			log:      log,
			client:   cclient,
		})
	}
}

func leaseStatusHandler(log log.Logger, cclient cluster.ReadClient, clusterSettings map[interface{}]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := fromctx.ApplyToContext(req.Context(), clusterSettings)

		leaseID := requestLeaseID(req)
		result := LeaseStatus{}

		found, manifestGroup, err := cclient.GetManifestGroup(req.Context(), leaseID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !found { // If the manifest doesn't exist, there is no lease
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var ipLeaseStatus []cip.LeaseIPStatus

		if clIP := clfromctx.ClientIPFromContext(ctx); clIP != nil {
			hasLeasedIPs := false

		ipManifestGroupSearchLoop:
			for _, service := range manifestGroup.Services {
				for _, expose := range service.Expose {
					if 0 != len(expose.IP) {
						hasLeasedIPs = true
						break ipManifestGroupSearchLoop
					}
				}
			}

			if hasLeasedIPs {
				log.Debug("querying for IP address status", "lease-id", leaseID)
				ipLeaseStatus, err = clIP.GetIPAddressStatus(req.Context(), leaseID.OrderID())
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				result.IPs = make(map[string][]LeasedIPStatus)

				for _, ipLease := range ipLeaseStatus {
					entries := result.IPs[ipLease.ServiceName]
					if entries == nil {
						entries = make([]LeasedIPStatus, 0)
					}

					entries = append(entries, LeasedIPStatus{
						Port:         ipLease.Port,
						ExternalPort: ipLease.ExternalPort,
						Protocol:     ipLease.Protocol,
						IP:           ipLease.IP,
					})

					result.IPs[ipLease.ServiceName] = entries
				}
			}
		}

		hasForwardedPorts := false
	portManifestGroupSearchLoop:
		for _, service := range manifestGroup.Services {
			for _, expose := range service.Expose {
				if expose.Global && expose.ExternalPort != 80 {
					hasForwardedPorts = true
					break portManifestGroupSearchLoop
				}
			}
		}
		if hasForwardedPorts {
			result.ForwardedPorts, err = cclient.ForwardedPortStatus(ctx, leaseID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		result.Services, err = cclient.LeaseStatus(ctx, leaseID)
		if err != nil {
			if errors.Is(err, kubeclienterrors.ErrNoDeploymentForLease) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, kubeclienterrors.ErrLeaseNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if kubeErrors.IsNotFound(err) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		writeJSON(log, w, result)
	}
}

func leaseServiceStatusHandler(log log.Logger, cclient cluster.ReadClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		status, err := cclient.ServiceStatus(req.Context(), requestLeaseID(req), requestService(req))
		if err != nil {
			if errors.Is(err, kubeclienterrors.ErrNoDeploymentForLease) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if errors.Is(err, kubeclienterrors.ErrLeaseNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if kubeErrors.IsNotFound(err) {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		writeJSON(log, w, status)
	}
}

func leaseLogsHandler(log log.Logger, cclient cluster.ReadClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			// At this point the connection either has a response sent already
			// or it has been closed
			return
		}

		wsLogWriter(r.Context(), ws, wsStreamConfig{
			lid:       requestLeaseID(r),
			services:  requestServices(r),
			follow:    requestLogFollow(r),
			tailLines: requestLogTailLines(r),
			log:       log,
			client:    cclient,
		})
	}
}

func wsSetupPongHandler(ws *websocket.Conn, cancel func()) error {
	if err := ws.SetReadDeadline(time.Time{}); err != nil {
		return err
	}

	ws.SetPongHandler(func(string) error {
		return ws.SetReadDeadline(time.Now().Add(pingWait))
	})

	go func() {
		var err error

		defer func() {
			if err != nil {
				cancel()
			}
		}()

		for {
			var mtype int
			if mtype, _, err = ws.ReadMessage(); err != nil {
				break
			}

			if mtype == websocket.CloseMessage {
				err = errors.Errorf("disconnect")
			}
		}
	}()

	return nil
}

func wsLogWriter(ctx context.Context, ws *websocket.Conn, cfg wsStreamConfig) {
	pingTicker := time.NewTicker(pingPeriod)

	cctx, cancel := context.WithCancel(ctx)
	defer func() {
		pingTicker.Stop()
		cancel()
		_ = ws.Close()
	}()

	logs, err := cfg.client.LeaseLogs(cctx, cfg.lid, cfg.services, cfg.follow, cfg.tailLines)
	if err != nil {
		cfg.log.Error("couldn't fetch logs", "error", err.Error())
		err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocketInternalServerErrorCode, ""))
		if err != nil {
			cfg.log.Error("couldn't push control message through websocket: %s", err.Error())
		}
		return
	}

	if len(logs) == 0 {
		_ = ws.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocketInternalServerErrorCode, "no running pods"))
		return
	}

	if err = wsSetupPongHandler(ws, cancel); err != nil {
		return
	}

	var scanners sync.WaitGroup

	logch := make(chan ServiceLogMessage)

	scanners.Add(len(logs))

	for _, lg := range logs {
		go func(name string, scan *bufio.Scanner) {
			defer scanners.Done()

			for scan.Scan() && ctx.Err() == nil {
				logch <- ServiceLogMessage{
					Name:    name,
					Message: scan.Text(),
				}
			}
		}(lg.Name, lg.Scanner)
	}

	donech := make(chan struct{})

	go func() {
		scanners.Wait()
		close(donech)
	}()

done:
	for {
		select {
		case line := <-logch:
			if err = ws.WriteJSON(line); err != nil {
				break done
			}
		case <-pingTicker.C:
			if err = ws.WriteControl(websocket.PingMessage, nil, time.Now().Add(5*time.Second)); err != nil {
				break done
			}
			if err = ws.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
				break done
			}
		case <-donech:
			break done
		}
	}

	cancel()

	for i := range logs {
		_ = logs[i].Stream.Close()
	}

	// drain logs channel in separate goroutine to unblock seeders waiting for write space
	go func() {
	drain:
		for {
			select {
			case <-donech:
				break drain
			case <-logch:
			}
		}
	}()
}

func wsEventWriter(ctx context.Context, ws *websocket.Conn, cfg wsStreamConfig) {
	pingTicker := time.NewTicker(pingPeriod)
	cctx, cancel := context.WithCancel(ctx)
	defer func() {
		pingTicker.Stop()
		cancel()
		_ = ws.Close()
	}()

	evts, err := cfg.client.LeaseEvents(cctx, cfg.lid, cfg.services, cfg.follow)
	if err != nil {
		cfg.log.Error("couldn't fetch events", "error", err.Error())
		err = ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocketInternalServerErrorCode, ""))
		if err != nil {
			cfg.log.Error("couldn't push control message through websocket", "error", err.Error())
		}
		return
	}

	if evts == nil {
		err = ws.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocketLeaseNotFound, ""))
		if err != nil {
			cfg.log.Error("couldn't push control message through websocket", "error", err.Error())
		}
		return
	}

	defer evts.Shutdown()

	if err = wsSetupPongHandler(ws, cancel); err != nil {
		return
	}

	sendClose := func() {
		_ = ws.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	}
done:
	for {
		select {
		case <-ctx.Done():
			sendClose()
			break done
		case <-evts.Done():
			sendClose()
			break done
		case evt := <-evts.ResultChan():
			if evt == nil {
				sendClose()
				break done
			}

			if err = ws.WriteJSON(cltypes.LeaseEvent{
				Type:                evt.Type,
				ReportingController: evt.ReportingController,
				ReportingInstance:   evt.ReportingInstance,
				Reason:              evt.Reason,
				Note:                evt.Note,
				Object: cltypes.LeaseEventObject{
					Kind:      evt.Regarding.Kind,
					Namespace: evt.Regarding.Namespace,
					Name:      evt.Regarding.Name,
				},
			}); err != nil {
				break done
			}
		case <-pingTicker.C:
			if err = ws.WriteControl(websocket.PingMessage, nil, time.Now().Add(5*time.Second)); err != nil {
				break done
			}
			if err = ws.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
				break done
			}
		}
	}
}

func writeJSON(log log.Logger, w http.ResponseWriter, obj interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentTypeJSON)

	_, err = w.Write(bytes)
	if err != nil {
		log.Error("error writing response", "err", err)
		return
	}
}
