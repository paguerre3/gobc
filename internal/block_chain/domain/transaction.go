package domain

import (
	"encoding/json"
	"time"
)

type Transaction interface {
	SenderAddress() string
	RecipientAddress() string
	Amount() float64
	TimeStamp() time.Time
}

type transaction struct {
	senderAddress    string
	recipientAddress string
	amount           float64
	timeStamp        time.Time
}

func newTransaction(senderAddress string, receipientAddress string, amount float64, timeStamp *time.Time) Transaction {
	tt := time.Now()
	if timeStamp != nil {
		tt = *timeStamp
	}
	return &transaction{
		senderAddress:    senderAddress,
		recipientAddress: receipientAddress,
		amount:           amount,
		timeStamp:        tt,
	}
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

func (t *transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		// requirement for json to marshal lower cappital fields:
		SenderAddress    string  `json:"senderAddress"`
		RecipientAddress string  `json:"recipientAddress"`
		Amount           float64 `json:"amount"`
		TimeStamp        string  `json:"timeStamp"`
	}{
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Amount:           t.amount,
		TimeStamp:        t.timeStamp.Format(time.RFC3339Nano),
	})
}
