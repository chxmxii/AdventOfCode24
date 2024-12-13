package main


import (
	"fmt"
	"strings"
	"os"
)

func getInput(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	data := strings.TrimRight(string(bytes), "\n")
	return data, nil
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
	fmt.Println(data)
}