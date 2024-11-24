package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/paguerre3/blockchain/configs"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_api "github.com/paguerre3/blockchain/internal/wallet/api"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
	wallet_web "github.com/paguerre3/blockchain/internal/wallet/infrastructure/web"
)

var (
	config = configs.Instance()
)

func registerWalletHandlers(e *echo.Echo, serverPort string, gateway string) {
	e.Renderer = common_web.NewTemplateRenderer(config.WalletTemplatesDir())

	// Enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			config.WalletFrontendDevServer(),  // React dev server URL
			config.WalletFrontendProdServer(), // Production
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
	common_web.NewServerNode("Wallet", config.WalletServerPort(),
		// The wallet gateway points to a BlockChain server address:
		config.WalletGateway(), registerWalletHandlers).InitAndRun()
}
