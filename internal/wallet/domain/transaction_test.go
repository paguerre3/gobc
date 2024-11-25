package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateMockKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	return privateKey
}

func TestNewTransaction(t *testing.T) {
	senderPrivateKey := generateMockKey(t)
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx, _ := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

	assert.NotNil(t, tx)
	assert.Equal(t, senderPrivateKey, tx.SenderPrivateKey())
	assert.Equal(t, &senderPrivateKey.PublicKey, tx.SenderPublicKey())
	assert.Equal(t, senderAddress, tx.SenderAddress())
	assert.Equal(t, recipientAddress, tx.RecipientAddress())
	assert.Equal(t, amount, tx.Amount())
	assert.NotNil(t, tx.TimeStamp())
}

func TestTransactionGenerateSignature(t *testing.T) {
	senderPrivateKey := generateMockKey(t)
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx, _ := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

	signature, err := tx.GenerateSignature()
	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestTransactionMarshalJSON(t *testing.T) {
	senderPrivateKey := generateMockKey(t)
	senderAddress := "sender-address"
	recipientAddress := "recipient-address"
	amount := 10.99

	tx, _ := NewTransaction(senderPrivateKey, senderAddress, recipientAddress, amount)

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

func TestNewTransactionAllMissingScenarios(t *testing.T) {
	mockKey := generateMockKey(t)
	validSender := "senderAddress"
	validRecipient := "recipientAddress"
	validAmount := 100.0

	// Test case: all valid inputs
	t.Run("ValidTransaction", func(t *testing.T) {
		tx, err := NewTransaction(mockKey, validSender, validRecipient, validAmount)
		assert.NoError(t, err)
		assert.NotNil(t, tx)
		assert.Equal(t, validSender, tx.SenderAddress())
		assert.Equal(t, validRecipient, tx.RecipientAddress())
		assert.Equal(t, validAmount, tx.Amount())
		assert.NotZero(t, tx.TimeStamp())
	})

	// Test case: missing senderPrivateKey
	t.Run("MissingSenderPrivateKey", func(t *testing.T) {
		_, err := NewTransaction(nil, validSender, validRecipient, validAmount)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "senderPrivateKey is missing")
	})

	// Test case: missing senderAddress
	t.Run("MissingSenderAddress", func(t *testing.T) {
		_, err := NewTransaction(mockKey, "", validRecipient, validAmount)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "senderAddress is missing")
	})

	// Test case: missing recipientAddress
	t.Run("MissingRecipientAddress", func(t *testing.T) {
		_, err := NewTransaction(mockKey, validSender, "", validAmount)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "recipientAddress is missing")
	})

	// Test case: invalid amount
	t.Run("InvalidAmount", func(t *testing.T) {
		_, err := NewTransaction(mockKey, validSender, validRecipient, 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "amount is invaid")
	})

	// Test case: multiple errors
	t.Run("MultipleErrors", func(t *testing.T) {
		_, err := NewTransaction(nil, "", "", 0)
		assert.Error(t, err)
		expectedError := "senderPrivateKey is missing, senderAddress is missing, recipientAddress is missing, amount is invaid"
		assert.Equal(t, expectedError, err.Error())
	})
}
