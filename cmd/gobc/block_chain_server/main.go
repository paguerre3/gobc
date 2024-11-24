package main

import (
	"github.com/labstack/echo/v4"

	"github.com/paguerre3/blockchain/configs"
	block_chain_api "github.com/paguerre3/blockchain/internal/block_chain/api"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
)

var (
	config = configs.Instance()
)

func registerBlockChainHandlers(e *echo.Echo, serverPort string, gateway string) {
	// Use cases
	getWalletUseCase := wallet_app.NewGetWalletUseCase(config.Wallet().ServerPort())
	wallet, _ := getWalletUseCase.Instance()
	getBlockChainUseCase := application.NewGetBlockChainUseCase(wallet,
		serverPort, config.BlockChain().CheckFunds())

	// Handlers
	blockChainApi := block_chain_api.NewBlockChainHandler(getBlockChainUseCase)
	e.GET("/block-chain", blockChainApi.GetBlockChain)
	e.GET("/ping", common_api.Ping)
}

func main() {
	common_web.NewServerNode("BlockChain", config.BlockChain().ServerPort(),
		// no gateway for blockChain servers:
		"", registerBlockChainHandlers).InitAndRun()
}
