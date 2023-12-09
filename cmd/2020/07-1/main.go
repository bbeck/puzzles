package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	rules := InputToRules()

	var count int
	for color := range rules {
		if Contains(color, "shiny gold", rules) {
			count++
		}
	}
	fmt.Println(count)
}

func Contains(color, target string, rules map[string]Contents) bool {
	for other := range rules[color] {
		if other == target || Contains(other, target, rules) {
			return true
		}
	}

	return false
}

type Contents map[string]int

func InputToRules() map[string]Contents {
	type Rule struct {
		Color    string
		Contents Contents
	}

	rs := aoc.InputLinesTo(2020, 7, func(line string) Rule {
		lhs, rhs, _ := strings.Cut(line, " bags contain ")
		rhs = strings.ReplaceAll(rhs, ".", "")
		rhs = strings.ReplaceAll(rhs, " bags", "")
		rhs = strings.ReplaceAll(rhs, " bag", "")

		contents := make(map[string]int)
		if rhs != "no other" {
			for _, child := range strings.Split(rhs, ", ") {
				parts := strings.SplitN(child, " ", 2)
				contents[parts[1]] = aoc.ParseInt(parts[0])
			}
		}

		return Rule{Color: lhs, Contents: contents}
	})

	rules := make(map[string]Contents)
	for _, rule := range rs {
		rules[rule.Color] = rule.Contents
	}
	return rules
}
