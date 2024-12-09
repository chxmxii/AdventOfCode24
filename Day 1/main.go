package main

import (
	"github.com/chxmxii/aoc-d1/utils"
	"fmt"
	"sort"
)

var (
	filename =  "input.txt"
) 
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(n int) int {
    if n >= 0 {
        return n
    }
    return -n
}

func main() {
	left, right := utils.ReadFile(filename)
	// left := []int{3,4,2,1,3,3}
	// right := []int{4,3,5,3,9,3}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		sum = sum + Abs(distance)
	}
	fmt.Println(sum)
}
