package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)


func loadInput(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}

func parseRule(rules string) ([][]int, error) {
	for _, rule := range rules {
		rules, err := strconv.Atoi(strings.Split(input, "|"))
		if err != nil {
			fmt.Println("Couldn't to parse rule", err)
		}
	}
	return rules, nil
}

func parseUpdates(input string) ([][]int, error) {	

}


func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := loadInput("test.input")
	errorHandler(err)

	fmt.Println(data)
}