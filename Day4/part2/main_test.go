package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "test_input.txt"
		data, err := getInput(input) 
		require.Nil(t, err)
		require.NotEmpty(t, data)

		grid := [][]string{
			{"M", "M", "S"},
			{"M", "A", "A"},
			{"M", "X", "S"},
		}
		assert.Equal(t, false, outOfBounds(grid, 1, 1))
		assert.Equal(t, true, outOfBounds(grid, 3, 1))

		assert.Equal(t, true, isQualified(grid, 1, 1))
		assert.Equal(t, false, isQualified(grid, 3, 1))
		assert.Equal(t, false, isQualified(grid, 0, 0))

		assert.Equal(t, true, isXMAS(grid, 1, 1))
		assert.Equal(t, false, isXMAS(grid, 0, 0))

		assert.Equal(t, 1, countXMAS(grid))
		assert.Equal(t, 9, countXMAS(createGrid(data)))
	})
}