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
		gridInput := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM"
		expectedGrid := [][]string{
			{"M","M","M","S","X","X","M","A","S","M"},
			{"M","S","A","M","X","M","S","M","S","A"},
			{"A","M","X","S","X","M","A","A","M","M"},
		}
		// Check if the createdGrid
		assert.Equal(t,expectedGrid ,createGrid(gridInput), "failed to create grid" )
		// find word occurrences
		grid := createGrid(data)
		assert.Equal(t, 18, searchWord(expectedGrid, "XMAS"), "failed to get the correct num of words")
		assert.Equal(t, 1, crossSearchWord(expectedGrid, "MAS"), "failed to get the correct num of cross words")
	})
}