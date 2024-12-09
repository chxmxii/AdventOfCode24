package utils

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filename string) ([]int, []int) {
	// read file
	f, err := os.Open(filename)
	check(err)

	//close file l8r
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	//create two lists by reading the input file lbl
	scanner := bufio.NewScanner(f)
	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		line := IntArray(scanner)
		left = append(left, line[0])
		right = append(right, line[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return left, right
}

func ClearString(input *bufio.Scanner) string{
	str := input.Text()
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.TrimSpace(str)
	str = strings.Trim(str, "\t \n")
	return str
}

func StringArray(input *bufio.Scanner) []string {
	line := ClearString(input)
	nums := strings.Split(line, "   ")
	return nums
}

func IntArray(input *bufio.Scanner) []int {
	nums := StringArray(input)
	array := make([]int, 0)
	for _, n := range nums {
		value, _ := strconv.Atoi(n)
		array = append(array, value)
	}
	return array
}