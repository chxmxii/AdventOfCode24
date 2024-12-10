package main

import (
	"fmt"
	"strings"
	"regexp"
	"strconv"
	_"reflect"
)

var (
	input = "};//how():mul(422,702)'how()'from()-&when(551,888)from()#mul(694,437)who()<,/ ~)"
	ptrn1 = `mul\([\d]{1,3},[\d]{1,3}\)` //first pattern used to fetch the safe funcs
	ptrn2 = `[\d]{1,3},[\d]{1,3}` // second pattern used to fetch the nums
)

func main() {
	data :=ClearString(input)
	fmt.Println("this is main", IntArray(data))
	// nums := []int
	// sum := 0
	// data := regexp.MustCompile(pattern).FindAllString(test,-1)
	// nums := regexp.MustComplie(pattern2).FindAllString(data,-1)
	// fmt.Println(data)
	// for _,n := range data {
	// 	nums := regexp.MustCompile(pattern2).FindAllString(n,-1)
	// 	fmt.Println(nums)
	// 	// fmt.Println(reflect.TypeOf(nums))
	// }
}

func ClearString(input string) []string {
	nums := make([]string, 0)
	data := regexp.MustCompile(ptrn1).FindAllString(input,-1)
	fmt.Println("this cs ",data)
	for i := range data {
		fmt.Println("this cs ddd",data[i])
		nums := regexp.MustCompile(ptrn2).FindAllString(i,-1)
		fmt.Println("this is cs too", nums)
		return nums
	}
	return nums
}

func IntArray(input []string) []int {
	// nums := ClearString(input)
	array := make([]int, 0)
	fmt.Println("this is intarr",input)
	for _, n := range input {
		fmt.Println("this intarr",strings.Split(n,","))
		str := strings.Split(n,",")
		for _, s := range str {
			value, _ := strconv.Atoi(s)
			fmt.Println("this intarr vv",value)
			array = append(array, value)
		}
	}
	return array
}

func mul(x,y int) int {
	return x*y
}