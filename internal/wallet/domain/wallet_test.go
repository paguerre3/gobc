package domain

import (
	"crypto/ecdsa"
	"encoding/json"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/stretchr/testify/assert"
)

func TestNewWallet(t *testing.T) {
	w := NewWallet()

	assert.NotNil(t, w)
	assert.NotNil(t, w.PrivateKey())
	assert.NotNil(t, w.PublicKey())
	assert.NotNil(t, w.BlockChainAddress())
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

func isBase58Encoded(s string) bool {
	_, _, err := base58.CheckDecode(s)
	return err == nil
}

func TestWalletBlockChainAddress(t *testing.T) {
	w := NewWallet()

	blockChainAddress := w.BlockChainAddress()
	assert.NotEmpty(t, blockChainAddress)
	assert.IsType(t, "", blockChainAddress)
	assert.True(t, isBase58Encoded(blockChainAddress))
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
	assert.Contains(t, jsonMap, "blockChainAddress")

	privateKey, ok := jsonMap["privateKey"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, privateKey)

	publicKey, ok := jsonMap["publicKey"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, publicKey)

	blockChainAddress, ok := jsonMap["blockChainAddress"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, blockChainAddress)
}
