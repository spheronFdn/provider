package rest

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"k8s.io/client-go/tools/remotecommand"

	cosmosclient "github.com/cosmos/cosmos-sdk/client"

	manifest "github.com/akash-network/akash-api/go/manifest/v2beta2"
	mtypes "github.com/akash-network/akash-api/go/node/market/v1beta4"

	"github.com/akash-network/provider"
	cltypes "github.com/akash-network/provider/cluster/types/v1beta3"
	"github.com/akash-network/provider/spheron"
)

const (
	schemeWSS   = "wss"
	schemeHTTPS = "https"
)

// Client defines the methods available for connecting to the gateway server.
type Client interface {
	Status(ctx context.Context) (*provider.Status, error)
	SubmitManifest(ctx context.Context, dseq uint64, mani manifest.Manifest) error
	LeaseStatus(ctx context.Context, id mtypes.LeaseID) (LeaseStatus, error)
	LeaseEvents(ctx context.Context, id mtypes.LeaseID, services string, follow bool) (*LeaseKubeEvents, error)
	LeaseLogs(ctx context.Context, id mtypes.LeaseID, services string, follow bool, tailLines int64) (*ServiceLogs, error)
	ServiceStatus(ctx context.Context, id mtypes.LeaseID, service string) (*cltypes.ServiceStatus, error)
	LeaseShell(ctx context.Context, id mtypes.LeaseID, service string, podIndex uint, cmd []string,
		stdin io.Reader,
		stdout io.Writer,
		stderr io.Writer,
		tty bool,
		tsq <-chan remotecommand.TerminalSize) error
	MigrateHostnames(ctx context.Context, hostnames []string, dseq uint64, gseq uint32) error
	MigrateEndpoints(ctx context.Context, endpoints []string, dseq uint64, gseq uint32) error
}

type LeaseKubeEvent struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ServiceLogMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type LeaseKubeEvents struct {
	Stream  <-chan cltypes.LeaseEvent
	OnClose <-chan string
}

type ServiceLogs struct {
	Stream  <-chan ServiceLogMessage
	OnClose <-chan string
}

// NewClient returns a new Client
func NewClient(spheronClient spheron.Client, addr string, authToken string) (Client, error) {
	// addres will be provider address so you need to check provider details and set client with provider uri*
	// TODO(spheron): query the chain for provider details and return the client.
	uri, err := url.Parse("https://localhost:8443")
	if err != nil {
		return nil, err
	}

	return newClient(spheronClient, addr, uri, authToken), nil
}

func newClient(spheronClient spheron.Client, addr string, uri *url.URL, authToken string) *client {
	cl := &client{
		host:          uri,
		addr:          addr,
		spheronClient: spheronClient,
	}

	tlsConfig := &tls.Config{
		// must use Hostname rather than Host field as certificate is issued for host without port
		ServerName:            uri.Hostname(),
		InsecureSkipVerify:    true, // nolint: gosec
		VerifyPeerCertificate: cl.verifyPeerCertificate,
		MinVersion:            tls.VersionTLS13,
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		// Never  follow redirects
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     nil,
		Timeout: 0,
	}

	cl.hclient = httpClient

	cl.wsclient = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		TLSClientConfig:  tlsConfig,
	}

	return cl
}

type ClientDirectory struct {
	cosmosContext cosmosclient.Context
	clients       map[string]Client
	clientCert    tls.Certificate

	lock sync.Mutex
}

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type client struct {
	host          *url.URL
	hclient       httpClient
	wsclient      *websocket.Dialer
	addr          string
	spheronClient spheron.Client
}

type ClientCustomClaims struct {
	AkashNamespace *AkashNamespace `json:"https://akash.network/"`
	jwt.RegisteredClaims
}

type AkashNamespace struct {
	V1 *ClaimsV1 `json:"v1"`
}

type ClaimsV1 struct {
	CertSerialNumber string `json:"cert_serial_number"`
}

var errRequiredCertSerialNum = errors.New("cert_serial_number must be present in claims")
var errNonNumericCertSerialNum = errors.New("cert_serial_number must be numeric in claims")

type ClientResponseError struct {
	Status  int
	Message string
}

func (err ClientResponseError) Error() string {
	return fmt.Sprintf("remote server returned %d", err.Status)
}

func (err ClientResponseError) ClientError() string {
	return fmt.Sprintf("Remote Server returned %d\n%s", err.Status, err.Message)
}

