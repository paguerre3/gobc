package configs

import (
	"os"
	"sync"

	common_web "github.com/paguerre3/blockchain/internal/common/infrastructure/web"
	"gopkg.in/yaml.v3"
)

var (
	instance           Config
	once               sync.Once
	configPathResolver = common_web.NewPathResolver()
)

// Interfaces
type Config interface {
	Test() TestConfig
	BlockChain() BlockChainConfig
	Wallet() WalletConfig
}

type TestConfig interface {
	ServerPort() string
}

type BlockChainConfig interface {
	ServerPort() string
	GenesisSenderAddress() string
	GenesisRecipientAddress() string
	MiningDifficulty() int       // Increasing difficulty means more time for guessing Nonce, e.g., 4=0000 takes around 10 minutes or more.
	MiningSenderAddress() string // Blockchain mining server address that "sends" rewards.
	MiningReward() float64
	MyRewardRecipientAddress() string // Block Chain address to obtain mining reward.
	CheckFunds() bool                 // Enable for prod mode to avoid sending money without sufficient funds.
}

type WalletConfig interface {
	ServerPort() string
	Gateway() string
	FrontendDevServer() string
	FrontendProdServer() string
	CopyrightYear() int
	TemplatesDir() string
}

// Private Struct Implementations

type config struct {
	test       testConfig
	blockChain blockChainConfig
	wallet     walletConfig
}

type testConfig struct {
	Port string `yaml:"serverPort"`
}

type blockChainConfig struct {
	Port                  string  `yaml:"serverPort"`
	GenesisSenderAddr     string  `yaml:"genesisSenderAddress"`
	GenesisRecipientAddr  string  `yaml:"genesisRecipientAddress"`
	MiningDiff            int     `yaml:"miningDifficulty"`
	MiningSenderAddr      string  `yaml:"miningSenderAddress"`
	MiningRewd            float64 `yaml:"miningReward"`
	MyRewardRecipientAddr string  `yaml:"myRewardRecipientAddress"`
	ChkFunds              bool    `yaml:"checkFunds"`
}

type walletConfig struct {
	Port            string `yaml:"serverPort"`
	Gwy             string `yaml:"gateway"`
	FrontendDevSrv  string `yaml:"frontendDevServer"`
	FrontendProdSrv string `yaml:"frontendProdServer"`
	CopyrYear       int    `yaml:"copyrightYear"`
	TempltDir       string `yaml:"templatesDir"`
}

// Singleton Instance Function

func Instance() Config {
	if instance == nil {
		once.Do(func() {
			if instance == nil {
				configPath := configPathResolver("configs/config.yaml")
				data, err := os.ReadFile(configPath)
				if err != nil {
					panic(err)
				}

				var raw struct {
					Test       testConfig       `yaml:"test"`
					BlockChain blockChainConfig `yaml:"blockChain"`
					Wallet     walletConfig     `yaml:"wallet"`
				}

				if err := yaml.Unmarshal(data, &raw); err != nil {
					panic(err)
				}

				instance = &config{
					test:       raw.Test,
					blockChain: raw.BlockChain,
					wallet:     raw.Wallet,
				}
			}
		})
	}
	return instance
}

// Methods to implement interfaces

func (c *config) Test() TestConfig {
	return &c.test
}

func (c *config) BlockChain() BlockChainConfig {
	return &c.blockChain
}

func (c *config) Wallet() WalletConfig {
	return &c.wallet
}

func (t *testConfig) ServerPort() string {
	return t.Port
}

func (b *blockChainConfig) ServerPort() string              { return b.Port }
func (b *blockChainConfig) GenesisSenderAddress() string    { return b.GenesisSenderAddr }
func (b *blockChainConfig) GenesisRecipientAddress() string { return b.GenesisRecipientAddr }
func (b *blockChainConfig) MiningDifficulty() int           { return b.MiningDiff }
func (b *blockChainConfig) MiningSenderAddress() string     { return b.MiningSenderAddr }
func (b *blockChainConfig) MiningReward() float64           { return b.MiningRewd }
func (b *blockChainConfig) MyRewardRecipientAddress() string {
	return b.MyRewardRecipientAddr
}
func (b *blockChainConfig) CheckFunds() bool { return b.ChkFunds }

func (w *walletConfig) ServerPort() string         { return w.Port }
func (w *walletConfig) Gateway() string            { return w.Gwy }
func (w *walletConfig) FrontendDevServer() string  { return w.FrontendDevSrv }
func (w *walletConfig) FrontendProdServer() string { return w.FrontendProdSrv }
func (w *walletConfig) CopyrightYear() int         { return w.CopyrYear }
func (w *walletConfig) TemplatesDir() string       { return w.TempltDir }
