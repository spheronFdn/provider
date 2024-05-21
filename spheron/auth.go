package spheron

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignMessage(key *keystore.Key, msg string) (string, error) {
	// Convert the message to a hash to be signed
	hash := crypto.Keccak256Hash([]byte(msg)).Bytes()

	// Sign the hash using the private key
	signature, err := crypto.Sign(hash, key.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign message: %v", err)
	}

	// Convert the signature to a hex string
	signatureHex := fmt.Sprintf("0x%x", signature)
	return signatureHex, nil
}

func validateSignature(pubKey *ecdsa.PublicKey, signedMsg string, originalMsg string) (bool, error) {
	publicKeyBytes := crypto.FromECDSAPub(pubKey)
	// Hash the original message
	msgHash := crypto.Keccak256Hash([]byte(originalMsg))

	// Convert the signature hex string to a byte slice
	signatureBytes, err := hex.DecodeString(signedMsg[2:]) // assuming the signature is prefixed with "0x"
	if err != nil {
		return false, fmt.Errorf("invalid signature hex: %v", err)
	}

	sigPublicKey, err := crypto.Ecrecover(msgHash.Bytes(), signatureBytes)
	if err != nil {
		return false, fmt.Errorf("invalid signature format: %v", err)
	}

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)

	return matches, nil
}

func validateTimestamp(originalMsg string) (bool, error) {
	timestamp, err := strconv.ParseInt(originalMsg, 10, 64)
	if err != nil {
		return false, fmt.Errorf("timestamp parsing failed: %v", err)
	}
	currentTime := time.Now().Unix()
	if currentTime-timestamp > 20 {
		return false, fmt.Errorf("timestamp is older than 20 seconds")
	}
	return true, nil
}

func ValidateAuthToken(pubKey *ecdsa.PublicKey, signedMsg string, originalMsg string) (bool, error) {
	validSig, err := validateSignature(pubKey, signedMsg, originalMsg)
	if err != nil {
		return false, err
	}

	validTimestamp, err := validateTimestamp(originalMsg)
	if err != nil {
		return false, err
	}

	return validSig && validTimestamp, nil
}

func EncodePublicKey(pubKey *ecdsa.PublicKey) string {
	if pubKey == nil {
		return ""
	}
	// Convert the public key to a byte slice
	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	// Encode the byte slice to a hex string
	pubKeyHex := hex.EncodeToString(pubKeyBytes)
	return pubKeyHex
}

func DecodePublicKey(pubKeyHex string) (*ecdsa.PublicKey, error) {
	// Decode the hex string to a byte slice
	pubKeyBytes, err := hex.DecodeString(pubKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode public key hex: %v", err)
	}

	// Unmarshal the byte slice to an ECDSA public key
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %v", err)
	}

	return pubKey, nil
}

func CreateAuthorizationToken(ctx context.Context, cctx *Context) (string, error) {
	ts := time.Now().Unix()
	signedTimestamp, err := SignMessage(cctx.Key, strconv.FormatInt(ts, 10))
	publicKeyHex := EncodePublicKey(&cctx.Key.PrivateKey.PublicKey)

	if err != nil {
		return "", err
	}

	body := AuthJson{
		Timestamp:       ts,
		PubKey:          publicKeyHex,
		SignedTimestamp: signedTimestamp,
	}
	// Convert authToken to a base64-encoded string
	authTokenBytes, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("unable to marshal auth token: %v", err.Error())
	}
	res := base64.StdEncoding.EncodeToString(authTokenBytes)
	return res, nil
}
