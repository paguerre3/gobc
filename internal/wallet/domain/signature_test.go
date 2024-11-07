package domain

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignatureMarshalJSON(t *testing.T) {
	s := newSignature(big.NewInt(1), big.NewInt(2))
	jsonBytes, err := json.Marshal(s)
	assert.NoError(t, err)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)
	assert.Contains(t, jsonMap, "r")
	assert.Contains(t, jsonMap, "s")
}

func TestSignatureUnmarshalJSON(t *testing.T) {
	s := newSignature(big.NewInt(1), big.NewInt(2))
	jsonBytes, err := json.Marshal(s)
	assert.NoError(t, err)
	var jsonMap map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonMap)
	assert.NoError(t, err)
	assert.Contains(t, jsonMap, "r")
	assert.Contains(t, jsonMap, "s")
}

func TestSignatureString(t *testing.T) {
	s := signature{r: big.NewInt(1), s: big.NewInt(2)}
	assert.Equal(t, "12", s.String())
}
