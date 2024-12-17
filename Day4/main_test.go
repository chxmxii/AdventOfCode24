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

		require.Nil(t, err, "Failed to read input")
		require.NotEmpty(t, data, "should not be empty")

		gridInput := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM"
		expectedGrid := [][]string{
			{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
			{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
			{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		}
		createdGrid := createGrid(gridInput)
		assert.Equal(t, expectedGrid, createdGrid, "Failed to create grid")
		grid := createGrid(data) 
		assert.Equal(t, 18, searchWord(grid, "XMAS"), "Wrong count!!")
	})
}