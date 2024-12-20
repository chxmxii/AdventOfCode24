package main

import (
	"github.com/chxmxii/aoc-d1/utils"
	"fmt"
	"sort"
	"flag"
)

var (
	filename =  flag.String("f","input.txt","input file")
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

func Frequency(right []int, num int) int {
	c := 0
	for _, item := range right {
		if item == num {
			c++
		}
	} 
	return c
}

func main() {
	flag.Parse()
	left, right := utils.ReadFile(*filename)
	// left := []int{3,4,2,1,3,3}
	// right := []int{4,3,5,3,9,3}
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := 0; i < len(left); i++ {
		sum += Abs(left[i] - right[i])
	}
	fmt.Println("Distance is:",sum)
	// part II
	ss := 0
	for i := 0; i < len(left); i++ {
		ss += left[i] * Frequency(right,left[i])
	}
	fmt.Println("Similarty Score is",ss)
}