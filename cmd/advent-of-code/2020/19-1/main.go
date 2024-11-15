package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"regexp"
	"strings"
)

func main() {
	rules, tests := InputToRulesAndTests()
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
	lines := puz.InputToLines()
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
