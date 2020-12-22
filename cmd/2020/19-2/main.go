package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	rules, messages := Input(2020, 19)

	// Apply overrides from puzzle statement.
	rules[8].clauses = [][]int{
		{42},
		{42, 42},
		{42, 42, 42},
		{42, 42, 42, 42},
		{42, 42, 42, 42, 42},
	}

	rules[11].clauses = [][]int{
		{42, 31},
		{42, 42, 31, 31},
		{42, 42, 42, 31, 31, 31},
		{42, 42, 42, 42, 31, 31, 31, 31},
		{42, 42, 42, 42, 42, 31, 31, 31, 31, 31},
	}

	regex := regexp.MustCompile("^" + ToRegexp(rules, 0) + "$")

	var sum int
	for _, message := range messages {
		if regex.MatchString(message) {
			sum++
		}
	}
	fmt.Println(sum)
}

// Walk a tree of rules converting them into a regular expression string.
func ToRegexp(rules map[int]*Rule, id int) string {
	rule := rules[id]

	if rule.literal != 0 {
		return string(rule.literal)
	}

	// Each clause is a potential sequence of rules that need to be met one after
	// the other, thus we'll concatenate the rules within a clause together.  When
	// there are multiple clauses only one has to match at a time, so we'll OR
	// different clauses together.
	var clauses []string
	for _, clause := range rule.clauses {
		var s string
		for _, part := range clause {
			s += ToRegexp(rules, part)
		}
		clauses = append(clauses, s)
	}

	return "(?:" + strings.Join(clauses, "|") + ")"
}

type Rule struct {
	literal byte
	clauses [][]int
}

func Input(year, day int) (map[int]*Rule, []string) {
	rules := make(map[int]*Rule)
	messages := make([]string, 0)

	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			continue
		}

		if strings.Contains(line, ":") {
			colon := strings.IndexRune(line, ':')
			pipe := strings.IndexRune(line, '|')
			quote := strings.IndexRune(line, '"')

			id := aoc.ParseInt(line[:colon])

			// If there's a quote character then the RHS is a literal.
			if quote != -1 {
				rules[id] = &Rule{literal: line[quote+1]}
				continue
			}

			// There was no quote, so this is a clause based rule.
			var clauses [][]int
			if pipe == -1 {
				clauses = [][]int{
					ParseInts(line[colon+1:]),
				}
			} else {
				clauses = [][]int{
					ParseInts(line[colon+1 : pipe]),
					ParseInts(line[pipe+1:]),
				}
			}

			rules[id] = &Rule{clauses: clauses}
			continue
		}

		messages = append(messages, line)
	}

	return rules, messages
}

// ParseInts parses a string containing one or more space separated integers
// into the integer values.
func ParseInts(s string) []int {
	var ns []int
	for _, n := range strings.Split(s, " ") {
		if len(n) > 0 {
			ns = append(ns, aoc.ParseInt(n))
		}
	}

	return ns
}
