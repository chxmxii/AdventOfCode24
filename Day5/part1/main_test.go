package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve5(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data, err := loadPuzzel("test.input")
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
		require.Nil(t, err)

		rules, updates, err := parseData(data)
		assert.Nil(t, err)
		assert.NotEmpty(t, rules)
		assert.NotEmpty(t, updates)
		require.Nil(t, err)
	})
}
