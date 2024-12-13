package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// Testing input data
		input := "test_input.txt"
		data, err := getInput(input)

		assert.Nil(t, err, "Passed")
		require.Nil(t, err, "Required")
		assert.NotEmpty(t, data, "Should not be empty")

		// searching word
		assert.Equal(t, )
	})
}