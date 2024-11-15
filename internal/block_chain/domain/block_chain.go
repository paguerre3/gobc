package domain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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
	TEST_SERVER_PORT               = ":0000"
)

type BlockChain interface {
	TransactionPool() []Transaction
	Chain() []Block
	BlockChainAddressOfRewardRecipient() string
	CheckFunds() bool
	ServerPort() string

	CreateAppendBlock(nonce int, previousHash [32]byte) *Block
	LastBlock() Block
	CreateAppendTransaction(senderAddress string, receiverAddress string, amount float64, timeStamp *time.Time,
		senderPublicKey *ecdsa.PublicKey, signature common.Signature) (Transaction, error)
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
	checkFunds                         bool
	serverPort                         string
}

func NewBlockchain(blockChainAddressOfRewardRecipient string, checkFunds bool, serverPort string) BlockChain {
	// only hash of empty block is stored at the beginning (using default fields):
	emptyBlock := &block{}
	bc := new(blockChain)
	bc.blockChainAddressOfRewardRecipient = blockChainAddressOfRewardRecipient
	bc.checkFunds = checkFunds
	bc.serverPort = serverPort
	// add genesis transactions to blockchain Pool
	// (the is no need of passing public key and signature for the genesis transaction scenario):
	bc.CreateAppendTransaction(GENESSIS_SENDER_ADDRESS, GENESSIS_RECIPIENT_ADDRESS, 0, nil, nil, nil)
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

func (bc *blockChain) CheckFunds() bool {
	return bc.checkFunds
}

func (bc *blockChain) ServerPort() string {
	return bc.serverPort
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

func (bc *blockChain) CreateAppendTransaction(senderAddress string, recipientAddress string, amount float64, timeStamp *time.Time,
	senderPublicKey *ecdsa.PublicKey, signature common.Signature) (Transaction, error) {
	t := newTransaction(senderAddress, recipientAddress, amount, timeStamp)
	if senderAddress == MINING_SENDER_ADDRESS || senderAddress == GENESSIS_SENDER_ADDRESS {
		// this is a transaction for reward so there is no need to verify the signature
		// (also, there is no need to verify a genesis transaction case)
		bc.transactionPool = append(bc.transactionPool, t)
		return t, nil
	}
	if bc.VerifyTransactionSignature(senderPublicKey, signature, t) {

		// Real case scenario, a wallet needs to have funds for sending cryptos
		// received via Mining rewards or transactions from other wallets before sending criptos!
		if bc.CheckFunds() && bc.CalculateTransactionTotal(senderAddress) < amount {
			return nil, fmt.Errorf("insufficient funds")
		}
		bc.transactionPool = append(bc.transactionPool, t)
		return t, nil
	}
	return nil, fmt.Errorf("invalid signature")
}

func (bc *blockChain) CopyTransactionPool() []Transaction {
	clonedTransactions := make([]Transaction, len(bc.TransactionPool()))
	for i, t := range bc.TransactionPool() {
		tt := t.TimeStamp()
		clonedTransactions[i] = newTransaction(t.SenderAddress(), t.RecipientAddress(), t.Amount(), &tt)
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
	// The blockChainn sender sends rewards to the blockChain address because of successfull mining
	// (no nee to pass public key and signature for "rewards" scenario):
	bc.CreateAppendTransaction(MINING_SENDER_ADDRESS, bc.BlockChainAddressOfRewardRecipient(), MINING_REWARD, nil, nil, nil)
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
	if senderPublicKey == nil || signature == nil {
		return false
	}
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
		CheckFunds                         bool          `json:"checkFunds"`
		ServerPort                         string        `json:"serverPort"`
	}{
		TransactionPool:                    bc.transactionPool,
		Chain:                              bc.chain,
		BlockChainAddressOfRewardRecipient: bc.blockChainAddressOfRewardRecipient,
		CheckFunds:                         bc.checkFunds,
		ServerPort:                         bc.serverPort,
	})
}
