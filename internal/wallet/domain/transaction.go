package domain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"time"

	"github.com/paguerre3/blockchain/internal/common"
)

type Transaction interface {
	SenderPrivateKey() *ecdsa.PrivateKey
	SenderPublicKey() *ecdsa.PublicKey
	SenderAddress() string
	RecipientAddress() string
	Amount() float64
	TimeStamp() time.Time

	GenerateSignature() (common.Signature, error)
}

type transaction struct {
	senderPrivateKey *ecdsa.PrivateKey
	senderPublicKey  *ecdsa.PublicKey
	senderAddress    string
	recipientAddress string
	amount           float64
	timeStamp        time.Time
}

func NewTransaction(senderPrivateKey *ecdsa.PrivateKey, senderAddress string, recipientAddress string, amount float64) Transaction {
	return &transaction{
		senderPrivateKey: senderPrivateKey,
		senderPublicKey:  &senderPrivateKey.PublicKey,
		senderAddress:    senderAddress,
		recipientAddress: recipientAddress,
		amount:           amount,
		timeStamp:        time.Now(),
	}
}

func (t *transaction) SenderPrivateKey() *ecdsa.PrivateKey {
	return t.senderPrivateKey
}

func (t *transaction) SenderPublicKey() *ecdsa.PublicKey {
	return t.senderPublicKey
}

func (t *transaction) SenderAddress() string {
	return t.senderAddress
}

func (t *transaction) RecipientAddress() string {
	return t.recipientAddress
}

func (t *transaction) Amount() float64 {
	return t.amount
}

func (t *transaction) TimeStamp() time.Time {
	return t.timeStamp
}

func (t *transaction) GenerateSignature() (common.Signature, error) {
	m, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	// hash
	h := sha256.Sum256([]byte(m))
	// sign hash of json
	r, s, err := ecdsa.Sign(rand.Reader, t.senderPrivateKey, h[:])
	return common.NewSignature(r, s), err
}

func (t *transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		// requirement for json to marshal lower cappital fields:
		//SenderPrivateKey string  `json:"senderPrivateKey"`
		//SenderPublicKey  string  `json:"senderPublicKey"`
		SenderAddress    string  `json:"senderAddress"`
		RecipientAddress string  `json:"recipientAddress"`
		Amount           float64 `json:"amount"`
		TimeStamp        string  `json:"timeStamp"`
	}{
		//Avoid using keys as they will be used by the signature of the transaction:
		//SenderPrivateKey: fmt.Sprintf("%x", t.senderPrivateKey.D.Bytes()), // hex.EncodeToString(hash[:]) OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f
		//SenderPublicKey:  fmt.Sprintf("%x%x", t.senderPublicKey.X.Bytes(), t.senderPublicKey.Y.Bytes()),
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Amount:           t.amount,
		TimeStamp:        t.timeStamp.Format(time.RFC3339Nano),
	})
}