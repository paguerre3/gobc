package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	address := MY_BLOCK_CHAIN_RECEIPT_ADDRESS
	bc := NewBlockchain(address)

	assert.NotNil(t, bc)
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
	assert.Equal(t, address, bc.BlockChainAddressOfRewardReceipient())
}

func TestBlockchainCreateAppendBlock(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	block := bc.CreateAppendBlock(1, [32]byte{})

	assert.NotNil(t, block)
	assert.Len(t, bc.Chain(), 2)
	assert.Empty(t, bc.TransactionPool())
	assert.Equal(t, MY_BLOCK_CHAIN_RECEIPT_ADDRESS, bc.BlockChainAddressOfRewardReceipient())
}

func TestBlockchainLastBlock(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	lastBlock := bc.LastBlock()

	assert.NotNil(t, lastBlock)
}

func TestBlockchainCreateAppendTransaction(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS) // moves genesis transaction from transaction pool to latest block then empty transactyiomn pool
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
	b := bc.Chain()[0]
	assert.Len(t, b.Transactions(), 1)

	transaction := bc.CreateAppendTransaction("sender", "receiver", 10.99)
	assert.NotNil(t, transaction)
	assert.Len(t, bc.TransactionPool(), 1)
}

func TestBlockchainMarshalJSON(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)
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
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS) // moves genesis transaction from transaction pool to latest block then empty transaction pool
	assert.Empty(t, bc.TransactionPool())
	assert.Len(t, bc.Chain(), 1)
	b := bc.Chain()[0]
	assert.Len(t, b.Transactions(), 1)
	tx := b.Transactions()[0]
	assert.Equal(t, GENESSIS_SENDER_ADDRESS, tx.SenderAddress())
	assert.Equal(t, GENESSIS_RECIPIENT_ADDRESS, tx.ReceiverAddress())

	// create 2 new transactions in blockchain pool after it was empty (for a future block):
	t1 := bc.CreateAppendTransaction("sender1", "receiver1", 10.99)
	t2 := bc.CreateAppendTransaction("sender2", "receiver2", 20.99)
	transactionPool := bc.TransactionPool()
	assert.Len(t, transactionPool, 2)
	assert.Equal(t, transactionPool[0], t1)
	assert.Equal(t, transactionPool[1], t2)
}

func TestBlockchainChain(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	bc.CreateAppendBlock(1, [32]byte{})
	bc.CreateAppendBlock(2, [32]byte{})

	chain := bc.Chain()

	assert.Len(t, chain, 3)
}

func TestIsValidProof(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	difficulty := 2 // find only 2 leading zeros for a really fast proof:
	for {
		if bc.IsValidProof(nonce, previousHash, bc.CopyTransactionPool(), difficulty) {
			break
		}
		nonce++
	}
	assert.Greater(t, nonce, 0)
}

func TestProofOfWork(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)
	// default difficulty level is set to "3" so the proof of work should be > 0
	// and it'll be found relatively quickly:
	proof := bc.ProofOfWork() // proof of work returns nonce, i.e. proof == nonce with leading zeros according tio the difficulty
	assert.Greater(t, proof, 0)
}

func TestBlockchainMining(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	// Call the Mining method
	miningSuccess := bc.Mining()

	// Verify that the mining was successful
	assert.True(t, miningSuccess)

	// Verify that a new block was created
	assert.Len(t, bc.Chain(), 2)

	// Verify that the transaction pool is empty
	assert.Empty(t, bc.TransactionPool())

	// Verify that the last block has a valid nonce
	lastBlock := bc.LastBlock()
	assert.Greater(t, lastBlock.Nonce(), 0)

	// Verify that the last block has a valid "previous hash"
	prevHashIndex := len(bc.Chain()) - 2
	assert.Equal(t, bc.Chain()[prevHashIndex].Hash(), lastBlock.PreviousHash())
}

func TestBlockchainCalculateTransactionTotal(t *testing.T) {
	bc := NewBlockchain(MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	// Create some transactions
	bc.CreateAppendTransaction("sender1", "receiver1", 10.99)
	bc.CreateAppendTransaction("sender2", "receiver2", 20.99)
	bc.CreateAppendTransaction("sender3", "receiver1", 30.99)

	assert.True(t, bc.Mining())

	// Calculate the transaction total for "receiver1"
	total := bc.CalculateTransactionTotal("receiver1")

	// Verify that the total is correct
	assert.Equal(t, 10.99+30.99, total)

	// Calculate the transaction total for "sender2"
	total = bc.CalculateTransactionTotal("sender2")

	// Verify that the total is correct
	assert.Equal(t, -20.99, total)
}
