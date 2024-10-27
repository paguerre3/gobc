package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block interface {
	Nonce() int
	PreviousHash() [32]byte
	TimeStamp() time.Time
	Transactions() []Transaction

	Hash() [32]byte
}

type block struct {
	nonce        int
	prevHash     [32]byte
	timeStamp    time.Time
	transactions []Transaction
}

func newBlock(nonce int, previousHash [32]byte, transactions []Transaction) Block {
	return &block{
		nonce:        nonce,
		prevHash:     previousHash,
		timeStamp:    time.Now(),
		transactions: transactions,
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

func (b *block) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		Nonce        int           `json:"nonce"`
		PrevHash     string        `json:"prevHash"`
		TimeStamp    string        `json:"timeStamp"`
		Transactions []Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PrevHash:     hex.EncodeToString(b.prevHash[:]),
		TimeStamp:    b.timeStamp.Format(time.RFC3339Nano),
		Transactions: b.transactions,
	})
}
