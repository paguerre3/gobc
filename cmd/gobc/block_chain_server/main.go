package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"

	block_chain_api "github.com/paguerre3/blockchain/internal/block_chain/api"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
)

const (
	MAX_BC_SERVERS = 3
)

func registerBlockChainHandlers(e *echo.Echo, serverPort string) {
	// Use cases
	getWalletUseCase := wallet_app.NewGetWalletUseCase(serverPort)
	getBlockChainUseCase := application.NewGetBlockChainUseCase(getWalletUseCase.Instance(),
		serverPort, false)

	// Handlers
	blockChainApi := block_chain_api.NewBlockChainHandler(getBlockChainUseCase)
	e.GET("/block-chain", blockChainApi.GetBlockChain)

	e.GET("/ping", common_api.Ping)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < MAX_BC_SERVERS; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			serverPort := fmt.Sprintf(":500%d", port)
			// no gateway for blockChain servers:
			common_api.NewServerNode("BlockChain", serverPort, "",
				registerBlockChainHandlers).InitAndRun()
		}(i)
	}
	wg.Wait()
}
