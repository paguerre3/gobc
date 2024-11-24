package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigs(t *testing.T) {
	config := Instance()
	assert.NotNil(t, config)
	assert.Equal(t, ":0000", config.Test().ServerPort())
	assert.Equal(t, ":5000", config.BlockChain().ServerPort())
	assert.Equal(t, "genesis_sender_address", config.BlockChain().GenesisSenderAddress())
	assert.Equal(t, "genesis_recipient_address", config.BlockChain().GenesisRecipientAddress())
	assert.Equal(t, 3, config.BlockChain().MiningDifficulty())
	assert.Equal(t, "THE_BLOCKCHAIN_MINING_SENDER_ADDRESS", config.BlockChain().MiningSenderAddress())
	assert.Equal(t, 1.0, config.BlockChain().MiningReward())
	assert.Equal(t, "MY_BLOCKCHAIN_RECIPIENT_ADDRESS_TO_OBTAIN_MINING_REWARD", config.BlockChain().MyRewardRecipientAddress())
	assert.Equal(t, false, config.BlockChain().CheckFunds())
	assert.Equal(t, ":8080", config.Wallet().ServerPort())
	assert.Equal(t, "http://localhost:5000", config.Wallet().Gateway())
	assert.Equal(t, "http://localhost:5173", config.Wallet().FrontendDevServer())
	assert.Equal(t, "http://localhost:4173", config.Wallet().FrontendProdServer())
	assert.Equal(t, 2022, config.Wallet().CopyrightYear())
	assert.Equal(t, "internal/wallet/infrastructure/templates", config.Wallet().TemplatesDir())
}
