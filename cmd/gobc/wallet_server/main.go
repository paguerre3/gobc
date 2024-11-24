package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_api "github.com/paguerre3/blockchain/internal/wallet/api"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
	wallet_web "github.com/paguerre3/blockchain/internal/wallet/infrastructure/web"
)

func registerWalletHandlers(e *echo.Echo, serverPort string, gateway string) {
	e.Renderer = common_web.NewTemplateRenderer(wallet_web.WALLET_TEMPLATES_PATH)

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:5173", // React dev server URL
			"http://localhost:4173", // Production
		},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// handlers
	getCopyrightUseCase := common_app.NewGetCopyrightUseCase()
	walletPage := wallet_web.NewWalletHandler(getCopyrightUseCase)
	e.GET("/contact", walletPage.Contact)

	getWalletUseCase := wallet_app.NewGetWalletUseCase(serverPort)
	walletApi := wallet_api.NewWalletHandler(getCopyrightUseCase, getWalletUseCase)
	e.GET("/copyright", walletApi.GetCopyright)
	e.GET("/wallet", walletApi.GetWallet)
	e.GET("/ping", common_api.Ping)
}

func main() {
	// The wallet gateway points to a BlockChain server address:
	portSuffix := 0
	gateway := fmt.Sprintf("http://localhost:500%d", portSuffix)
	serverPort := fmt.Sprintf(":808%d", portSuffix)
	common_web.NewServerNode("Wallet", serverPort, gateway,
		registerWalletHandlers).InitAndRun()
}
