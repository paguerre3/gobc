package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
	"github.com/paguerre3/blockchain/internal/common"
	wallet_domain "github.com/paguerre3/blockchain/internal/wallet/domain"
)

func newBlockChainWithFmt() func() domain.BlockChain {
	bc := domain.NewBlockchain(domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS, false)
	return func() domain.BlockChain {
		fmt.Println(strings.Repeat("#", 75))
		json, _ := json.MarshalIndent(bc, "", "  ")
		fmt.Println(string(json))
		return bc
	}
}

// calculate and print transaction total
func fmtTransactionTotal(blockChain *domain.BlockChain, senderOrReceipientAddress string) {
	fmt.Println(strings.Repeat("@", 75))
	fmt.Printf("%s Transaction total: %.1f\n", senderOrReceipientAddress, (*blockChain).CalculateTransactionTotal(senderOrReceipientAddress))
}

func newWalletWithfmt() func() wallet_domain.Wallet {
	w := wallet_domain.NewWallet()
	return func() wallet_domain.Wallet {
		fmt.Println(strings.Repeat("*", 75))
		json, _ := json.MarshalIndent(w, "", "  ")
		fmt.Println(string(json))
		return w
	}
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

func main() {
	// Wallets:
	walletA := newWalletWithfmt()()
	walletB := newWalletWithfmt()()
	walletC := newWalletWithfmt()()
	walletD := newWalletWithfmt()()

	// Wallet A address is Sender and Wallet B address is Recipient.
	tx1 := wallet_domain.NewTransaction(walletA.PrivateKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)
	tt1 := tx1.TimeStamp()
	signature1 := fmtTransactionSignature(&tx1)

	// Blockchain:
	fmtBlockChain := newBlockChainWithFmt()
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
	fmtTransactionTotal(&blockChain, walletA.BlockChainAddress())
	fmtTransactionTotal(&blockChain, walletB.BlockChainAddress())
	fmtTransactionTotal(&blockChain, walletC.BlockChainAddress())
	fmtTransactionTotal(&blockChain, walletD.BlockChainAddress())
	fmtTransactionTotal(&blockChain, domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS) // check rewards
}
