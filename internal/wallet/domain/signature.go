package domain

import (
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

func newSignature(r, s *big.Int) Signature {
	return &signature{r: r, s: s}
}

func (s *signature) R() *big.Int {
	return s.r
}

func (s *signature) S() *big.Int {
	return s.s
}

func (s *signature) String() string {
	return fmt.Sprintf("%x%x", s.r, s.s) // hex.EncodeToString(hash[:]) OR: fmt.Sprintf("%x", hash) // %x	base 16, with lower-case letters for a-f
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
