package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_web "github.com/paguerre3/blockchain/internal/wallet/infrastructure/web"
)

const (
	MAX_WALLETS = 3
)

func registerWalletHandlers(e *echo.Echo, serverPort string) {
	e.Renderer = common_web.NewTemplateRenderer(wallet_web.WALLET_TEMPLATES_PATH)

	// handlers
	walletApi := wallet_web.NewWalletHandler()
	e.GET("/", walletApi.Index)
	e.GET("/welcome", walletApi.Index)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < MAX_WALLETS; i++ {
		wg.Add(1)
		go func(portSufix int) {
			defer wg.Done()
			// The wallet gateway points to a BlockChain server address:
			gateway := fmt.Sprintf("http://localhost:500%d", portSufix)
			serverPort := fmt.Sprintf(":808%d", portSufix)
			common_web.NewServerNode("Wallet", serverPort, gateway,
				registerWalletHandlers).InitAndRun()
		}(i)
	}
	wg.Wait()
}
