package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
	"github.com/paguerre3/blockchain/internal/common"
	walletd "github.com/paguerre3/blockchain/internal/wallet/domain"
)

func newBlockChainWithFmt() func() domain.BlockChain {
	bc := domain.NewBlockchain(domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS)
	return func() domain.BlockChain {
		fmt.Println(strings.Repeat("#", 75))
		json, _ := json.MarshalIndent(bc, "", "  ")
		fmt.Println(string(json))
		return bc
	}
}

func fmtTransactionTotal(blockChain *domain.BlockChain, senderOrReceipientAddress string) {
	fmt.Printf("%s Transaction total: %.1f\n", senderOrReceipientAddress, (*blockChain).CalculateTransactionTotal(senderOrReceipientAddress))
}

func newWalletWithfmt() func() walletd.Wallet {
	w := walletd.NewWallet()
	return func() walletd.Wallet {
		fmt.Println(strings.Repeat("*", 75))
		json, _ := json.MarshalIndent(w, "", "  ")
		fmt.Println(string(json))
		return w
	}
}

func fmtTransactionSignature(transation *walletd.Transaction) common.Signature {
	signature, _ := (*transation).GenerateSignature()
	fmt.Printf("Transaction Signature: %s\n", signature) // it uses signature.String()
	return signature
}

func main() {
	blockChain.CreateAppendTransaction("sender_address_2", "receiver_address_2", 2.5)
	blockChain.CreateAppendTransaction("sender_address_3", "receiver_address_2", 5.0)
	//previousHash = blockChain.LastBlock().Hash()
	//nonce = blockChain.ProofOfWork()
	//blockChain.CreateAppendBlock(nonce, previousHash)
	// Mining already covers calculating the nonce/PoW and creating the Block with previous Hash (previous 3 sentences):
	// In addition, Mining adds an additional transaction to the block with the receiver address for the reward.
	blockChain.Mining()
	fmtBlockChain()

	fmtTransactionTotal(&blockChain, "sender_address_1")
	fmtTransactionTotal(&blockChain, "receiver_address_1")
	fmtTransactionTotal(&blockChain, "sender_address_2")
	fmtTransactionTotal(&blockChain, "sender_address_3")
	fmtTransactionTotal(&blockChain, "receiver_address_2")
	fmtTransactionTotal(&blockChain, domain.MY_BLOCK_CHAIN_RECEIPT_ADDRESS)

	// Wallets:
	fmtWalletA := newWalletWithfmt()
	fmtWalletB := newWalletWithfmt()
	fmtWalletC := newWalletWithfmt()
	walletA := fmtWalletA()
	walletB := fmtWalletB()
	walletC := fmtWalletC()

	// Wallet A address is Sender and Wallet B address is Recipient.
	tx1 := walletd.NewTransaction(walletA.PrivateKey(), walletA.BlockChainAddress(), walletB.BlockChainAddress(), 1.0)
	signature1 := fmtTransactionSignature(&tx1)

	// Blockchain:
	fmtBlockChain := newBlockChainWithFmt()
	blockChain := fmtBlockChain()
	_, err := blockChain.CreateAppendTransaction(walletA.BlockChainAddress(), walletB.BlockChainAddress(), tx1.Amount(), walletA.PublicKey(), signature1)
	if err != nil {
		panic(err)
	}
	//previousHash := blockChain.LastBlock().Hash()
	//nonce := blockChain.ProofOfWork()
	//blockChain.CreateAppendBlock(nonce, previousHash)
	// Mining already covers calculating the nonce/PoW and creating the Block with previous Hash (previous 3 sentences).
	// In addition, Mining adds an additional transaction to the block with the receiver address for the reward.
	blockChain.Mining()
	fmtBlockChain()

}
