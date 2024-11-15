package tests

import (
	"encoding/json"
	"fmt"
	"strings"

	"testing"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
	"github.com/paguerre3/blockchain/internal/common"
	wallet_domain "github.com/paguerre3/blockchain/internal/wallet/domain"
	"github.com/stretchr/testify/assert"
)

func newBlockChainWithFmt(checkFunds bool) func() domain.BlockChain {
	bc := domain.NewBlockchain(domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS, checkFunds, common.TEST_SERVER_PORT)
	return func() domain.BlockChain {
		fmt.Println(strings.Repeat("#", 75))
		json, _ := json.MarshalIndent(bc, "", "  ")
		fmt.Println(string(json))
		return bc
	}
}

// calculate and print transaction total
func fmtTransactionTotal(blockChain *domain.BlockChain, senderOrReceipientAddress string) float64 {
	fmt.Println(strings.Repeat("@", 75))
	total := (*blockChain).CalculateTransactionTotal(senderOrReceipientAddress)
	fmt.Printf("%s Transaction total: %.1f\n", senderOrReceipientAddress, total)
	return total
}

func newWalletFmtd() wallet_domain.Wallet {
	w := wallet_domain.NewWallet()
	fmt.Println(strings.Repeat("*", 75))
	json, _ := json.MarshalIndent(w, "", "  ")
	fmt.Println(string(json))
	return w
}

// generate transactuion signature and print
func fmtTransactionSignature(transation *wallet_domain.Transaction) common.Signature {
	fmt.Println(strings.Repeat(">", 75))
	tx := (*transation)
	signature, _ := tx.GenerateSignature()
	json, _ := json.MarshalIndent(tx, "", "  ")
	fmt.Printf("Transaction: %s \nSignature: %s\n", string(json), signature) // it uses signature.String()
	return signature
}

func toggleBlockChainIntegration(checkFunds bool) (float64, float64, float64, float64, float64) {
	// Wallets:
	walletA := newWalletFmtd()
	walletB := newWalletFmtd()
	walletC := newWalletFmtd()
	walletD := newWalletFmtd()

	// Wallet A address is Sender and Wallet B address is Recipient.
	tx1 := wallet_domain.NewTransaction(walletA.PrivateKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)
	tt1 := tx1.TimeStamp()
	signature1 := fmtTransactionSignature(&tx1)

	// Blockchain:
	fmtBlockChain := newBlockChainWithFmt(checkFunds)
	blockChain := fmtBlockChain()
	blockChain.CreateAppendTransaction(walletA.BlockChainAddress(), walletB.BlockChainAddress(), tx1.Amount(), &tt1,
		walletA.PublicKey(), signature1)
	//previousHash := blockChain.LastBlock().Hash()
	//nonce := blockChain.ProofOfWork()
	//blockChain.CreateAppendBlock(nonce, previousHash)
	// Mining already covers calculating the nonce/PoW and creating the Block with previous Hash (previous 3 sentences).
	// In addition, Mining adds an additional transaction to the block with the receiver address for the reward.
	blockChain.Mining() // creates block containing one transaction and one reward transaction
	fmtBlockChain()

	// Wallet B address is Sender and Wallet D address is Recipient.
	tx2 := wallet_domain.NewTransaction(walletB.PrivateKey(), walletB.BlockChainAddress(), walletD.BlockChainAddress(), 2.5)
	tt2 := tx2.TimeStamp()
	signature2 := fmtTransactionSignature(&tx2)
	blockChain.CreateAppendTransaction(walletB.BlockChainAddress(), walletD.BlockChainAddress(), tx2.Amount(), &tt2,
		walletB.PublicKey(), signature2)

	// Wallet C address is Sender and Wallet D address is Recipient.
	tx3 := wallet_domain.NewTransaction(walletC.PrivateKey(), walletC.BlockChainAddress(), walletD.BlockChainAddress(), 5.0)
	tt3 := tx3.TimeStamp()
	signature3 := fmtTransactionSignature(&tx3)
	blockChain.CreateAppendTransaction(walletC.BlockChainAddress(), walletD.BlockChainAddress(), tx3.Amount(), &tt3,
		walletC.PublicKey(), signature3)
	//previousHash = blockChain.LastBlock().Hash()
	//nonce = blockChain.ProofOfWork()
	//blockChain.CreateAppendBlock(nonce, previousHash)
	// Mining already covers calculating the nonce/PoW and creating the Block with previous Hash (previous 3 sentences):
	// In addition, Mining adds an additional transaction to the block with the receiver address for the reward.
	blockChain.Mining() // creates block containing 2 transactions and one reward transaction
	fmtBlockChain()
	t1 := fmtTransactionTotal(&blockChain, walletA.BlockChainAddress())
	t2 := fmtTransactionTotal(&blockChain, walletB.BlockChainAddress())
	t3 := fmtTransactionTotal(&blockChain, walletC.BlockChainAddress())
	t4 := fmtTransactionTotal(&blockChain, walletD.BlockChainAddress())
	t5 := fmtTransactionTotal(&blockChain, domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS) // check rewards
	return t1, t2, t3, t4, t5
}

func TestBlockChainIntegrationWithoutCheckingFunds(t *testing.T) {
	t1, t2, t3, t4, t5 := toggleBlockChainIntegration(false)
	assert.Equal(t, -1.0, t1) // wallet A
	assert.Equal(t, -1.5, t2) // wallet B
	assert.Equal(t, -5.0, t3) // wallet C
	assert.Equal(t, 7.5, t4)  // wallet D
	assert.Equal(t, 2.0, t5)  // rewards
}

func TestBlockChainIntegrationCheckingFunds(t *testing.T) {
	t1, t2, t3, t4, t5 := toggleBlockChainIntegration(true)
	assert.Equal(t, 0.0, t1) // wallet A
	assert.Equal(t, 0.0, t2) // wallet B
	assert.Equal(t, 0.0, t3) // wallet C
	assert.Equal(t, 0.0, t4) // wallet D
	assert.Equal(t, 2.0, t5) // rewards
}
