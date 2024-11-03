package domain

import (
	"crypto/ecdsa"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	w := NewWallet()

	assert.NotNil(t, w)
	assert.NotNil(t, w.PrivateKey())
	assert.NotNil(t, w.PublicKey())
}

func TestWalletPrivateKey(t *testing.T) {
	w := NewWallet()

	privateKey := w.PrivateKey()
	assert.NotNil(t, privateKey)
	assert.IsType(t, &ecdsa.PrivateKey{}, privateKey)
}

func TestWalletPublicKey(t *testing.T) {
	w := NewWallet()

	publicKey := w.PublicKey()
	assert.NotNil(t, publicKey)
	assert.IsType(t, &ecdsa.PublicKey{}, publicKey)
}

func TestWalletMarshalJSON(t *testing.T) {
	w := NewWallet()

	jsonBytes, err := json.Marshal(w)
	assert.NoError(t, err)
	assert.NotNil(t, jsonBytes)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "privateKey")
	assert.Contains(t, jsonMap, "publicKey")

	privateKey, ok := jsonMap["privateKey"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, privateKey)

	publicKey, ok := jsonMap["publicKey"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, publicKey)
}
