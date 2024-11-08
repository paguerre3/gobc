package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	senderPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

	assert.NotNil(t, tx)
	assert.Equal(t, senderPrivateKey, tx.SenderPrivateKey())
	assert.Equal(t, &senderPrivateKey.PublicKey, tx.SenderPublicKey())
	assert.Equal(t, senderAddress, tx.SenderAddress())
	assert.Equal(t, recipientAddress, tx.RecipientAddress())
	assert.Equal(t, amount, tx.Amount())
	assert.NotNil(t, tx.TimeStamp())
}

func TestTransactionGenerateSignature(t *testing.T) {
	senderPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

	signature, err := tx.GenerateSignature()
	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestTransactionMarshalJSON(t *testing.T) {
	senderPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

	jsonBytes, err := json.Marshal(tx)
	assert.NoError(t, err)
	assert.NotNil(t, jsonBytes)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "senderAddress")
	assert.Contains(t, jsonMap, "recipientAddress")
	assert.Contains(t, jsonMap, "amount")
	assert.Contains(t, jsonMap, "timeStamp")

	senderAddressValue, ok := jsonMap["senderAddress"].(string)
	assert.True(t, ok)
	assert.Equal(t, senderAddress, senderAddressValue)

	recipientAddressValue, ok := jsonMap["recipientAddress"].(string)
	assert.True(t, ok)
	assert.Equal(t, recipientAddress, recipientAddressValue)

	amountValue, ok := jsonMap["amount"].(float64)
	assert.True(t, ok)
	assert.Equal(t, amount, amountValue)

	timeStampValue, ok := jsonMap["timeStamp"].(string)
	assert.True(t, ok)
	assert.NotNil(t, timeStampValue)
}
