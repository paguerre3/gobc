package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	nonce := 1
	previousHash := [32]byte{}
	transactions := []Transaction{}

	block := newBlock(nonce, previousHash, transactions)

	assert.NotNil(t, block)
	assert.Equal(t, nonce, block.Nonce())
	assert.Equal(t, previousHash, block.PreviousHash())
	assert.NotNil(t, block.TimeStamp())
	assert.Empty(t, block.Transactions())
}

func TestBlockHash(t *testing.T) {
	nonce := 1
	previousHash := [32]byte{}
	transactions := []Transaction{}

	block := newBlock(nonce, previousHash, transactions)

	hash := block.Hash()

	assert.NotNil(t, hash)
}

func TestBlockMarshalJSON(t *testing.T) {
	nonce := 1
	previousHash := [32]byte{}
	transactions := []Transaction{}

	block := newBlock(nonce, previousHash, transactions)

	jsonBytes, err := json.Marshal(block)
	assert.NoError(t, err)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "nonce")
	assert.Contains(t, jsonMap, "prevHash")
	assert.Contains(t, jsonMap, "timeStamp")
	assert.Contains(t, jsonMap, "transactions")
}

func TestBlockHashCollision(t *testing.T) {
	nonce := 1
	previousHash := [32]byte{}
	transactions := []Transaction{}

	block1 := newBlock(nonce, previousHash, transactions)
	block2 := newBlock(nonce, previousHash, transactions)

	hash1 := block1.Hash()
	hash2 := block2.Hash()

	assert.NotEqual(t, hash1, hash2)
}
