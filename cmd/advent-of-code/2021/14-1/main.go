package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	template, rules := InputToTemplateAndRules()

	for iter := 0; iter < 10; iter++ {
		var sb strings.Builder
		for i := 0; i < len(template)-1; i++ {
			lhs, mid := template[i:i+1], rules[template[i:i+2]]
			sb.WriteString(lhs)
			sb.WriteString(mid)
		}

		sb.WriteString(template[len(template)-1:])
		template = sb.String()
	}

	var fc lib.FrequencyCounter[rune]
	for _, c := range template {
		fc.Add(c)
	}

	entries := fc.Entries()
	fmt.Println(entries[0].Count - entries[len(entries)-1].Count)
}

func InputToTemplateAndRules() (string, map[string]string) {
	lines := lib.InputToLines()

	template := lines[0]

	rules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		lhs, rhs, _ := strings.Cut(lines[i], " -> ")
		rules[lhs] = rhs
	}
	return template, rules
}
