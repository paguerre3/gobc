package domain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Transaction interface {
	Sender() string
	Receiver() string
	Amount() int
}

type Block interface {
	Nonce() int
	PreviousHash() [32]byte
	TimeStamp() time.Time
	Transactions() []string
	Hash() [32]byte
}

type BlockChain interface {
	TransactionPool() []string
	Chain() []Block
	CreateBlock(nonce int, previousHash [32]byte) Block
}

type transaction struct {
	sender   string
	receiver string
	amount   int
}

type block struct {
	nonce        int
	prevHash     [32]byte
	timeStamp    time.Time
	transactions []string
}

type blockChain struct {
	transactionPool []string
	chain           []Block
}

func newTransaction(sender string, receiver string, amount int) Transaction {
	return &transaction{
		sender:   sender,
		receiver: receiver,
		amount:   amount,
	}
}

func (t *transaction) Sender() string {
	return t.sender
}

func (t *transaction) Receiver() string {
	return t.receiver
}

func (t *transaction) Amount() int {
	return t.amount
}

func (t *transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender   string `json:"sender"`
		Receiver string `json:"receiver"`
		Amount   int    `json:"amount"`
	}{
		Sender:   t.Sender(),
		Receiver: t.Receiver(),
		Amount:   t.Amount(),
	})
}

func newBlock(nonce int, previousHash [32]byte) Block {
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

func (b *block) PreviousHash() [32]byte {
	return b.prevHash
}

func (b *block) TimeStamp() time.Time {
	return b.timeStamp
}

func (b *block) Transactions() []string {
	return b.transactions
}

func (b *block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *block) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		Nonce        int      `json:"nonce"`
		PrevHash     [32]byte `json:"prevHash"`
		TimeStamp    string   `json:"timeStamp"`
		Transactions []string `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		TimeStamp:    b.timeStamp.Format(time.RFC3339Nano),
		Transactions: b.transactions,
	})
}

func NewBlockchain() BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	genesisBlock := &block{}
	bc := new(blockChain)
	//bc.transactionPool = []string{"Genesis transaction"}
	bc.CreateBlock(0, genesisBlock.Hash())
	return bc
}

func (bc *blockChain) TransactionPool() []string {
	return bc.transactionPool
}

func (bc *blockChain) Chain() []Block {
	return bc.chain
}

func (bc *blockChain) CreateBlock(nonce int, previousHash [32]byte) Block {
	b := newBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
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
