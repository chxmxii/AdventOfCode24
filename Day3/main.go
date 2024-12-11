package main


import (
	"regexp"
	"strings"
	"strconv"
	_ "reflect"
	"fmt"
	_ "embed"
)

//go:embed input.txt
var input string

func main() {
	input := strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input!!!")
	}
	sum := 0
	data := clearString(input)
	for _, d := range data {
		sum += getNums(d)
	}
	fmt.Println(sum)
}

func clearString(input string) []string {
	data := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`).FindAllString(input,-1)
	return data 
}

func getNums(input string) int {
	nums := make([]int, 0)
	values := regexp.MustCompile(`\d+`).FindAllString(input,-1)
	for _, v := range values {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Cannot convert string to int:", err)
			continue
		}
		nums = append(nums,num)
	} 
	return mul(nums[0],nums[1])
}

func mul(x,y int) int {
	return x*y
}