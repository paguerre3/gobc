package domain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type Block interface {
	Nonce() int
	PreviousHash() [32]byte
	TimeStamp() time.Time
	Transactions() []Transaction
	Hash() [32]byte
	CreateAppendTransaction(sender string, receiver string, amount float64) Transaction
}

type block struct {
	nonce        int
	prevHash     [32]byte
	timeStamp    time.Time
	transactions []Transaction
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

func (b *block) Transactions() []Transaction {
	return b.transactions
}

func (b *block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *block) CreateAppendTransaction(sender string, receiver string, amount float64) Transaction {
	t := newTransaction(sender, receiver, amount)
	b.transactions = append(b.transactions, t)
	return t
}

func (b *block) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		Nonce        int           `json:"nonce"`
		PrevHash     [32]byte      `json:"prevHash"`
		TimeStamp    string        `json:"timeStamp"`
		Transactions []Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		TimeStamp:    b.timeStamp.Format(time.RFC3339Nano),
		Transactions: b.transactions,
	})
}
