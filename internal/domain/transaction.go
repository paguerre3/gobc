package domain

import (
	"encoding/json"
	"time"
)

type Transaction interface {
	Sender() string
	Receiver() string
	Amount() float64
	TimeStamp() time.Time
}

type transaction struct {
	sender    string
	receiver  string
	amount    float64
	timeStamp time.Time
}

func newTransaction(sender string, receiver string, amount float64) Transaction {
	return &transaction{
		sender:    sender,
		receiver:  receiver,
		amount:    amount,
		timeStamp: time.Now(),
	}
}

func (t *transaction) Sender() string {
	return t.sender
}

func (t *transaction) Receiver() string {
	return t.receiver
}

func (t *transaction) Amount() float64 {
	return t.amount
}

func (t *transaction) TimeStamp() time.Time {
	return t.timeStamp
}

func (t *transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender"`
		Receiver  string  `json:"receiver"`
		Amount    float64 `json:"amount"`
		TimeStamp string  `json:"timeStamp"`
	}{
		Sender:    t.Sender(),
		Receiver:  t.Receiver(),
		Amount:    t.Amount(),
		TimeStamp: t.timeStamp.Format(time.RFC3339Nano),
	})
}
