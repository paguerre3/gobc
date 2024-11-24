package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/blockchain/configs"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
	"github.com/stretchr/testify/assert"
)

var (
	config = configs.Instance()
)

func TestBlockChainHandler(t *testing.T) {
	// Create a test GetBlockChainUseCase instance
	walletServerPort := config.TestServerPort()
	blockChainServerPort := strings.ReplaceAll(walletServerPort, "0", "1")
	getWalletUseCase := wallet_app.NewGetWalletUseCase(walletServerPort)
	wallet, _ := getWalletUseCase.Instance()
	getBlockChainUseCase := application.NewGetBlockChainUseCase(wallet,
		blockChainServerPort, false)

	// Create a test BlockChainHandler instance
	bch := NewBlockChainHandler(getBlockChainUseCase)

	// Test case 1: GetBlockChain with existing block chain
	req, err := http.NewRequest("GET", "/block-chain-mock", nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err = bch.GetBlockChain(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &jsonMap)
	assert.NoError(t, err)

	assert.Contains(t, jsonMap, "transactionPool")
	assert.Contains(t, jsonMap, "chain")
	assert.Contains(t, jsonMap, "blockChainAddressOfRewardRecipient")
	assert.Contains(t, jsonMap, "checkFunds")
}
