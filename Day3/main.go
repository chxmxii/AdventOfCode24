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
	sum2 := 0
	do := true
	data := clearString(input)
	for _, d := range data {
		if strings.Compare("do()",d) == 0 {
			do = true
		} else if strings.Compare("don't()",d) == 0 {
			do = false
		} else {
			sum += mul(d)
			if do {
				sum2 += mul(d)
			}
		}
	}
	fmt.Println("sum of multiplication is", sum)
	fmt.Println("sum of multiplication is", sum2)

}

func clearString(input string) []string {
	data := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(input,-1)
	return data 
}

func mul(input string) int {
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
	return nums[0] * nums[1]
}