package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
	"math"
)

func main() {
	template, rules := InputToTemplateAndRules()

	// For part 2 we need to perform 40 steps of the insertion algorithm which will
	// result in a massive string.  Instead of building a string we'll just keep
	// track of the pairs that exist as well as how many times they appear.  This
	// representation loses the ordering of the pairs, but that's not needed in
	// order to compute the final answer.
	//
	// An insertion operation will take one pair, remove it, and add in two new
	// pairs.
	//
	// Finally, as we proceed we'll also keep track of the count of each character
	// that's been used.  This is straightforward as each insertion adds a single
	// character to the overall "string" being generated.

	// Seed the counts with the characters we have in the original template.
	counts := make(map[string]int)
	for _, c := range template {
		counts[string(c)]++
	}

	// Seed the list of pairs from the original template.
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for step := 1; step <= 40; step++ {
		next := make(map[string]int)
		for pair, count := range pairs {
			lhs, rhs := string(pair[0]), string(pair[1])
			insertion := rules[pair]

			// Add the new pairs
			next[lhs+insertion] += count
			next[insertion+rhs] += count

			// Count the new character we just inserted
			counts[insertion] += count
		}
		pairs = next
	}

	min, max := math.MaxInt, math.MinInt
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
