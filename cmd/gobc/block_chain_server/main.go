package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	block_chain_api "github.com/paguerre3/blockchain/internal/block_chain/api"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
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
	portSuffix := 0
	serverPort := fmt.Sprintf(":500%d", portSuffix)
	// no gateway for blockChain servers:
	common_web.NewServerNode("BlockChain", serverPort, "",
		registerBlockChainHandlers).InitAndRun()
}
