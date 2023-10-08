package utils_test

import (
	"cinemy-auth-api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	t.Run("Valid Email", func(t *testing.T) {
		isValid, err :=utils.IsValidEmail("user@example.com")
		assert.NoError(t, err)
		assert.True(t, isValid)
	})

	t.Run("InValid Email", func(t *testing.T) {
		isValid, err :=utils.IsValidEmail("user.com")
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "invalid mail")
		assert.False(t, isValid)
	})
}