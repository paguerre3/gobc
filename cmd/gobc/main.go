package main

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	block_chain_api "github.com/paguerre3/blockchain/internal/block_chain/api"
	"github.com/paguerre3/blockchain/internal/block_chain/application"
	common_api "github.com/paguerre3/blockchain/internal/common/api"
	wallet_app "github.com/paguerre3/blockchain/internal/wallet/application"
)

const (
	MAX_BC_SERVERS = 3
)

type ServerNode interface {
	InitAndRun()
}

type serverNode struct {
	port string
}

func newServerNode(port string) ServerNode {
	return &serverNode{port: port}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < MAX_BC_SERVERS; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			newServerNode(fmt.Sprintf(":500%d", port)).InitAndRun()
		}(i)
	}
	wg.Wait()
}

func (s *serverNode) InitAndRun() {
	log.Infof("Starting Blockchain Server on TCP Port %s", s.port)
	// Initialize a new Echo instance
	e := echo.New()

	// Define a simple GET route
	e.GET("/ping", common_api.Ping)

	// use cases
	getWalletUseCase := wallet_app.NewGetWalletUseCase(s.port)
	getBlockChainUseCase := application.NewGetBlockChainUseCase(getWalletUseCase.Instance(), s.port, false)

	// handlers
	blockChainApi := block_chain_api.NewBlockChainHandler(getBlockChainUseCase)
	e.GET("/block-chain", blockChainApi.GetBlockChain)

	// Start the server on the specified port
	e.Start(s.port)
}
