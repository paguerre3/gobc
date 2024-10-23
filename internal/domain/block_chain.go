package domain

import (
	"encoding/json"
)

type BlockChain interface {
	TransactionPool() []string
	Chain() []Block
	CreateAppendBlock(nonce int, previousHash [32]byte) *Block
	LastBlock() Block
}

type blockChain struct {
	transactionPool []string
	chain           []Block
}

func NewBlockchain() BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	emptyBlock := &block{}
	bc := new(blockChain)
	//bc.transactionPool = []string{"Genesis transaction pool"}
	genesisBlock := bc.CreateAppendBlock(1, emptyBlock.Hash())
	(*genesisBlock).CreateAppendTransaction("Genesis Sender", "Genesis Receiver", 0)
	return bc
}

func (bc *blockChain) TransactionPool() []string {
	return bc.transactionPool
}

func (bc *blockChain) Chain() []Block {
	return bc.chain
}

func (bc *blockChain) CreateAppendBlock(nonce int, previousHash [32]byte) *Block {
	b := newBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return &b
}

func (bc *blockChain) LastBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *blockChain) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		TransactionPool []string `json:"transactionPool"`
		Chain           []Block  `json:"chain"`
	}{
		TransactionPool: bc.transactionPool,
		Chain:           bc.chain,
	})
}
