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

		for i := 0; i < len(template)-1; i++ {
			replacement := rules[template[i:i+2]]
			sb.WriteByte(template[i])
			sb.WriteString(replacement)
		}
		sb.WriteByte(template[len(template)-1])

		template = sb.String()
	}

	counts := make(map[int32]int)
	for _, c := range template {
		counts[c]++
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
