package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"regexp"
	"strings"
)

func main() {
	rules, tests := InputToRulesAndTests()

	// Apply overrides from the puzzle statement.  Because this introduces cycles
	// we can't introduce them as defined as that would cause ToRegex to
	// infinitely loop.  Instead, we'll manually unroll them to prevent the
	// infinite recursion.
	//
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	rules["8"] = []Clause{
		{"42"},
		{"42", "42"},
		{"42", "42", "42"},
		{"42", "42", "42", "42"},
		{"42", "42", "42", "42", "42"},
	}
	rules["11"] = []Clause{
		{"42", "31"},
		{"42", "42", "31", "31"},
		{"42", "42", "42", "31", "31", "31"},
		{"42", "42", "42", "42", "31", "31", "31", "31"},
		{"42", "42", "42", "42", "42", "31", "31", "31", "31", "31"},
	}
	regex := regexp.MustCompile("^" + ToRegex(rules, "0") + "$")

	var count int
	for _, test := range tests {
		if regex.MatchString(test) {
			count++
		}
	}
	fmt.Println(count)
}

func ToRegex(rules map[string][]Clause, id string) string {
	var clauses []string
	for _, c := range rules[id] {
		var sb strings.Builder
		for _, cid := range c {
			if cid == "a" || cid == "b" {
				sb.WriteString(cid)
			} else {
				sb.WriteString(ToRegex(rules, cid))
			}
		}

		clauses = append(clauses, sb.String())
	}

	return "(" + strings.Join(clauses, "|") + ")"
}

type Clause []string

func InputToRulesAndTests() (map[string][]Clause, []string) {
	lines := aoc.InputToLines(2020, 19)
	rules := make(map[string][]Clause)

	var index int
	for ; lines[index] != ""; index++ {
		id, rhs, _ := strings.Cut(lines[index], ": ")
		rhs = strings.ReplaceAll(rhs, "\"", "")

		var clauses []Clause
		for _, clause := range strings.Split(rhs, " | ") {
			clauses = append(clauses, strings.Fields(clause))
		}

		rules[id] = clauses
	}

	index++

	return rules, lines[index:]
}
