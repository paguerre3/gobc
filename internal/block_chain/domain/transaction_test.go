package domain

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	senderAddress := "sender"
	receiverAddress := "receiver"
	amount := 10.99

	tt := time.Now().Add(13 * time.Hour)
	tx := newTransaction(senderAddress, receiverAddress, amount, &tt)

	assert.NotNil(t, tx)
	assert.Equal(t, senderAddress, tx.SenderAddress())
	assert.Equal(t, receiverAddress, tx.RecipientAddress())
	assert.Equal(t, amount, tx.Amount())
	assert.NotNil(t, tx.TimeStamp())
	assert.Equal(t, tt, tx.TimeStamp())
}

func TestTransactionMarshalJSON(t *testing.T) {
	tx := newTransaction("sender", "receiver", 10.99, nil)

	jsonBytes, err := json.Marshal(tx)
	assert.NoError(t, err)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)

	expectedKeys := []string{"senderAddress", "recipientAddress", "amount", "timeStamp"}
	for _, key := range expectedKeys {
		assert.Contains(t, jsonMap, key)
	}
}
