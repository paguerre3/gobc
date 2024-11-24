package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigs(t *testing.T) {
	config := Instance()
	assert.NotNil(t, config)
	assert.Equal(t, ":0000", config.TestServerPort())
	assert.Equal(t, "genesis_sender_address", config.GenesisSenderAddress())
	assert.Equal(t, "genesis_recipient_address", config.GenesisRecipientAddress())
	assert.Equal(t, 3, config.MiningDifficulty())
	assert.Equal(t, "THE_BLOCKCHAIN_MINING_SENDER_ADDRESS", config.MinigSenderAddress())
	assert.Equal(t, 1.0, config.MiningReward())
	assert.Equal(t, "MY_BLOCKCHAIN_RECIPIENT_ADDRESS_TO_OBTAIN_MINING_REWARD", config.MyBlockChainRecipientAddres())
	assert.Equal(t, false, config.CheckFunds())
	assert.Equal(t, "http://localhost:5173", config.WalletFrontendDevServer())
	assert.Equal(t, "http://localhost:4173", config.WalletFrontendProdServer())
	assert.Equal(t, 2022, config.WalletCopyrightYear())
	assert.Equal(t, "internal/wallet/infrastructure/templates", config.WalletTemplatesDir())
	assert.Equal(t, "/cmd/", config.CmdDir())
	assert.Equal(t, "/internal/", config.InternalDir())
}
