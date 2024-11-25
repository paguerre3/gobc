package domain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

type Signature interface {
	R() *big.Int
	S() *big.Int
}

type signature struct {
	r, s *big.Int
}

func NewSignature(r, s *big.Int) Signature {
	return &signature{r: r, s: s}
}

func (s *signature) R() *big.Int {
	return s.r
}

func (s *signature) S() *big.Int {
	return s.s
}

func (s *signature) String() string {
	return fmt.Sprintf("%64x%64x", s.r, s.s) // hex.EncodeToString(hash[:]) OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f
}

func (s *signature) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		R string `json:"r"`
		S string `json:"s"`
	}{
		R: s.r.String(),
		S: s.s.String(),
	})
}

func stringToBigIntTuple(s string) (big.Int, big.Int) {
	bx, _ := hex.DecodeString(s[:64]) // 0 - 63 (64)
	by, _ := hex.DecodeString(s[64:]) // 64 - 127 (64) => (64 + 64) = 128

	var bix, biy big.Int
	bix.SetBytes(bx)
	biy.SetBytes(by)

	return bix, biy
}

func PublicKeyFromString(s string) *ecdsa.PublicKey {
	if len(s) != 128 {
		return nil
	}
	x, y := stringToBigIntTuple(s)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: &x, Y: &y}
}

func PrivateKeyFromString(s string, publicKey *ecdsa.PublicKey) *ecdsa.PrivateKey {
	if len(s) != 64 || publicKey == nil {
		return nil
	}
	b, _ := hex.DecodeString(s[:])

	var bi big.Int
	bi.SetBytes(b)

	return &ecdsa.PrivateKey{D: &bi, PublicKey: *publicKey}
}

func SignatureFromString(s string) Signature {
	if len(s) != 128 {
		return nil
	}
	x, y := stringToBigIntTuple(s)
	// signature uses public key
	return &signature{r: &x, s: &y}
}
