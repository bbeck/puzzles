package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	template, rules := InputToTemplateAndRules()

	// The string grows way too long to keep building at higher iteration counts.
	// Instead, we'll keep all present pairs in a frequency counter and deal with
	// all pairs of the same letters at one time.
	var counts lib.FrequencyCounter[string]
	for i := 0; i < len(template)-1; i++ {
		counts.Add(template[i : i+2])
	}

	for iter := 0; iter < 40; iter++ {
		var next lib.FrequencyCounter[string]
		for _, entry := range counts.Entries() {
			mid := rules[entry.Value]
			lhs := entry.Value[:1] + mid
			rhs := mid + entry.Value[1:]

			next.AddWithCount(lhs, entry.Count)
			next.AddWithCount(rhs, entry.Count)
		}
		counts = next
	}

	// Count the elements keeping in mind that each is going to be double counted
	// due to appearing as the first and second character in a pair.  The very
	// first and last character of the template will be off by one due to
	// appearing as only one character of a pair.
	var elements lib.FrequencyCounter[string]
	for _, entry := range counts.Entries() {
		elements.AddWithCount(string(entry.Value[0]), entry.Count)
		elements.AddWithCount(string(entry.Value[1]), entry.Count)
	}

	entries := elements.Entries()
	first, last := entries[0], entries[len(entries)-1]
	fmt.Println((first.Count+1)/2 - (last.Count+1)/2)
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
