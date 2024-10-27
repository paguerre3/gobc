package domain

import (
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"
)

const (
	GENESSIS_SENDER_ADDRESS    = "genesis_sender_address"
	GENESSIS_RECIPIENT_ADDRESS = "genesis_recipient_address"
	MINING_DIFFICULTY          = 3 // increasing difficulty means more time for guessing Nonce, e.g. 4 is arround 10 minutes or more
)

type BlockChain interface {
	TransactionPool() []Transaction
	Chain() []Block

	CreateAppendBlock(nonce int, previousHash [32]byte) *Block
	LastBlock() Block
	CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64) Transaction
	CopyTransactionPool() []Transaction
	IsValidProof(nonce int, previousHash [32]byte, transactions []Transaction, difficulty int) bool
}

type blockChain struct {
	transactionPool []Transaction
	chain           []Block
}

func NewBlockchain() BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	emptyBlock := &block{}
	bc := new(blockChain)
	// add genesis transactions to blockchain Pool:
	bc.CreateAppendTransaction(GENESSIS_SENDER_ADDRESS, GENESSIS_RECIPIENT_ADDRESS, 0)
	bc.CreateAppendBlock(0, emptyBlock.Hash()) // transfer transacton "pool" from blockhain to new block and empty it
	return bc
}

func (bc *blockChain) TransactionPool() []Transaction {
	return bc.transactionPool
}

func (bc *blockChain) Chain() []Block {
	return bc.chain
}

func (bc *blockChain) CreateAppendBlock(nonce int, previousHash [32]byte) *Block {
	// 1. Create new block and transfer transacion pool of blockchain to the new block:
	b := newBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	// 2. Empty transaction pool of blockchain:
	bc.transactionPool = []Transaction{}
	return &b
}

func (bc *blockChain) LastBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *blockChain) CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64) Transaction {
	t := newTransaction(senderAddress, receiverAddress, amount)
	bc.transactionPool = append(bc.transactionPool, t)
	return t
}

func (bc *blockChain) CopyTransactionPool() []Transaction {
	clonedTransactions := make([]Transaction, len(bc.TransactionPool()))
	for i, t := range bc.TransactionPool() {
		clonedTransactions[i] = newTransaction(t.SenderAddress(), t.ReceiverAddress(), t.Amount())
	}
	return clonedTransactions
}

func (bc *blockChain) IsValidProof(nonce int, previousHash [32]byte, transactions []Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	var guessBlock Block = &block{
		nonce:        nonce,
		prevHash:     previousHash,
		timeStamp:    time.Time{}, // 0001-01-01 00:00:00 +0000 UTC ==> t.IsZero() == true
		transactions: transactions,
	}
	hash := guessBlock.Hash()
	guessHashBlockStr := hex.EncodeToString(hash[:]) // OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f

	// From "0 index" to "difficulty index" exclusively, i.e. it starts with "zeros" up to "difficulty" exclusively,
	// e.g. hash starts with "000" (leading zeros).
	// Note this is a quick way to compare hash instead to check the entire hash.
	return guessHashBlockStr[:difficulty] == zeros
}

func (bc *blockChain) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		TransactionPool []Transaction `json:"transactionPool"`
		Chain           []Block       `json:"chain"`
	}{
		TransactionPool: bc.transactionPool,
		Chain:           bc.chain,
	})
}
