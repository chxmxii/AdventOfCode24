package main


import (
	"fmt"
	"strings"
	"os"
	_ "reflect"
)

var directions = [8][2]int{
	{-1, 0}, // up
	{-1, -1}, // up left
	{-1, 1}, // up right
	{1, 1}, // down right
	{1, 0}, // down
	{1, -1}, // down left
	{0, -1}, // left
	{0, 1}, // right
}

func getInput(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}

func createGrid(input string) [][]string {
	lines := strings.Split(input, "\n")
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}
	grid := make([][]string, len(lines))
	for i, line := range lines {
		row := strings.Split(line, "")
		for len(row) < maxLen {
			row = append(row, "")
		}
		grid[i] = row
	}
	return grid
}

func searchWord(grid [][]string, word string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	ans := 0
	for r := 0; r < rows; r++ {
		for c:= 0; c < cols; c++ {
			// search on all directions
			if grid[r][c] == string(word[0]) {
				for _, d := range directions {
					found := true
					for i := 1; i < len(word); i++ {
						nr := r + i*d[0]
						nc := c + i*d[1]

						if nr < 0 || nc < 0 || nr >= rows || nc >= cols || grid[nr][nc] != string(word[i]) {
							found = false
							break
						} 	
					}
					if found {
						ans++
					}
				}
			}
		}
	}
	return ans
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	filename := "input.txt"
	data, err := getInput(filename)
	handleError(err)

	grid := createGrid(data)
	ans := searchWord(grid, "XMAS")
	fmt.Println("the word XMAS has appread around", ans)
}