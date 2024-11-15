package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	rules := InputToRules()
	fmt.Println(Count("shiny gold", rules))
}

func Count(color string, rules map[string]Contents) int {
	var count int
	for other, n := range rules[color] {
		count += n + n*Count(other, rules)
	}
	return count
}

type Contents map[string]int

func InputToRules() map[string]Contents {
	type Rule struct {
		Color    string
		Contents Contents
	}

	rs := puz.InputLinesTo(func(line string) Rule {
		lhs, rhs, _ := strings.Cut(line, " bags contain ")
		rhs = strings.ReplaceAll(rhs, ".", "")
		rhs = strings.ReplaceAll(rhs, " bags", "")
		rhs = strings.ReplaceAll(rhs, " bag", "")

		contents := make(map[string]int)
		if rhs != "no other" {
			for _, child := range strings.Split(rhs, ", ") {
				parts := strings.SplitN(child, " ", 2)
				contents[parts[1]] = puz.ParseInt(parts[0])
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
