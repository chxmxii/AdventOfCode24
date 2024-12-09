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

func main() {
	left, right := utils.ReadFile(*filename)
	sort.Ints(left)
	sort.Ints(right)

}