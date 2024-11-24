package application

import (
	"sync"

	"github.com/paguerre3/blockchain/internal/wallet/domain"
)

var (
	walletCache map[string]domain.Wallet = make(map[string]domain.Wallet)
	mutex       sync.Mutex
)

type GetWalletUseCase interface {
	Instance() (domain.Wallet, bool)
}

type getWalletUseCase struct {
	serverPort string
}

func NewGetWalletUseCase(serverPort string) GetWalletUseCase {
	return &getWalletUseCase{
		serverPort: serverPort,
	}
}

func (gwc *getWalletUseCase) Instance() (domain.Wallet, bool) {
	wallet, ok := walletCache[gwc.serverPort]
	if !ok {
		mutex.Lock()
		defer mutex.Unlock()
		wallet, ok = walletCache[gwc.serverPort]
		if !ok {
			wallet = domain.NewWallet()
			walletCache[gwc.serverPort] = wallet
		}
	}
	return wallet, ok
}
