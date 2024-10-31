package runes

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WhenValidInputRangeDuringAddToRune_ThenSuccessAggregation(t *testing.T) {
	ro := NewRunOperator()
	c := 'A'
	fmt.Println("uppercase", strings.Repeat("*", 10))
	ro.FmtRune(c) // 65 is the ASCII code of A
	ro.FmtRune('Z')
	fmt.Println("lowercase", strings.Repeat("*", 10))
	ro.FmtRune('a')
	ro.FmtRune('z')
	res, err := ro.AddToRune('A', 3)
	assert.Nil(t, err)
	assert.Equal(t, 'D', res)

	res, err = ro.AddToRune('a', 3)
	assert.Nil(t, err)
	assert.Equal(t, 'd', res)
}

func Test_WhenInalidInputRangeDuringAddToRune_ThenErrorAggregation(t *testing.T) {
	ro := NewRunOperator()
	res, err := ro.AddToRune('A', 1000)
	ro.FmtRune(res)
	assert.Equal(t, "invalid range", err.Error())
}

func Test_WhenValidInputRangeDuringAddToRuneGuaranteed_ThenSuccessAggregation(t *testing.T) {
	ro := NewRunOperator()
	res := ro.AddToRuneGuaranteed('A', 3)
	ro.FmtRune(res)
	assert.Equal(t, 'D', res)

	res = ro.AddToRuneGuaranteed('A', 4)
	ro.FmtRune(res)
	assert.Equal(t, 'E', res)

	res = ro.AddToRuneGuaranteed('a', 3)
	ro.FmtRune(res)
	assert.Equal(t, 'd', res)

	res = ro.AddToRuneGuaranteed('a', 500)
	ro.FmtRune(res)
	assert.Equal(t, 'a', res)

	res = ro.AddToRuneGuaranteed('a', 27)
	ro.FmtRune(res)
	assert.Equal(t, 'c', res)

	res = ro.AddToRuneGuaranteed('c', 3)
	ro.FmtRune(res)
	assert.Equal(t, 'f', res)
}
