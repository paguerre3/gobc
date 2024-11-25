package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet interface {
	PrivateKey() *ecdsa.PrivateKey
	PublicKey() *ecdsa.PublicKey
	BlockChainAddress() string
}

type wallet struct {
	privateKey        *ecdsa.PrivateKey
	publicKey         *ecdsa.PublicKey
	blockChainAddress string
}

func generateKeys() (*ecdsa.PrivateKey, error) {
	// 1. Create ECDSA private key (32 bytes), public key (64 bytes)
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func createBlockChainAddress(privateKey *ecdsa.PrivateKey) string {
	// 2. Perform SHA-256 hashing on the public key (32 bytes)
	h2 := sha256.New()
	h2.Write(privateKey.PublicKey.X.Bytes())
	h2.Write(privateKey.PublicKey.Y.Bytes())
	digest2 := h2.Sum(nil)
	// 3. Perform RIPEMD-160 hashing on the resulting SHA-256 hash (20 bytes)
	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)
	// 4. Add version byte in front of RIPEMD-160 hash (1 byte: "0x00=Main Network" or "Test Network")
	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])
	// 5. Perform SHA-256 hashing on the extended RIPEMD-160 hash result (20 bytes)
	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)
	//6. Perform SHA-256 hashing on the resulting SHA-256 hash (32 bytes: 2nd SHA)
	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)
	//7. Take the first "4 bytes" of the second SHA-256 hash for checksum (4 bytes)
	chsum := digest6[:4]
	//8. Add the 4 checksum bytes from 7 at the end of the extended RIPEMD-160 hash from 4 (25 bytes)
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])
	//9. Convert the result from a byte string into base58
	return base58.Encode(dc8)
}

func NewWallet() Wallet {
	pk, err := generateKeys()
	if err != nil {
		panic(err)
	}
	address := createBlockChainAddress(pk)
	return &wallet{
		privateKey: pk,
		// publicKey is a part of the private key
		publicKey:         &pk.PublicKey,
		blockChainAddress: address,
	}
}

func (w *wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

func (w *wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

func (w *wallet) BlockChainAddress() string {
	return w.blockChainAddress
}

func (w *wallet) MarshalJSON() ([]byte, error) {
	// its required to marshal lower cappital fields for json:
	return json.Marshal(struct {
		PrivateKey        string `json:"privateKey"`
		PublicKey         string `json:"publicKey"`
		BlockChainAddress string `json:"blockChainAddress"`
	}{
		PrivateKey:        fmt.Sprintf("%x", w.privateKey.D.Bytes()), // hex.EncodeToString(hash[:]) OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f
		PublicKey:         fmt.Sprintf("%064x%064x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes()),
		BlockChainAddress: w.blockChainAddress,
	})
}
