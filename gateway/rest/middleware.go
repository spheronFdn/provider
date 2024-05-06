package rest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
	"github.com/akash-network/provider/spheron"
)

type contextKey int

const (
	leaseContextKey contextKey = iota + 1
	deploymentContextKey
	logFollowContextKey
	tailLinesContextKey
	serviceContextKey
	ownerContextKey
	providerContextKey
	servicesContextKey
)

func requestLeaseID(req *http.Request) mtypes.LeaseID {
	return context.Get(req, leaseContextKey).(mtypes.LeaseID)
}

func requestLogFollow(req *http.Request) bool {
	return context.Get(req, logFollowContextKey).(bool)
}

func requestLogTailLines(req *http.Request) *int64 {
	return context.Get(req, tailLinesContextKey).(*int64)
}

func requestService(req *http.Request) string {
	return context.Get(req, serviceContextKey).(string)
}

func requestServices(req *http.Request) string {
	return context.Get(req, servicesContextKey).(string)
}

func requestProvider(req *http.Request) string {
	return context.Get(req, providerContextKey).(string)
}

func requestOwner(req *http.Request) string {
	return context.Get(req, ownerContextKey).(string)
}

func requestDeploymentID(req *http.Request) dtypes.DeploymentID {
	return context.Get(req, deploymentContextKey).(dtypes.DeploymentID)
}

func requireOwner() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO(spheron) : Use our custom authorization
			// Extract and decode the authorization header
			authHeader := r.Header.Get("Auth-Spheron")
			if authHeader == "" {
				http.Error(w, "missing Auth-Spheron header", http.StatusUnauthorized)
				return
			}

			authHeaderDecoded, err := base64.StdEncoding.DecodeString(authHeader)
			if err != nil {
				http.Error(w, "invalid Auth-Spheron header", http.StatusBadRequest)
				return
			}

			// Unmarshal the decoded string into the AuthJson struct
			var authData spheron.AuthJson
			if err := json.Unmarshal(authHeaderDecoded, &authData); err != nil {
				http.Error(w, "invalid JSON format", http.StatusBadRequest)
				return
			}

			// TODO(spheron): Implement custom authorization using the `authData` values
			// Check if timestamp is in range of -20sec:now, and if user actually signed it with private key !
			pubKey := authData.PubKey

			// Set the owner information into the request context
			context.Set(r, ownerContextKey, pubKey)
			next.ServeHTTP(w, r)
		})
	}
}

func requireDeploymentID() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id, err := parseDeploymentID(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			context.Set(req, deploymentContextKey, id)
			next.ServeHTTP(w, req)
		})
	}
}

func requireLeaseID() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			id, err := parseLeaseID(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fmt.Printf("requireLeaseID %+v\n", id)

			context.Set(req, leaseContextKey, id)
			next.ServeHTTP(w, req)
		})
	}
}

func requireService() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			vars := mux.Vars(req)

			svc := vars["serviceName"]
			if svc == "" {
				http.Error(w, "empty service name", http.StatusBadRequest)
				return
			}

			context.Set(req, serviceContextKey, svc)
			next.ServeHTTP(w, req)
		})
	}
}

func parseDeploymentID(req *http.Request) (dtypes.DeploymentID, error) {
	var parts []string
	parts = append(parts, requestOwner(req))
	parts = append(parts, mux.Vars(req)["dseq"])

	// used to return : return dtypes.ParseDeploymentPath(parts)
	// Spheron fix
	dseq, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return dtypes.DeploymentID{}, err
	}

	return dtypes.DeploymentID{
		Owner: parts[0],
		DSeq:  dseq,
	}, nil
}

func parseLeaseID(req *http.Request) (mtypes.LeaseID, error) {
	vars := mux.Vars(req)

	fmt.Printf("parseLeaseID %+v\n", vars)

	parts := []string{
		requestOwner(req),
		vars["dseq"],
		vars["gseq"],
		vars["oseq"],
		requestProvider(req),
	}

	fmt.Printf("parseLeaseID parts %+v\n", parts)

	dseq, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return mtypes.LeaseID{}, err
	}
	gseq, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return mtypes.LeaseID{}, err
	}
	oseq, err := strconv.ParseUint(parts[3], 10, 32)
	if err != nil {
		return mtypes.LeaseID{}, err
	}

	// Spheron fix: replace provider extraction to not use cosmosdk
	// used to run: return mquery.ParseLeasePath(parts)
	return mtypes.LeaseID{
		Owner:    parts[0],
		DSeq:     dseq,
		GSeq:     uint32(gseq),
		OSeq:     uint32(oseq),
		Provider: parts[4],
	}, nil
}

func requestStreamParams() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			vars := req.URL.Query()

			var err error

			defer func() {
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}()

			var tailLines *int64

			services := vars.Get("service")
			if strings.HasSuffix(services, ",") {
				err = errors.Errorf("parameter \"service\" must not contain trailing comma")
				return
			}

			follow := false

			if val := vars.Get("follow"); val != "" {
				follow, err = strconv.ParseBool(val)
				if err != nil {
					return
				}
			}

			vl := new(int64)
			if val := vars.Get("tail"); val != "" {
				*vl, err = strconv.ParseInt(val, 10, 32)
				if err != nil {
					return
				}

				if *vl < -1 {
					err = errors.Errorf("parameter \"tail\" contains invalid value")
					return
				}
			} else {
				*vl = -1
			}

			if *vl > -1 {
				tailLines = vl
			}

			context.Set(req, logFollowContextKey, follow)
			context.Set(req, tailLinesContextKey, tailLines)
			context.Set(req, servicesContextKey, services)

			next.ServeHTTP(w, req)
		})
	}
}
