package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/paguerre3/blockchain/internal/domain"
)

func newBlockChainWithFmt() func() domain.BlockChain {
	bc := domain.NewBlockchain()
	return func() domain.BlockChain {
		fmt.Println(strings.Repeat("#", 75))
		json, _ := json.MarshalIndent(bc, "", "  ")
		fmt.Println(string(json))
		return bc
	}
}

func main() {
	blockChainFmt := newBlockChainWithFmt()
	blockChain := blockChainFmt()

	blockChain.CreateAppendTransaction("sender_address_A", "receiver_address_A", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateAppendBlock(1, previousHash)
	blockChainFmt()

	blockChain.CreateAppendTransaction("sender_address_B_1", "receiver_address_B_1", 2.5)
	blockChain.CreateAppendTransaction("sender_address_B_2", "receiver_address_B_2", 5.0)
	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateAppendBlock(2, previousHash)
	blockChainFmt()
}
