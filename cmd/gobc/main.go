package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
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

func main() {
	fmtBlockChain := newBlockChainWithFmt()
	blockChain := fmtBlockChain()

	blockChain.CreateAppendTransaction("sender_address_1", "receiver_address_1", 1.0)
	//previousHash := blockChain.LastBlock().Hash()
	//nonce := blockChain.ProofOfWork()
	//blockChain.CreateAppendBlock(nonce, previousHash)
	// Mining already covers calculating the nonce/PoW and creating the Block with previous Hash (previous 3 sentences).
	// In addition, Mining adds an additional transaction to the block with the receiver address for the reward.
	blockChain.Mining()
	fmtBlockChain()

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
}
