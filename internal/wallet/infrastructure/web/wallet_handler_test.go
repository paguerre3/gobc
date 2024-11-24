package web

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	common_app "github.com/paguerre3/blockchain/internal/common/application"
	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	"github.com/stretchr/testify/assert"
)

func TestWalletHandlerContact(t *testing.T) {
	// Create a test Echo instance
	e := echo.New()
	e.Renderer = common_web.NewTemplateRenderer(WALLET_TEMPLATES_PATH)

	getCopyrightUseCase := common_app.NewGetCopyrightUseCase()
	// Create a test WalletHandler instance
	walletHandler := NewWalletHandler(getCopyrightUseCase)

	// Test case 1: Index method
	req, err := http.NewRequest("GET", "/contact", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err = walletHandler.Contact(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Header().Get("Content-Type"), "text/html")
	strBody := rec.Body.String()
	assert.Contains(t, strBody, "<!DOCTYPE html>")
	assert.Contains(t, strBody, "<title>Contact</title>")
}
