package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
)

type BlockChainHandler interface {
	GetBlockChain(c echo.Context) error
}

type blockChainHandler struct {
	getBlockChainUseCase application.GetBlockChainUseCase
}

func NewBlockChainHandler(getBlockChainUseCase application.GetBlockChainUseCase) BlockChainHandler {
	return &blockChainHandler{
		getBlockChainUseCase: getBlockChainUseCase,
	}
}

func (bch *blockChainHandler) GetBlockChain(c echo.Context) error {
	bc, ok := bch.getBlockChainUseCase.Instance()
	status := http.StatusOK
	if !ok {
		status = http.StatusCreated
	}
	// The c.JSON(status, bc) function call in Echo automatically marshals the bc object into a proper JSON response.
	return c.JSON(status, bc)
}
