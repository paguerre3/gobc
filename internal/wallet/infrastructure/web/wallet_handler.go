package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
)

const (
	WALLET_TEMPLATES_PATH = "internal/wallet/infrastructure/templates"
	WALLET_COPYRIGHT_YEAR = 2022
)

type WalletHandler interface {
	Contact(c echo.Context) error
	Copyright(c echo.Context) error
}

type walletHandler struct {
}

func NewWalletHandler() WalletHandler {
	return &walletHandler{}
}

func (w *walletHandler) Contact(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", hydrateYear())
}

func (w *walletHandler) Copyright(c echo.Context) error {
	return c.JSON(http.StatusOK, hydrateYear())
}

func hydrateYear() common_web.PageData {
	data := common_web.PageData{
		Year: WALLET_COPYRIGHT_YEAR,
	}
	return data
}
