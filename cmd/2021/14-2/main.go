package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"math"
)

func main() {
	template, rules := InputToTemplateAndRules()

	counts := make(map[string]int)
	for _, c := range template {
		counts[string(c)]++
	}

	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for step := 1; step <= 40; step++ {
		next := make(map[string]int)
		for pair, count := range pairs {
			lhs, rhs := string(pair[0]), string(pair[1])

			middle := rules[pair]
			next[lhs+middle] += count
			next[middle+rhs] += count
			counts[middle] += count
		}
		pairs = next
	}

	var min = math.MaxInt
	var max = 0
	for _, count := range counts {
		min = aoc.MinInt(min, count)
		max = aoc.MaxInt(max, count)
	}
	fmt.Println(max - min)
}

func InputToTemplateAndRules() (string, map[string]string) {
	lines := aoc.InputToLines(2021, 14)

	template := lines[0]

	rules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		var lhs, rhs string
		if _, err := fmt.Sscanf(lines[i], "%s -> %s", &lhs, &rhs); err != nil {
			log.Fatal(err)
		}
		rules[lhs] = rhs
	}
	return template, rules
}
