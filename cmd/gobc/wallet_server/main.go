package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	wallet_web "github.com/paguerre3/blockchain/internal/wallet/infrastructure/web"
)

func registerWalletHandlers(e *echo.Echo, serverPort string) {
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
	walletApi := wallet_web.NewWalletHandler()
	e.GET("/contact", walletApi.Contact)
	e.GET("/copyright", walletApi.Copyright)
}

func main() {
	// The wallet gateway points to a BlockChain server address:
	portSuffix := 0
	gateway := fmt.Sprintf("http://localhost:500%d", portSuffix)
	serverPort := fmt.Sprintf(":808%d", portSuffix)
	common_web.NewServerNode("Wallet", serverPort, gateway,
		registerWalletHandlers).InitAndRun()
}
