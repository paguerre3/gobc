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
	fmtBlockChain := newBlockChainWithFmt()
	blockChain := fmtBlockChain()

	blockChain.CreateAppendTransaction("sender_address_1_flow_A", "receiver_address_1_flow_A", 1.0)
	previousHash := blockChain.LastBlock().Hash()
	nonce := blockChain.ProofOfWork()
	blockChain.CreateAppendBlock(nonce, previousHash)
	fmtBlockChain()

	blockChain.CreateAppendTransaction("sender_address_2_flow_B", "receiver_address_2_flow_B", 2.5)
	blockChain.CreateAppendTransaction("sender_address_3_flow_B", "receiver_address_3_flow_B", 5.0)
	previousHash = blockChain.LastBlock().Hash()
	nonce = blockChain.ProofOfWork()
	blockChain.CreateAppendBlock(nonce, previousHash)
	fmtBlockChain()
}
