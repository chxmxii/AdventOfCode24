package main

import(
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve5(t *testing.T) {
	t.Run("", func(t *testing.T) {
		data, err := loadInput("test.input")
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
		require.Nil(t, err)
		
		assert.Equal(t, {53,47} , parseRule("47|53"))
	})
}