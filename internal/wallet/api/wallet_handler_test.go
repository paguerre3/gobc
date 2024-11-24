package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/blockchain/configs"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
	"github.com/stretchr/testify/assert"
)

var (
	config = configs.Instance()
)

func TestWalletHandlerCopyright(t *testing.T) {
	// Create a test Echo instance
	e := echo.New()

	getCopyrightUseCase := common_app.NewGetCopyrightUseCase()
	getWalletUseCase := wallet_app.NewGetWalletUseCase(config.Test().ServerPort())
	// Create a test WalletHandler instance
	walletHandler := NewWalletHandler(getCopyrightUseCase, getWalletUseCase)

	// Test case 1: Year method
	req, err := http.NewRequest("GET", "/copyright", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = walletHandler.GetCopyright(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Header().Get("Content-Type"), "application/json")
	var data map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &data)
	assert.NoError(t, err)
	assert.Equal(t, float64(config.Wallet().CopyrightYear()), data["year"])
}

func TestWalletHandlerGetWallet(t *testing.T) {
	getCopyrightUseCase := common_app.NewGetCopyrightUseCase()
	getWalletUseCase := wallet_app.NewGetWalletUseCase(config.Test().ServerPort())
	// Create a test WalletHandler instance
	walletHandler := NewWalletHandler(getCopyrightUseCase, getWalletUseCase)

	// Test case 1: GetWallet (creation of new wallet)
	req, err := http.NewRequest("GET", "/wallet-mock", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err = walletHandler.GetWallet(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "privateKey")
	assert.Contains(t, jsonMap, "publicKey")
	assert.Contains(t, jsonMap, "blockChainAddress")
}
