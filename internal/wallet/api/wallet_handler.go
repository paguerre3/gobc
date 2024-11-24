package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
)

type WalletHandler interface {
	GetCopyright(c echo.Context) error
	GetWallet(c echo.Context) error
}

type walletHandler struct {
	getCopyrightUseCase common_app.GetCopyrightUseCase
	getWalletUseCase    wallet_app.GetWalletUseCase
}

func NewWalletHandler(getCopyrightUseCase common_app.GetCopyrightUseCase,
	getWalletUseCase wallet_app.GetWalletUseCase) WalletHandler {
	return &walletHandler{
		getCopyrightUseCase: getCopyrightUseCase,
		getWalletUseCase:    getWalletUseCase,
	}
}

func (w *walletHandler) GetCopyright(c echo.Context) error {
	return c.JSON(http.StatusOK, w.getCopyrightUseCase.GetCopyright())
}

func (w *walletHandler) GetWallet(c echo.Context) error {
	wallet, ok := w.getWalletUseCase.Instance()
	status := http.StatusOK
	if !ok {
		status = http.StatusCreated
	}
	// The c.JSON(status, w) function call in Echo automatically marshals the w object into a proper JSON response.
	return c.JSON(status, wallet)
}
