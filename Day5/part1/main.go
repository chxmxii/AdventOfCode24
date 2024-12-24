package main

import (
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	r1, r2 int
}

func rule(r1, r2 int) Rule {
	return Rule{r1: r1, r2: r2}
}

func loadPuzzel(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func parsePuzzel(input string) ([]Rule, [][]int) {
	data := strings.Split(string(input), "\n")
	updates := make([][]int, 0, len(data))
	rules := make([]Rule, 0, len(data))

	l := 0
	sec := 0
	for l < len(data) && sec < 2 {
		if data[l] == "" {
			sec++
			l++
			continue
		}
		switch sec {

		case 0:
			rule := Rule{}
			fmt.Sscanf(data[l], "%d|%d\n", &rule.r1, &rule.r2)
			rules = append(rules, rule)

		case 1:
			p := strings.Split(data[l], ",")
			u := make([]int, len(p))
			for i, v := range p {
				fmt.Sscanf(v, "%d", &u[i])
			}
			updates = append(updates, u)
		}
		l++
	}
	return rules, updates
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := loadPuzzel("test.input")
	errHandler(err)
	for _, v := range data {
		fmt.Println(v)
	}
	rules, updates := parsePuzzel(data)
	fmt.Println(rules)
	fmt.Println(updates)
}
