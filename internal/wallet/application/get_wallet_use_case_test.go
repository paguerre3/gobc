package application

import (
	"strings"
	"sync"
	"testing"

	common_domain "github.com/paguerre3/blockchain/internal/common/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetWalletUseCase(t *testing.T) {
	// Create a test server port
	serverPort := common_domain.TEST_SERVER_PORT

	// Create a test GetWalletUseCase instance
	gwc := NewGetWalletUseCase(serverPort)

	// Test case 1: cache hit
	wallet1 := gwc.Instance()
	assert.NotNil(t, wallet1)
	wallet1Cahched := gwc.Instance()
	assert.NotNil(t, wallet1Cahched)
	assert.Equal(t, wallet1.BlockChainAddress(), wallet1Cahched.BlockChainAddress())
	assert.Equal(t, wallet1.PrivateKey(), wallet1Cahched.PrivateKey())
	assert.Equal(t, wallet1.PublicKey(), wallet1Cahched.PublicKey())
	assert.Equal(t, wallet1, wallet1Cahched)

	// Test case 2: cache miss
	serverPort2 := strings.ReplaceAll(serverPort, "0", "1")
	gwc2 := NewGetWalletUseCase(serverPort2)
	wallet2 := gwc2.Instance()
	assert.NotNil(t, wallet2)
	assert.NotEqual(t, wallet1, wallet2)
	assert.NotEqual(t, wallet1.BlockChainAddress(), wallet2.BlockChainAddress())

	// Test case 3: concurrent access
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gwc3 := NewGetWalletUseCase(serverPort)
			wallet3 := gwc3.Instance()
			assert.NotNil(t, wallet3)
			assert.Equal(t, wallet1, wallet3)
		}()
	}
	wg.Wait()
}
