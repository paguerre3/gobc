package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignatureMarshalJSON(t *testing.T) {
	s := NewSignature(big.NewInt(1), big.NewInt(2))
	jsonBytes, err := json.Marshal(s)
	assert.NoError(t, err)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)
	assert.Contains(t, jsonMap, "r")
	assert.Contains(t, jsonMap, "s")
}

func TestSignatureUnmarshalJSON(t *testing.T) {
	s := NewSignature(big.NewInt(1), big.NewInt(2))
	jsonBytes, err := json.Marshal(s)
	assert.NoError(t, err)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)
	assert.Contains(t, jsonMap, "r")
	assert.Contains(t, jsonMap, "s")
}

func TestSignatureString(t *testing.T) {
	s := signature{r: big.NewInt(1), s: big.NewInt(2)}
	spaces := strings.Repeat(" ", 63)
	assert.Equal(t, fmt.Sprintf("%s%s%s%s", spaces, "1", spaces, "2"), s.String())
}

func TestStringToBigIntTuple(t *testing.T) {
	r := big.NewInt(12345)
	s := big.NewInt(67890)
	input := fmt.Sprintf("%064x%064x", r, s)

	bix, biy := stringToBigIntTuple(input)
	assert.Equal(t, r, &bix, "BigInt r mismatch")
	assert.Equal(t, s, &biy, "BigInt s mismatch")
}

func TestPublicKeyFromString(t *testing.T) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "Failed to generate key")
	pubKeyHex := fmt.Sprintf("%064x%064x", priv.PublicKey.X, priv.PublicKey.Y)

	pubKey := PublicKeyFromString(pubKeyHex)
	require.NotNil(t, pubKey, "Public key should not be nil")
	assert.Equal(t, priv.PublicKey.X, pubKey.X, "Public key X mismatch")
	assert.Equal(t, priv.PublicKey.Y, pubKey.Y, "Public key Y mismatch")
}

func TestPrivateKeyFromString(t *testing.T) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "Failed to generate key")
	privKeyHex := fmt.Sprintf("%064x", priv.D)
	pubKey := &priv.PublicKey

	privateKey := PrivateKeyFromString(privKeyHex, pubKey)
	require.NotNil(t, privateKey, "Private key should not be nil")
	assert.Equal(t, priv.D, privateKey.D, "Private key D mismatch")
}

func TestSignatureFromString(t *testing.T) {
	r := big.NewInt(12345)
	s := big.NewInt(67890)
	input := fmt.Sprintf("%064x%064x", r, s)

	sig := SignatureFromString(input)
	require.NotNil(t, sig, "Signature should not be nil")
	assert.Equal(t, r, sig.R(), "Signature R mismatch")
	assert.Equal(t, s, sig.S(), "Signature S mismatch")
}
