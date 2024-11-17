package api

import (
	"path"

	"github.com/labstack/echo/v4"
)

const (
	TEMPLATE_DIR = "./templates"
)

type WalletHandler interface {
	Index(c echo.Context) error
}

type walletHandler struct {
}

func NewWalletHandler() WalletHandler {
	return &walletHandler{}
}

func (w *walletHandler) Index(c echo.Context) error {
	return c.File(path.Join(TEMPLATE_DIR, "index.html"))
}
