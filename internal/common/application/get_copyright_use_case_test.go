package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCopyrightUseCase(t *testing.T) {
	// Create a test GetCopyrightUseCase instance
	getCopyrightUseCase := NewGetCopyrightUseCase()

	// Test case 1: GetCopyright method
	pageData := getCopyrightUseCase.GetCopyright()
	assert.NotNil(t, pageData)
	assert.Equal(t, WALLET_COPYRIGHT_YEAR, pageData.Year)
}
