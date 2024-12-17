package main


import (
	"fmt"
	"strings"
	"os"
	_ "reflect"
)

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

func outOfBounds(grid [][]string, x, y int) bool {
	return x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0])
}

func isQualified(grid [][]string, x, y int) bool {
	switch {
	case outOfBounds(grid, x+1, y+1),
		outOfBounds(grid, x+1, y-1),
		outOfBounds(grid, x-1, y+1),
		outOfBounds(grid, x-1, y-1):
		return false
	}
	return grid[x][y] == "A"
}

func isXMAS(grid [][]string, x, y int) bool {
	if ((grid[x+1][y+1] == "M" && grid[x-1][y-1] == "S") || 
		(grid[x+1][y+1] == "S" && grid[x-1][y-1] == "M")) && 
	   ((grid[x+1][y-1] == "M" && grid[x-1][y+1] == "S") || 
	   	(grid[x+1][y-1] == "S" && grid[x-1][y+1] == "M")) {
		return true
	}
	return false
}

func countXMAS(grid [][]string) int{
	count := 0 
	for x := 0; x <= len(grid); x++ {
		for y := 0; y <= len(grid); y++ {
			if isQualified(grid, x ,y) && isXMAS(grid, x, y) {
				count++
			}
		}
	}
	return count
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
	fmt.Println("the word MAS has appread around", countXMAS(grid))
}