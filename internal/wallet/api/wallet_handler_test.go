package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	"github.com/stretchr/testify/assert"
)

func TestWalletHandler(t *testing.T) {
	// Create a test Echo instance
	e := echo.New()
	e.Renderer = common_api.NewTemplateRenderer(WALLET_TEMPLATES_PATH)

	// Create a test WalletHandler instance
	walletHandler := NewWalletHandler()

	// Test case 1: Index method
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = walletHandler.Index(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Header().Get("Content-Type"), "text/html")
	strBody := rec.Body.String()
	assert.Contains(t, strBody, "<!DOCTYPE html>")
	assert.Contains(t, strBody, "<title>Cami Wallet</title>")
}
