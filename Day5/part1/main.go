package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	r1, r2 int
}

var filename = flag.String("f", "puzzel.input", "Filename of the input data")

func (r *Rule) rule(str string) {
	fmt.Sscanf(str, "%d|%d", &r.r1, &r.r2)
}

func loadPuzzel(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(bytes), "\n"), nil
}

func parseData(data []string) ([]Rule, [][]int, error) {
	rules := []Rule{}
	updates := [][]int{}

	d := false

	for _, line := range data {
		line = strings.TrimSpace(line)
		if line == "" {
			d = true
			continue
		}
		if d {
			update := []int{}
			for _, value := range strings.Split(line, ",") {
				v, err := strconv.Atoi(value)
				if err != nil {
					return nil, nil, err
				}
				update = append(update, v)
			}
			updates = append(updates, update)
		} else {
			rule := Rule{}
			rule.rule(line)
			rules = append(rules, rule)
		}
	}

	return rules, updates, nil
}

func getMiddle(pages []int) int {
	if len(pages)%2 == 0 {
		return pages[len(pages)/2-1]
	} else {
		return pages[len(pages)/2]
	}
}

func solve(rules []Rule, updates [][]int) int {
	count := 0
	page := p()
	page.rulesMapping(rules)
	for _, update := range updates {
		if page.checkOrder(update) {
			count += getMiddle(update)
		}
	}
	return count
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	data, err := loadPuzzel(*filename)
	handleError(err)
	rules, updates, err := parseData(data)
	handleError(err)
	fmt.Println(solve(rules, updates))
}
