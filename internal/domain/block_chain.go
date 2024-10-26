package domain

import (
	"encoding/json"
)

const (
	GenesisSenderAddress    = "genesis_sender_address"
	GenesisRecipientAddress = "genesis_recipient_address"
)

type BlockChain interface {
	TransactionPool() []Transaction
	Chain() []Block

	CreateAppendBlock(nonce int, previousHash [32]byte) *Block
	LastBlock() Block
	CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64) Transaction
}

type blockChain struct {
	transactionPool []Transaction
	chain           []Block
}

func NewBlockchain() BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	emptyBlock := &block{}
	bc := new(blockChain)
	// add genesis transactions to blockchain Pool:
	bc.CreateAppendTransaction(GenesisSenderAddress, GenesisRecipientAddress, 0)
	bc.CreateAppendBlock(1, emptyBlock.Hash()) // transfer transacton "pool" from blockhain to new block and empty it
	return bc
}

func (bc *blockChain) TransactionPool() []Transaction {
	return bc.transactionPool
}

func (bc *blockChain) Chain() []Block {
	return bc.chain
}

func (bc *blockChain) CreateAppendBlock(nonce int, previousHash [32]byte) *Block {
	// 1. Create new block and transfer transacion pool of blockchain to the new block:
	b := newBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	// 2. Empty transaction pool of blockchain:
	bc.transactionPool = []Transaction{}
	return &b
}

func (bc *blockChain) LastBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *blockChain) CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64) Transaction {
	t := newTransaction(senderAddress, receiverAddress, amount)
	bc.transactionPool = append(bc.transactionPool, t)
	return t
}

func (bc *blockChain) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		TransactionPool []Transaction `json:"transactionPool"`
		Chain           []Block       `json:"chain"`
	}{
		TransactionPool: bc.transactionPool,
		Chain:           bc.chain,
	})
}
