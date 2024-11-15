package application

import (
	"sync"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
)

var (
	blockChainCache map[string]domain.BlockChain = make(map[string]domain.BlockChain)
	mutex           sync.Mutex
)

type GetBlockChainUseCase interface {
	Instance() domain.BlockChain
}

type getBlockChainUseCase struct {
	wallet     domain.Wallet
	serverPort string
	checkFunds bool
}

func NewGetBlockChainUseCase(wallet domain.Wallet, serverPort string,
	checkFunds bool) GetBlockChainUseCase {
	return &getBlockChainUseCase{
		wallet:     wallet,
		serverPort: serverPort,
		checkFunds: checkFunds,
	}
}

func (gbc *getBlockChainUseCase) Instance() domain.BlockChain {
	bc, ok := blockChainCache[gbc.serverPort]
	if !ok {
		mutex.Lock()
		defer mutex.Unlock()
		bc, ok = blockChainCache[gbc.serverPort]
		if !ok {
			bc = domain.NewBlockchain(gbc.wallet.BlockChainAddress(), gbc.checkFunds, gbc.serverPort)
			blockChainCache[bc.ServerPort()] = bc
		}
	}
	return bc
}
