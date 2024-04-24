package rest

import (
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/context"
	gcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	dtypes "github.com/akash-network/akash-api/go/node/deployment/v1beta3"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"
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
			// ILIJA FIX 1
			// if r.TLS == nil || len(r.TLS.PeerCertificates) == 0 {
			// 	http.Error(w, "", http.StatusUnauthorized)
			// 	return
			// }
			// ILIJA FIX 2

			// // at this point client certificate has been validated
			// // so only thing left to do is get account id stored in the CommonName
			// owner, err := sdk.AccAddressFromBech32("owner")

			// fmt.Printf("requireOwner %+v\n", owner)

			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusUnauthorized)
			// 	return
			// }

			context.Set(r, ownerContextKey, "owner")
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

	//ILIJA FIX MASTER ZA WALLETI
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

	//ILIJA FIX
	// return mquery.ParseLeasePath(parts)
	//ILIJA FIX
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

func resourceServerAuth(log log.Logger, providerAddr sdk.Address, publicKey *ecdsa.PublicKey) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// verify the provided JWT
			token, err := jwt.ParseWithClaims(r.Header.Get("Authorization"), &ClientCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				// return the public key to be used for JWT verification
				return publicKey, nil
			})
			if err != nil {
				log.Error("falied to parse JWT", "error", err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// delete the Authorization header as it is no more needed
			r.Header.Del("Authorization")

			// store the owner & provider address in request context to be used in later handlers
			customClaims, ok := token.Claims.(*ClientCustomClaims)
			if !ok {
				log.Error("failed to parse JWT claims")
				http.Error(w, "Invalid JWT", http.StatusUnauthorized)
				return
			}
			ownerAddress, err := sdk.AccAddressFromBech32(customClaims.Subject)
			if err != nil {
				log.Error("failed parsing owner address", "error", err, "address", customClaims.Subject)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			gcontext.Set(r, ownerContextKey, ownerAddress)
			//ILIJA FIX 1
			// gcontext.Set(r, providerContextKey, providerAddr)
			//ILIJA FIX 2
			gcontext.Set(r, providerContextKey, "provider")

			next.ServeHTTP(w, r)
		})
	}
}
