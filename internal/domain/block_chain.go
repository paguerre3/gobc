package domain

import "time"

type block struct {
	nonce        int
	prevHash     string
	timeStamp    time.Time
	transactions []string
}

type blockChain struct {
	transactionPool []string
	chain           []*block
}

func newBlock(nonce int, prevHash string) *block {
	return &block{
		nonce:     nonce,
		prevHash:  prevHash,
		timeStamp: time.Now(),
		//transactions: transactions,
	}
}
