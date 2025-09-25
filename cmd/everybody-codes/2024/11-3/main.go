package main

import (
	"fmt"
	"math"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	rules := InputToRules()

	smallest, largest := math.MaxInt, 0
	for _, t := range Keys(rules) {
		num := Run(rules, t)
		smallest = Min(smallest, num)
		largest = Max(largest, num)
	}

	fmt.Println(largest - smallest)
}

func Run(m map[string][]string, start string) int {
	var sum int
	current := map[string]int{start: 1}
	for day := 0; day < 20; day++ {
		sum = 0
		next := make(map[string]int)
		for t, count := range current {
			for _, u := range m[t] {
				next[u] += count
				sum += count
			}
		}

		current = next
	}

	return sum
}

func InputToRules() map[string][]string {
	rules := make(map[string][]string)
	for in.HasNext() {
		lhs, rhs := in.Cut(":")
		rules[lhs] = strings.Split(rhs, ",")
	}
	return rules
}
