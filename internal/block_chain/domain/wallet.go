package domain

import "crypto/ecdsa"

// Wallet interface to be used by use case in order to avoid cross domain dependencies
// when injecting wallets
type Wallet interface {
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
	BlockChainAddress() string
}
