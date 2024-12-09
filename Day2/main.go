package main

import (
	"fmt"
	_ "embed"
	"strings"
	"strconv"
)

//go:embed input.txt
var input string

func main() {
	input := strings.TrimRight(input, "\n")
	safes := 0
	tolerated := 0
	if len(input) == 0 {
		panic("empty input!!!")
	}
	for _, report := range strings.Split(input, "\n") {
		report := strings.Fields(report)
		reports := make([]int, len(report))
		for i := range report {
			reports[i], _ = strconv.Atoi(report[i])
		}
		// part 1
		if isSafe(reports) {
			safes++
			continue
		}
		// part 2
		for i := range reports {
			if !isSafe(reports) {
				save := make([]int, 0, len(reports)-1)
				save = append(save, reports[:i]...)
				save = append(save, reports[i+1:]...)
				if isSafe(save) {
					tolerated++
					break
				}
			}
		}
	}
	fmt.Printf("There is %d safe reports, and %d with toleration.",safes,safes + tolerated)
}


func isSafe(reports []int) bool {
	isOk  := true
	isInc := true
	isDec := true
	for i := range len(reports)-1 {
		if reports[i] >= reports[i+1] {
			isInc = false
		} 
		if reports[i] <= reports[i+1] {
			isDec = false
		}
		diff := Abs(reports[i] - reports[i+1])
		if diff < 1 || diff > 3 {
			isOk = false
		}
		if !(isDec || isInc || isOk) {
			break
		}
	}
	return isDec && isOk || isInc && isOk
}


func Abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}