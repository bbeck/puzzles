package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	rules := InputToRules(2020, 7)
	count := Count("shiny gold", rules)

	fmt.Println(count)
}

func Count(color string, rules Rules) int {
	var count int
	for _, child := range rules[color] {
		count += child.count + child.count*Count(child.color, rules)
	}
	return count
}

type Rules map[string][]Contents

type Contents struct {
	count int
	color string
}

func InputToRules(year, day int) Rules {
	rules := make(Rules)
	for _, line := range aoc.InputToLines(year, day) {
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, " bags", "")
		line = strings.ReplaceAll(line, " bag", "")

		fields := strings.Split(line, " contain ")
		color := fields[0]

		rhs := fields[1]
		if rhs == "no other" {
			rules[color] = nil
			continue
		}

		var contents []Contents
		for _, child := range strings.Split(rhs, ", ") {
			parts := strings.SplitN(child, " ", 2)

			contents = append(contents, Contents{
				count: aoc.ParseInt(parts[0]),
				color: parts[1],
			})
		}

		rules[color] = contents
	}

	return rules
}
