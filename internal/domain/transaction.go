package domain

import (
	"encoding/json"
	"time"
)

type Transaction interface {
	SenderAddress() string
	ReceiverAddress() string
	Amount() float64
	TimeStamp() time.Time
}

type transaction struct {
	senderAddress   string
	receiverAddress string
	amount          float64
	timeStamp       time.Time
}

func newTransaction(senderAddress string, receiverAddress string, amount float64) Transaction {
	return &transaction{
		senderAddress:   senderAddress,
		receiverAddress: receiverAddress,
		amount:          amount,
		timeStamp:       time.Now(),
	}
}

func (t *transaction) SenderAddress() string {
	return t.senderAddress
}

func (t *transaction) ReceiverAddress() string {
	return t.receiverAddress
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
		SenderAddress   string  `json:"senderAddress"`
		ReceiverAddress string  `json:"receiverAddress"`
		Amount          float64 `json:"amount"`
		TimeStamp       string  `json:"timeStamp"`
	}{
		SenderAddress:   t.senderAddress,
		ReceiverAddress: t.receiverAddress,
		Amount:          t.amount,
		TimeStamp:       t.timeStamp.Format(time.RFC3339Nano),
	})
}
