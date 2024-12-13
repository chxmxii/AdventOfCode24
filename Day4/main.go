package main


import (
	"fmt"
	"os"
)

func getInput(filename string) (string, err) {
	bytes, err := os.ReadFile(filename)
	
	if err != nil {
		return err
	}
	
	return bytes, nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "test_input.txt"
	data, err := getInput(filename)
	handleError(err)	
}