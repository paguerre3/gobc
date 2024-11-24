package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
)

type WalletHandler interface {
	Contact(c echo.Context) error
}

type walletHandler struct {
	getCopyrightUseCase common_app.GetCopyrightUseCase
}

func NewWalletHandler(getCopyrightUseCase common_app.GetCopyrightUseCase) WalletHandler {
	return &walletHandler{getCopyrightUseCase: getCopyrightUseCase}
}

func (w *walletHandler) Contact(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", w.getCopyrightUseCase.GetCopyright())
}
