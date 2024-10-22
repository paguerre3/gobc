package domain

import "time"

type Block interface {
	Nonce() int
	PreviousHash() string
	TimeStamp() time.Time
	Transactions() []string
}

type BlockChain interface {
	TransactionPool() []string
	Chain() []Block
	CreateBlock(nonce int, previousHash string) Block
}

type block struct {
	nonce        int
	prevHash     string
	timeStamp    time.Time
	transactions []string
}

type blockChain struct {
	transactionPool []string
	chain           []Block
}

func newBlock(nonce int, previousHash string) Block {
	return &block{
		nonce:     nonce,
		prevHash:  previousHash,
		timeStamp: time.Now(),
		//transactions: transactions,
	}
}

func (b *block) Nonce() int {
	return b.nonce
}

func (b *block) PreviousHash() string {
	return b.prevHash
}

func (b *block) TimeStamp() time.Time {
	return b.timeStamp
}

func (b *block) Transactions() []string {
	return b.transactions
}

func NewBlockchain() BlockChain {
	bc := new(blockChain)
	//bc.transactionPool = []string{"Genesis transaction"}
	bc.CreateBlock(0, "Genesis block")
	return bc
}

func (bc *blockChain) TransactionPool() []string {
	return bc.transactionPool
}

func (bc *blockChain) Chain() []Block {
	return bc.chain
}

func (bc *blockChain) CreateBlock(nonce int, previousHash string) Block {
	b := newBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
