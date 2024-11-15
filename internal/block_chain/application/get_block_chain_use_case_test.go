package application

import (
	"strings"
	"sync"
	"testing"

	"github.com/paguerre3/blockchain/internal/block_chain/domain"
	wallet_domain "github.com/paguerre3/blockchain/internal/wallet/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetBlockChainUseCase(t *testing.T) {

	// Test case 1: cache hit
	wallet1 := wallet_domain.NewWallet()
	serverPort := domain.TEST_SERVER_PORT
	gbc := NewGetBlockChainUseCase(wallet1, serverPort, true)
	bc1 := gbc.Instance()
	assert.NotNil(t, bc1)
	assert.Equal(t, serverPort, bc1.ServerPort())
	bc1Cached := gbc.Instance()
	assert.NotNil(t, bc1Cached)
	assert.Equal(t, bc1, bc1Cached)

	// Test case 2: cache miss, so create and get new instance
	wallet2 := wallet_domain.NewWallet()
	serverPort2 := strings.ReplaceAll(domain.TEST_SERVER_PORT, "0", "1")
	gbc2 := NewGetBlockChainUseCase(wallet2, serverPort2, true)
	bc2 := gbc2.Instance()
	assert.NotNil(t, bc2)
	assert.Equal(t, serverPort2, bc2.ServerPort())

	// Test case 3: concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gbc3 := NewGetBlockChainUseCase(wallet1, serverPort, true)
			bc3 := gbc3.Instance()
			assert.NotNil(t, bc3)
			assert.Equal(t, serverPort, bc3.ServerPort())
		}()
	}
	wg.Wait()

	// Test case 4: check port
	gbc4 := NewGetBlockChainUseCase(wallet1, serverPort, false)
	bc4 := gbc4.Instance()
	assert.NotNil(t, bc4)
	assert.Equal(t, serverPort, bc4.ServerPort())
}
