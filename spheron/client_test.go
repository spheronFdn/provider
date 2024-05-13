package spheron

import (
	"context"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthTokenValidation(t *testing.T) {

	key := ReadKey("./testdata/wallet1.json", "testPassword")
	assert.NotNil(t, key)

	cctx := Context{
		HomeDir: "./testdata",
		Key:     key,
	}

	authToken, err := CreateAuthorizationToken(context.TODO(), &cctx)
	assert.Nil(t, err)

	testData := []struct {
		tc       string
		positive bool
		wallet   string
	}{
		{"postive auth test", true, "wallet1.json"},
		{"negative auth test", false, "wallet2.json"},
	}

	for i, test := range testData {
		t.Logf("running test case %d, %s", i, test.tc)

		authTokenDecoded, err := base64.StdEncoding.DecodeString(authToken)
		assert.Nil(t, err)
		assert.NotNil(t, authTokenDecoded)

		var authData AuthJson
		err = json.Unmarshal(authTokenDecoded, &authData)
		assert.Nil(t, err)
		var pubKey *ecdsa.PublicKey
		if test.positive {
			pubKey, err = DecodePublicKey(authData.PubKey)
			assert.Nil(t, err)
		} else {
			// Try to pass in a malicious users public key
			testKey := ReadKey("./testdata/"+test.wallet, "testPassword")
			pubKey = &testKey.PrivateKey.PublicKey
		}
		valid, err := ValidateAuthToken(pubKey, authData.SignedTimestamp, strconv.FormatInt(authData.Timestamp, 10))
		assert.Nil(t, err)
		if test.positive {
			assert.Equal(t, valid, true)
		} else {
			assert.Equal(t, valid, false)
		}

	}
}
