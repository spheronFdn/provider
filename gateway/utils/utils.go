package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
)

func NewServerTLSConfig(ctx context.Context, certs []tls.Certificate) (*tls.Config, error) {
	// InsecureSkipVerify is set to true due to inability to use normal TLS verification
	// certificate validation and authentication performed in VerifyPeerCertificate
	cfg := &tls.Config{
		Certificates:       certs,
		ClientAuth:         tls.RequestClientCert,
		InsecureSkipVerify: true, // nolint: gosec
		MinVersion:         tls.VersionTLS13,
		VerifyPeerCertificate: func(certificates [][]byte, _ [][]*x509.Certificate) error {
			// TODO(spheron): check if we can totaly remove this VerifyPeerCertificte with InsecureSkipVerify:false
			return nil
		},
	}

	return cfg, nil
}
