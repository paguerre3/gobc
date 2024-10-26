package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain()

	assert.NotNil(t, bc)
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
}

func TestBlockchainCreateAppendBlock(t *testing.T) {
	bc := NewBlockchain()

	block := bc.CreateAppendBlock(1, [32]byte{})

	assert.NotNil(t, block)
	assert.Len(t, bc.Chain(), 2)
	assert.Empty(t, bc.TransactionPool())
}

func TestBlockchainLastBlock(t *testing.T) {
	bc := NewBlockchain()

	lastBlock := bc.LastBlock()

	assert.NotNil(t, lastBlock)
}

func TestBlockchainCreateAppendTransaction(t *testing.T) {
	bc := NewBlockchain() // moves genesis transaction from transaction pool to latest block then empty transactyiomn pool
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
	b := bc.Chain()[0]
	assert.Len(t, b.Transactions(), 1)

	transaction := bc.CreateAppendTransaction("sender", "receiver", 10.99)
	assert.NotNil(t, transaction)
	assert.Len(t, bc.TransactionPool(), 1)
}

func TestBlockchainMarshalJSON(t *testing.T) {
	bc := NewBlockchain()
	assert.Empty(t, bc.TransactionPool())

	jsonBytes, err := json.Marshal(bc)
	assert.NoError(t, err)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "transactionPool")
	assert.Contains(t, jsonMap, "chain")
}

func TestBlockchainTransactionPool(t *testing.T) {
	bc := NewBlockchain() // moves genesis transaction from transaction pool to latest block then empty transaction pool
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
	b := bc.Chain()[0]
	assert.Len(t, b.Transactions(), 1)
	tx := b.Transactions()[0]
	assert.Equal(t, GenesisSenderAddress, tx.SenderAddress())
	assert.Equal(t, GenesisRecipientAddress, tx.ReceiverAddress())

	// create 2 new transactions in blockchain pool after it was empty (for a future block):
	t1 := bc.CreateAppendTransaction("sender1", "receiver1", 10.99)
	t2 := bc.CreateAppendTransaction("sender2", "receiver2", 20.99)
	transactionPool := bc.TransactionPool()
	assert.Len(t, transactionPool, 2)
	assert.Equal(t, transactionPool[0], t1)
	assert.Equal(t, transactionPool[1], t2)
}

func TestBlockchainChain(t *testing.T) {
	bc := NewBlockchain()

	bc.CreateAppendBlock(1, [32]byte{})
	bc.CreateAppendBlock(2, [32]byte{})

	chain := bc.Chain()

	assert.Len(t, chain, 3)
}