func (c *client) verifyPeerCertificate(certificates [][]byte, _ [][]*x509.Certificate) error {
	if len(certificates) != 1 {
		return errors.Errorf("tls: invalid certificate chain")
	}

	cert, err := x509.ParseCertificate(certificates[0])
	if err != nil {
		return errors.Wrap(err, "tls: failed to parse certificate")
	}

	// TODO(spheron): return validation back here maybe ?

	certPool := x509.NewCertPool()
	certPool.AddCert(cert)

	opts := x509.VerifyOptions{
		DNSName:                   c.host.Hostname(),
		Roots:                     certPool,
		CurrentTime:               time.Now(),
		KeyUsages:                 []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		MaxConstraintComparisions: 0,
	}

	if _, err = cert.Verify(opts); err != nil {
		return errors.Wrap(err, "tls: unable to verify certificate")
	}

	return nil
}

func (c *client) Status(ctx context.Context) (*provider.Status, error) {
	uri, err := makeURI(c.host, statusPath())
	if err != nil {
		return nil, err
	}
	var obj provider.Status

	if err := c.getStatus(ctx, uri, &obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func (c *client) SubmitManifest(ctx context.Context, dseq uint64, mani manifest.Manifest) error {
	uri, err := makeURI(c.host, submitManifestPath(dseq))
	if err != nil {
		return err
	}

	buf, err := json.Marshal(mani)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", uri, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentTypeJSON)
	resp, err := c.hclient.Do(req)
	if err != nil {
		return err
	}
	responseBuf := &bytes.Buffer{}
	_, err = io.Copy(responseBuf, resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return err
	}

	return createClientResponseErrorIfNotOK(resp, responseBuf)
}

func (c *client) MigrateEndpoints(ctx context.Context, endpoints []string, dseq uint64, gseq uint32) error {
	uri, err := makeURI(c.host, "endpoint/migrate")
	if err != nil {
		return err
	}

	body := endpointMigrateRequestBody{
		EndpointsToMigrate: endpoints,
		DestinationDSeq:    dseq,
		DestinationGSeq:    gseq,
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := c.hclient.Do(req)
	if err != nil {
		return err
	}
	responseBuf := &bytes.Buffer{}
	_, err = io.Copy(responseBuf, resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return err
	}

	return createClientResponseErrorIfNotOK(resp, responseBuf)
}

func (c *client) MigrateHostnames(ctx context.Context, hostnames []string, dseq uint64, gseq uint32) error {
	uri, err := makeURI(c.host, "hostname/migrate")
	if err != nil {
		return err
	}

	body := migrateRequestBody{
		HostnamesToMigrate: hostnames,
		DestinationDSeq:    dseq,
		DestinationGSeq:    gseq,
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := c.hclient.Do(req)
	if err != nil {
		return err
	}
	responseBuf := &bytes.Buffer{}
	_, err = io.Copy(responseBuf, resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return err
	}

	return createClientResponseErrorIfNotOK(resp, responseBuf)
}

func (c *client) LeaseStatus(ctx context.Context, id mtypes.LeaseID) (LeaseStatus, error) {
	uri, err := makeURI(c.host, leaseStatusPath(id))
	if err != nil {
		return LeaseStatus{}, err
	}

	var obj LeaseStatus
	if err := c.getStatus(ctx, uri, &obj); err != nil {
		return LeaseStatus{}, err
	}

	return obj, nil
}

func (c *client) LeaseEvents(ctx context.Context, id mtypes.LeaseID, _ string, follow bool) (*LeaseKubeEvents, error) {
	endpoint, err := url.Parse(c.host.String() + "/" + leaseEventsPath(id))
	if err != nil {
		return nil, err
	}

	switch endpoint.Scheme {
	case schemeWSS, schemeHTTPS:
		endpoint.Scheme = schemeWSS
	default:
		return nil, errors.Errorf("invalid uri scheme %q", endpoint.Scheme)
	}

	query := url.Values{}
	query.Set("follow", strconv.FormatBool(follow))

	endpoint.RawQuery = query.Encode()
	conn, response, err := c.wsclient.DialContext(ctx, endpoint.String(), nil)
	if err != nil {
		if errors.Is(err, websocket.ErrBadHandshake) {
			buf := &bytes.Buffer{}
			_, _ = io.Copy(buf, response.Body)

			return nil, ClientResponseError{
				Status:  response.StatusCode,
				Message: buf.String(),
			}
		}

		return nil, err
	}

	streamch := make(chan cltypes.LeaseEvent)
	onclose := make(chan string, 1)
	logs := &LeaseKubeEvents{
		Stream:  streamch,
		OnClose: onclose,
	}

	processOnCloseErr := func(err error) {
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok { // nolint: gosimple
				onclose <- parseCloseMessage(err.Error())
			} else {
				onclose <- err.Error()
			}
		}
	}

	if err = conn.SetReadDeadline(time.Now().Add(pingWait)); err != nil {
		return nil, err
	}

	conn.SetPingHandler(func(string) error {
		err := conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(time.Second))
		if err != nil {
			return err
		}

		return conn.SetReadDeadline(time.Now().Add(pingWait))
	})

	go func(conn *websocket.Conn) {
		defer func() {
			close(streamch)
			close(onclose)
			_ = conn.Close()
		}()

		for {
			mType, msg, e := conn.ReadMessage()
			if e != nil {
				processOnCloseErr(e)
				return
			}

			switch mType {
			case websocket.TextMessage:
				var evt cltypes.LeaseEvent
				if e = json.Unmarshal(msg, &evt); e != nil {
					onclose <- e.Error()
					return
				}

				streamch <- evt
			case websocket.CloseMessage:
				onclose <- parseCloseMessage(string(msg))
				return
			default:
			}
		}
	}(conn)

	return logs, nil
}

func (c *client) ServiceStatus(ctx context.Context, id mtypes.LeaseID, service string) (*cltypes.ServiceStatus, error) {
	uri, err := makeURI(c.host, serviceStatusPath(id, service))
	if err != nil {
		return nil, err
	}

	var obj cltypes.ServiceStatus
	if err := c.getStatus(ctx, uri, &obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func (c *client) getStatus(ctx context.Context, uri string, obj interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", contentTypeJSON)

	resp, err := c.hclient.Do(req)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	if err != nil {
		return err
	}

	err = createClientResponseErrorIfNotOK(resp, buf)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(buf)
	return dec.Decode(obj)
}

func createClientResponseErrorIfNotOK(resp *http.Response, responseBuf *bytes.Buffer) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	return ClientResponseError{
		Status:  resp.StatusCode,
		Message: responseBuf.String(),
	}
}

// makeURI
// for client queries path must not include owner id
func makeURI(uri *url.URL, path string) (string, error) {
	endpoint, err := url.Parse(uri.String() + "/" + path)
	if err != nil {
		return "", err
	}

	return endpoint.String(), nil
}

func (c *client) LeaseLogs(ctx context.Context,
	id mtypes.LeaseID,
	services string,
	follow bool,
	tailLines int64) (*ServiceLogs, error) {

	endpoint, err := url.Parse(c.host.String() + "/" + serviceLogsPath(id))
	if err != nil {
		return nil, err
	}

	switch endpoint.Scheme {
	case schemeWSS, schemeHTTPS:
		endpoint.Scheme = schemeWSS
	default:
		return nil, errors.Errorf("invalid uri scheme \"%s\"", endpoint.Scheme)
	}

	query := url.Values{}

	query.Set("follow", strconv.FormatBool(follow))

	if services != "" {
		query.Set("services", services)
	}

	endpoint.RawQuery = query.Encode()

	conn, response, err := c.wsclient.DialContext(ctx, endpoint.String(), nil)
	if err != nil {
		if errors.Is(err, websocket.ErrBadHandshake) {
			buf := &bytes.Buffer{}
			_, _ = io.Copy(buf, response.Body)

			return nil, ClientResponseError{
				Status:  response.StatusCode,
				Message: buf.String(),
			}
		}

		return nil, err
	}

	streamch := make(chan ServiceLogMessage)
	onclose := make(chan string, 1)
	logs := &ServiceLogs{
		Stream:  streamch,
		OnClose: onclose,
	}

	if err = conn.SetReadDeadline(time.Now().Add(pingWait)); err != nil {
		return nil, err
	}

	conn.SetPingHandler(func(string) error {
		err := conn.WriteControl(websocket.PongMessage, nil, time.Now().Add(time.Second))
		if err != nil {
			return err
		}

		return conn.SetReadDeadline(time.Now().Add(pingWait))
	})

	go func(conn *websocket.Conn) {
		defer func() {
			close(streamch)
			close(onclose)
			_ = conn.Close()
		}()

		for {
			mType, msg, e := conn.ReadMessage()
			if e != nil {
				onclose <- parseCloseMessage(e.Error())
				return
			}

			switch mType {
			case websocket.TextMessage:
				var logLine ServiceLogMessage
				if e = json.Unmarshal(msg, &logLine); e != nil {
					return
				}

				streamch <- logLine
			case websocket.CloseMessage:
				onclose <- parseCloseMessage(string(msg))
				return
			default:
			}
		}
	}(conn)

	return logs, nil
}

// parseCloseMessage extract close reason from websocket close message
// "websocket: [error code]: [client reason]"
func parseCloseMessage(msg string) string {
	errmsg := strings.SplitN(msg, ": ", 3)
	if len(errmsg) == 3 {
		return errmsg[2]
	}

	return ""
}
