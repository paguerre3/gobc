package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
)

const (
	WALLET_TEMPLATES_PATH = "internal/wallet/api/templates"
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
	data := common_api.PageData{
		Year: time.Now().Year(),
	}
	return c.Render(http.StatusOK, "index.html", data)
}
