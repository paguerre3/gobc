package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
)

type Wallet interface {
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
}

type wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewWallet() Wallet {
	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return &wallet{
		privateKey: pk,
		// publicKey is a part of the private key
		publicKey: &pk.PublicKey,
	}
}

func (w *wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *wallet) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
	}{
		PrivateKey: fmt.Sprintf("%x", w.privateKey.D.Bytes()), // hex.EncodeToString(hash[:]) OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f
		PublicKey:  fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes()),
	})
}
