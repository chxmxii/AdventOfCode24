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

func loadPuzzel(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	data := strings.Split(string(bytes), "\n")
	return data, nil
}

func parsePuzzel(input []string) ([]Rule, [][]int) {
	data := input
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
			fmt.Sscanf(data[l], "%d|%d", &rule.r1, &rule.r2)
			rules = append(rules, rule)
			break
		case 1:
			p := strings.Split(data[l], ",")
			u := make([]int, len(p))
			for i, v := range p {
				fmt.Sscanf(v, "%d", &u[i])
			}
			updates = append(updates, u)
			break
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
	rules, updates := parsePuzzel(data)
	fmt.Println(rules)
	fmt.Println(updates)
}
