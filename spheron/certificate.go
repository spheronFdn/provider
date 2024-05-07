package spheron

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"

	"go.step.sm/crypto/pemutil"
)

func (client *Client) ReadX509KeyPair(homeDirectory string, fin ...io.Reader) (*x509.Certificate, tls.Certificate, error) {
	certData, privKeyData, _, err := client.ReadTlsCertificate(homeDirectory, fin...)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	x509cert, err := x509.ParseCertificate(certData)
	if err != nil {
		return nil, tls.Certificate{}, fmt.Errorf("could not parse x509 cert: %w", err)
	}

	result := tls.Certificate{
		Certificate: [][]byte{certData},
	}

	result.PrivateKey, err = x509.ParsePKCS8PrivateKey(privKeyData)
	if err != nil {
		return nil, tls.Certificate{}, fmt.Errorf("%w: failed parsing private key data", err)
	}

	return x509cert, result, err
}

func (client *Client) ReadTlsCertificate(homeDirectory string, fin ...io.Reader) ([]byte, []byte, []byte, error) {
	var pemIn io.Reader
	var closeMe io.ReadCloser

	if len(fin) != 0 {
		if len(fin) != 1 {
			return nil, nil, nil, fmt.Errorf("Read() takes exactly 1 or 0 arguments, not %d", len(fin))
		}
		pemIn = fin[0]
	}

	if pemIn == nil {
		fopen, err := os.OpenFile(homeDirectory+"/"+"provider"+".pem", os.O_RDONLY, 0x0)
		if err != nil {
			return nil, nil, nil, fmt.Errorf("could not open certificate PEM file: %w", err)
		}
		closeMe = fopen
		pemIn = fopen
	}

	cert, privKey, pubKey, err := client.readTlsCertificateImpl(pemIn)

	if closeMe != nil {
		closeErr := closeMe.Close()
		if closeErr != nil {
			return nil, nil, nil, fmt.Errorf("could not close PEM file: %w", closeErr)
		}
	}

	return cert, privKey, pubKey, err
}

func (client *Client) readTlsCertificateImpl(fin io.Reader) ([]byte, []byte, []byte, error) {
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, fin)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed reading certificate PEM file: %w", err)
	}
	data := buf.Bytes()

	// Read certificate
	block, remaining := pem.Decode(data)
	if block == nil {
		return nil, nil, nil, fmt.Errorf("errCertificateNotFoundInPEM")
	}
	cert := block.Bytes

	// Read private key
	block, _ = pem.Decode(remaining)
	if block == nil {
		return nil, nil, nil, fmt.Errorf("errPrivateKeyNotFoundInPEMfmt")
	}

	var privKeyPlaintext []byte
	var privKeyI interface{}
	//TODO(spheron): Replace sig with signature of wallet address (down is example how it's do it with cosmosdk)
	// sig, _, err := cctx.Keyring.SignByAddress(fromAddress, []byte(fromAddress.String())), sig.passwordBytes is used after that
	sig := []byte("mockPassword")

	// PKCS#8 header defined in RFC7468 section 11
	// nolint: gocritic
	if block.Type == "ENCRYPTED PRIVATE KEY" {
		privKeyPlaintext, err = pemutil.DecryptPKCS8PrivateKey(block.Bytes, sig)
	} else if block.Headers["Proc-Type"] == "4,ENCRYPTED" {
		// nolint: staticcheck
		privKeyPlaintext, _ = x509.DecryptPEMBlock(block, sig)

		// DecryptPEMBlock may not return IncorrectPasswordError.
		// Try parse private key instead and if it fails give another try with legacy password
		privKeyI, err = x509.ParsePKCS8PrivateKey(privKeyPlaintext)
	} else {
		return nil, nil, nil, fmt.Errorf("errUnsupportedEncryptedPEM")
	}
	if err != nil {
		return nil, nil, nil, fmt.Errorf("%w: failed decrypting x509 block with private key", err)
	}

	if privKeyI == nil {
		if privKeyI, err = x509.ParsePKCS8PrivateKey(privKeyPlaintext); err != nil {
			return nil, nil, nil, fmt.Errorf("%w: failed parsing private key data", err)
		}
	}

	eckey, valid := privKeyI.(*ecdsa.PrivateKey)
	if !valid {
		return nil, nil, nil, fmt.Errorf("%w: unexpected private key type, expected %T but got %T",
			fmt.Errorf("errPublicKeyNotFoundInPEM"),
			&ecdsa.PrivateKey{},
			privKeyI)
	}

	var pubKey []byte
	if pubKey, err = x509.MarshalPKIXPublicKey(eckey.Public()); err != nil {
		return nil, nil, nil, fmt.Errorf("%w: failed extracting public key", err)
	}

	return cert, privKeyPlaintext, pubKey, nil
}
