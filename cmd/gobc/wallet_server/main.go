package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	wallet_api "github.com/paguerre3/blockchain/internal/wallet/api"
)

const (
	SERVER_PORT = ":8080"
)

func registerWalletHandlers(e *echo.Echo, serverPort string) {
	e.Renderer = common_api.NewTemplateRenderer(wallet_api.WALLET_TEMPLATES_PATH)

	// handlers
	walletApi := wallet_api.NewWalletHandler()
	e.GET("/", walletApi.Index)
	e.GET("/welcome", walletApi.Index)
}

func main() {
	// The wallet gateway points to a BlockChain server address:
	gateway := fmt.Sprintf("http://localhost:500%d", 0)
	common_api.NewServerNode("Wallet", SERVER_PORT, gateway,
		registerWalletHandlers).InitAndRun()
}
