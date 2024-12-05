package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	rules, updates := InputToRulesAndUpdates()

	var sum int
	for _, update := range updates {
		if IsCorrect(rules, update) {
			sum += ParseInt(update[len(update)/2])
		}
	}
	fmt.Println(sum)
}

func IsCorrect(rules map[string][]string, update []string) bool {
	var seen Set[string]
	for _, s := range update {
		for _, successor := range rules[s] {
			if seen.Contains(successor) {
				return false
			}
		}
		seen.Add(s)
	}
	return true
}

func InputToRulesAndUpdates() (map[string][]string, [][]string) {
	rules := make(map[string][]string)
	updates := make([][]string, 0)

	for _, line := range InputToLines() {
		switch {
		case strings.Contains(line, "|"):
			lhs, rhs, _ := strings.Cut(line, "|")
			rules[lhs] = append(rules[lhs], rhs)
		case strings.Contains(line, ","):
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return rules, updates
}
