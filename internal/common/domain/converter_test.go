package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSafeStr(t *testing.T) {
	t.Run("Nil string pointer", func(t *testing.T) {
		var str *string
		result := ToSafeStr(str)
		assert.Equal(t, "", result, "Expected empty string for nil input")
	})

	t.Run("Non-nil string pointer", func(t *testing.T) {
		value := "hello"
		str := &value
		result := ToSafeStr(str)
		assert.Equal(t, "hello", result, "Expected the value of the string pointer")
	})
}

func TestToSafeFloat64(t *testing.T) {
	t.Run("Nil float64 pointer", func(t *testing.T) {
		var f *float64
		result := ToSafeFloat64(f)
		assert.Equal(t, 0.0, result, "Expected 0.0 for nil input")
	})

	t.Run("Non-nil float64 pointer", func(t *testing.T) {
		value := 42.0
		f := &value
		result := ToSafeFloat64(f)
		assert.Equal(t, 42.0, result, "Expected the value of the float64 pointer")
	})
}
