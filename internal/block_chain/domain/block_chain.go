package domain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"time"

	"github.com/paguerre3/blockchain/internal/common"
)

const (
	GENESSIS_SENDER_ADDRESS        = "genesis_sender_address"
	GENESSIS_RECIPIENT_ADDRESS     = "genesis_recipient_address"
	MINING_DIFFICULTY              = 3                                      // increasing difficulty means more time for guessing Nonce, e.g. 4 is arround 10 minutes or more
	MINING_SENDER_ADDRESS          = "THE_BLOCKCHAIN_MINING_SENDER_ADDRESS" // block chain mining server address that "sends" rewards
	MINING_REWARD                  = 1.0
	MY_BLOCK_CHAIN_RECEIPT_ADDRESS = "MY_BLOCKCHAIN_RECEIPT_ADDRESS_TO_OBTAIN_MINING_REWARD" // address for receiving mining rewards
)

type BlockChain interface {
	TransactionPool() []Transaction
	Chain() []Block
	BlockChainAddressOfRewardRecipient() string

	CreateAppendBlock(nonce int, previousHash [32]byte) *Block
	LastBlock() Block
	CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64) Transaction
	CopyTransactionPool() []Transaction
	IsValidProof(nonce int, previousHash [32]byte, transactions []Transaction, difficulty int) bool
	ProofOfWork() int
	Mining() bool
	CalculateTransactionTotal(blockChainAddressOfReceipientOrSender string) float64

	VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, signature common.Signature, transaction Transaction) bool
}

type blockChain struct {
	transactionPool                    []Transaction
	chain                              []Block
	blockChainAddressOfRewardRecipient string // server address registered to "receive" rewards of succesffull mining (the 1st sending the right PoW)
}

func NewBlockchain(blockChainAddressOfRewardRecipient string) BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	emptyBlock := &block{}
	bc := new(blockChain)
	bc.blockChainAddressOfRewardRecipient = blockChainAddressOfRewardRecipient
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

func (bc *blockChain) BlockChainAddressOfRewardRecipient() string {
	return bc.blockChainAddressOfRewardRecipient
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
		clonedTransactions[i] = newTransaction(t.SenderAddress(), t.RecipientAddress(), t.Amount())
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
	//fmt.Printf("guessHashBlockStr: %s\n", guessHashBlockStr) // uncomment to see the hash with leading zeros

	// From "0 index" to "difficulty index" exclusively, i.e. it starts with "zeros" up to "difficulty" exclusively,
	// e.g. hash starts with "000" (leading zeros).
	// Note this is a quick way to compare hash instead to check the entire hash.
	return guessHashBlockStr[:difficulty] == zeros
}

func (bc *blockChain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.IsValidProof(nonce, previousHash, transactions, MINING_DIFFICULTY) {
		nonce++
	}
	return nonce
}

func (bc *blockChain) Mining() bool {
	// The blockChainn sender sends rewards to the blockChain address because of successfull mining:
	bc.CreateAppendTransaction(MINING_SENDER_ADDRESS, bc.BlockChainAddressOfRewardRecipient(), MINING_REWARD)
	nonce := bc.ProofOfWork()
	var b *Block = nil
	if nonce > 0 {
		previousHash := bc.LastBlock().Hash()
		b = bc.CreateAppendBlock(nonce, previousHash)
	}
	return b != nil
}

func (bc *blockChain) CalculateTransactionTotal(blockChainAddressOfRecipientOrSender string) float64 {
	var total float64 = 0
	for _, b := range bc.Chain() { // iterates the entire block chain
		for _, t := range b.Transactions() {
			if t.SenderAddress() == blockChainAddressOfRecipientOrSender {
				total -= t.Amount()
			}
			if t.RecipientAddress() == blockChainAddressOfRecipientOrSender {
				total += t.Amount()
			}
		}
	}
	return total
}

func (bc *blockChain) VerifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, signature common.Signature, transaction Transaction) bool {
	m, err := json.Marshal(transaction)
	if err != nil {
		return false
	}
	hash := sha256.Sum256([]byte(m))
	return ecdsa.Verify(senderPublicKey, hash[:], signature.R(), signature.S())
}

func (bc *blockChain) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		TransactionPool                    []Transaction `json:"transactionPool"`
		Chain                              []Block       `json:"chain"`
		BlockChainAddressOfRewardRecipient string        `json:"blockChainAddressOfRewardRecipient"`
	}{
		TransactionPool:                    bc.transactionPool,
		Chain:                              bc.chain,
		BlockChainAddressOfRewardRecipient: bc.blockChainAddressOfRewardRecipient,
	})
}
