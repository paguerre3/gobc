package configs

import (
	"os"
	"sync"

	coommon_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"

	"gopkg.in/yaml.v3"
)

var (
	instance           Config
	once               sync.Once
	configPathResolver = coommon_web.NewPathResolver()
)

type Config interface {
	TestServerPort() string

	GenesisSenderAddress() string
	GenesisRecipientAddress() string
	MiningDifficulty() int      // increasing difficulty means more time for guessing Nonce, e.g. 4=0000 is arround 10 minutes or more
	MinigSenderAddress() string // block chain mining server address that "sends" rewards
	MiningReward() float64
	MyBlockChainRecipientAddres() string // address for receiving mining rewards
	CheckFunds() bool                    // enable for prod mode, i.e. to avoid sending money without founds

	WalletFrontendDevServer() string
	WalletFrontendProdServer() string
	WalletCopyrightYear() int
	WalletTemplatesDir() string
}

type config struct {
	Test struct {
		ServerPort string `yaml:"serverPort"`
	} `yaml:"test"`
	BlockChain struct {
		GenesisSenderAddress         string  `yaml:"genesisSenderAddress"`
		GenesisRecipientAddress      string  `yaml:"genesisRecipientAddress"`
		MiningDifficulty             int     `yaml:"miningDifficulty"`
		MiningSenderAddress          string  `yaml:"miningSenderAddress"`
		MiningReward                 float64 `yaml:"miningReward"`
		MyBlockChainRecipientAddress string  `yaml:"myBlockChainRecipientAddress"`
		CheckFunds                   bool    `yaml:"checkFunds"`
	} `yaml:"blockChain"`
	Wallet struct {
		FrontendDevServer  string `yaml:"frontendDevServer"`
		FrontendProdServer string `yaml:"frontendProdServer"`
		CopyrightYear      int    `yaml:"copyrightYear"`
		TemplatesDir       string `yaml:"templatesDir"`
	} `yaml:"wallet"`
}

func Instance() Config {
	if instance == nil {
		once.Do(func() {
			if instance == nil {
				// TODO: load from config file accordimng to proper environment
				configPath := configPathResolver("configs/config.yaml")
				data, err := os.ReadFile(configPath)
				if err != nil {
					panic(err)
				}

				var conf config
				if err := yaml.Unmarshal(data, &conf); err != nil {
					panic(err)
				}
				instance = &conf
			}
		})
	}
	return instance
}

func (c *config) TestServerPort() string {
	return c.Test.ServerPort
}

func (c *config) GenesisSenderAddress() string {
	return c.BlockChain.GenesisSenderAddress
}

func (c *config) GenesisRecipientAddress() string {
	return c.BlockChain.GenesisRecipientAddress
}

func (c *config) MiningDifficulty() int {
	return c.BlockChain.MiningDifficulty
}

func (c *config) MinigSenderAddress() string {
	return c.BlockChain.MiningSenderAddress
}

func (c *config) MiningReward() float64 {
	return c.BlockChain.MiningReward
}

func (c *config) MyBlockChainRecipientAddres() string {
	return c.BlockChain.MyBlockChainRecipientAddress
}

func (c *config) CheckFunds() bool {
	return c.BlockChain.CheckFunds
}

func (c *config) WalletFrontendDevServer() string {
	return c.Wallet.FrontendDevServer
}

func (c *config) WalletFrontendProdServer() string {
	return c.Wallet.FrontendProdServer
}

func (c *config) WalletCopyrightYear() int {
	return c.Wallet.CopyrightYear
}

func (c *config) WalletTemplatesDir() string {
	return c.Wallet.TemplatesDir
}
