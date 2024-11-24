package application

import (
	"strings"
	"sync"
	"testing"

	"github.com/paguerre3/blockchain/configs"
	wallet_domain "github.com/paguerre3/blockchain/internal/wallet/domain"
	"github.com/stretchr/testify/assert"
)

var (
	config = configs.Instance()
)

func TestGetBlockChainUseCase(t *testing.T) {

	// Test case 1: cache hit
	wallet1 := wallet_domain.NewWallet()
	serverPort := config.Test().ServerPort()
	gbc := NewGetBlockChainUseCase(wallet1, serverPort, true)
	bc1, ok1 := gbc.Instance()
	assert.NotNil(t, bc1)
	assert.Equal(t, serverPort, bc1.ServerPort())
	assert.False(t, ok1)
	bc1Cached, ok2 := gbc.Instance()
	assert.NotNil(t, bc1Cached)
	assert.Equal(t, bc1, bc1Cached)
	assert.Equal(t, serverPort, bc1Cached.ServerPort())
	assert.Equal(t, wallet1.BlockChainAddress(), bc1Cached.BlockChainAddressOfRewardRecipient())
	assert.Equal(t, bc1.BlockChainAddressOfRewardRecipient(), bc1Cached.BlockChainAddressOfRewardRecipient())
	assert.True(t, ok2)

	// Test case 2: cache miss, so create and get new instance
	wallet2 := wallet_domain.NewWallet()
	serverPort2 := strings.ReplaceAll(config.Test().ServerPort(), "0", "1")
	gbc2 := NewGetBlockChainUseCase(wallet2, serverPort2, true)
	bc2, _ := gbc2.Instance()
	assert.NotNil(t, bc2)
	assert.Equal(t, serverPort2, bc2.ServerPort())

	// Test case 3: concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gbc3 := NewGetBlockChainUseCase(wallet1, serverPort, true)
			bc3, _ := gbc3.Instance()
			assert.NotNil(t, bc3)
			assert.Equal(t, serverPort, bc3.ServerPort())
		}()
	}
	wg.Wait()

	// Test case 4: check port
	gbc4 := NewGetBlockChainUseCase(wallet1, serverPort, false)
	bc4, _ := gbc4.Instance()
	assert.NotNil(t, bc4)
	assert.Equal(t, serverPort, bc4.ServerPort())
}
