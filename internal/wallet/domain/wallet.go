package domain

import "crypto/ecdsa"

type Wallet interface {
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
}

type wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() Wallet {
	return &wallet{}
}

func (w *wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}
