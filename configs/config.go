package configs

import (
	"sync"
)

const (
	testServerPort              = ":0000"
	genesisSenderAddress        = "genesis_sender_address"
	genesisRecipientAddress     = "genesis_recipient_address"
	miningDifficulty            = 3                                      // increasing difficulty means more time for guessing Nonce, e.g. 4=0000 is arround 10 minutes or more
	minigSenderAddress          = "THE_BLOCKCHAIN_MINING_SENDER_ADDRESS" // block chain mining server address that "sends" rewards
	miningReward                = 1.0
	myBlockChainRecipientAddres = "MY_BLOCKCHAIN_RECIPIENT_ADDRESS_TO_OBTAIN_MINING_REWARD" // address for receiving mining rewards
	checkFunds                  = false                                                     // enable for prod mode, i.e. to avoid sending money without founds
	walletFrontendDevServer     = "http://localhost:5173"
	walletFrontendProdServer    = "http://localhost:4173"
	walletCopyrightYear         = 2022
	walletTemplatesDir          = "internal/wallet/infrastructure/templates"
	cmdDir                      = "/cmd/"
	internalDir                 = "/internal/"
)

var (
	instance Config
	once     sync.Once
)

type Config interface {
	TestServerPort() string

	GenesisSenderAddress() string
	GenesisRecipientAddress() string
	MiningDifficulty() int
	MinigSenderAddress() string
	MiningReward() float64
	MyBlockChainRecipientAddres() string
	CheckFunds() bool

	WalletFrontendDevServer() string
	WalletFrontendProdServer() string
	WalletCopyrightYear() int
	WalletTemplatesDir() string

	CmdDir() string
	InternalDir() string
}

type config struct {
	testServerPort              string
	genesisSenderAddress        string
	genesisRecipientAddress     string
	miningDifficulty            int
	minigSenderAddress          string
	miningReward                float64
	myBlockChainRecipientAddres string
	checkFunds                  bool
	walletFrontendDevServer     string
	walletFrontendProdServer    string
	walletCopyrightYear         int
	walletTemplatesDir          string
	cmdDir                      string
	internalDir                 string
}

func Instance() Config {
	if instance == nil {
		once.Do(func() {
			if instance == nil {
				// TODO: load from config file accordimng to proper environment
				instance = &config{
					testServerPort:              testServerPort,
					genesisSenderAddress:        genesisSenderAddress,
					genesisRecipientAddress:     genesisRecipientAddress,
					miningDifficulty:            miningDifficulty,
					minigSenderAddress:          minigSenderAddress,
					miningReward:                miningReward,
					myBlockChainRecipientAddres: myBlockChainRecipientAddres,
					checkFunds:                  checkFunds,
					walletFrontendDevServer:     walletFrontendDevServer,
					walletFrontendProdServer:    walletFrontendProdServer,
					walletCopyrightYear:         walletCopyrightYear,
					walletTemplatesDir:          walletTemplatesDir,
					cmdDir:                      cmdDir,
					internalDir:                 internalDir,
				}
			}
		})
	}
	return instance
}

func (c *config) TestServerPort() string {
	return c.testServerPort
}

func (c *config) GenesisSenderAddress() string {
	return c.genesisSenderAddress
}

func (c *config) GenesisRecipientAddress() string {
	return c.genesisRecipientAddress
}

func (c *config) MiningDifficulty() int {
	return c.miningDifficulty
}

func (c *config) MinigSenderAddress() string {
	return c.minigSenderAddress
}

func (c *config) MiningReward() float64 {
	return c.miningReward
}

func (c *config) MyBlockChainRecipientAddres() string {
	return c.myBlockChainRecipientAddres
}

func (c *config) CheckFunds() bool {
	return c.checkFunds
}

func (c *config) WalletFrontendDevServer() string {
	return c.walletFrontendDevServer
}

func (c *config) WalletFrontendProdServer() string {
	return c.walletFrontendProdServer
}

func (c *config) WalletCopyrightYear() int {
	return c.walletCopyrightYear
}

func (c *config) WalletTemplatesDir() string {
	return c.walletTemplatesDir
}

func (c *config) CmdDir() string {
	return c.cmdDir
}

func (c *config) InternalDir() string {
	return c.internalDir
}
