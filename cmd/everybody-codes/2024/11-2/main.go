package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	rules := InputToRules()

	var sum int
	current := map[string]int{"Z": 1}
	for day := 0; day < 10; day++ {
		sum = 0
		next := make(map[string]int)
		for t, count := range current {
			for _, u := range rules[t] {
				next[u] += count
				sum += count
			}
		}
		current = next
	}
	fmt.Println(sum)
}

func InputToRules() map[string][]string {
	rules := make(map[string][]string)
	for _, line := range InputToLines() {
		lhs, rhs, _ := strings.Cut(line, ":")
		rules[lhs] = strings.Split(rhs, ",")
	}
	return rules
}
