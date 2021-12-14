package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	template, rules := InputToTemplateAndRules()

	for step := 0; step < 10; step++ {
		var sb strings.Builder

		// Take the characters from the template in pairs, looking up the character
		// to insert between them from the rules.  When emitting the new string it's
		// important to not emit the RHS of the pair.  If we emitted it then we'd
		// be duplicating that character since it's about to be the LHS of the next
		// pair that's processed.
		for i := 0; i < len(template)-1; i++ {
			lhs, rhs := string(template[i]), string(template[i+1])
			insertion := rules[lhs+rhs]

			sb.WriteString(lhs)
			sb.WriteString(insertion)
		}

		// The very last character of the template never appears as an LHS, because
		// of this it was lost from the template.  This fixes that and adds it back.
		sb.WriteByte(template[len(template)-1])

		template = sb.String()
	}

	counts := make(map[string]int)
	for _, c := range template {
		s := string(c)
		counts[s]++
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
