package rest

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"

	gwutils "github.com/akash-network/provider/gateway/utils"
	"github.com/akash-network/provider/spheron"
)

func NewServer(t testing.TB, qclient spheron.Client, handler http.Handler, certs []tls.Certificate) *httptest.Server {
	t.Helper()

	ts := httptest.NewUnstartedServer(handler)

	var err error
	ts.TLS, err = gwutils.NewServerTLSConfig(context.Background(), certs)
	if err != nil {
		t.Fatal(err.Error())
	}

	ts.StartTLS()

	return ts
}
