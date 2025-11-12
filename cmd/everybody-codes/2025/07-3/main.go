package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	names, rules := InputToNamesAndRules()

	// Generate all possible 6 character names by growing the valid input names
	// up to length 6.
	names = Generate6(names, rules)

	// Now consider each rule letter as the last letter of a 6 character name.
	// Compute how many valid suffixes can be appended to it to form a name
	// between 7 and 11 characters long.
	counts := CountSuffixes(rules)

	var sum int
	for _, name := range names {
		sum += counts[name[len(name)-1:]]
	}
	fmt.Println(sum)
}

func Generate6(names []string, rules map[string]Set[string]) []string {
	var seen Set[string]

	var valid = func(name string) bool {
		for i := 0; i < len(name)-1; i++ {
			current, next := name[i:i+1], name[i+1:i+2]
			if !rules[current].Contains(next) {
				return false
			}
		}
		return true
	}

	var visit func(name string)
	visit = func(name string) {
		if len(name) == 6 {
			seen.Add(name)
			return
		}

		for suffix := range rules[name[len(name)-1:]] {
			visit(name + suffix)
		}
	}

	for _, name := range names {
		if valid(name) {
			visit(name)
		}
	}
	return seen.Entries()
}

func CountSuffixes(rules map[string]Set[string]) map[string]int {
	var count func(letter string, length int) int
	count = func(letter string, length int) int {
		if length == 0 {
			return 0
		}

		var sum int
		for suffix := range rules[letter] {
			sum += 1 + count(suffix, length-1)
		}
		return sum
	}

	result := make(map[string]int)
	for letter := range rules {
		result[letter] = count(letter, 5)
	}
	return result
}

func InputToNamesAndRules() ([]string, map[string]Set[string]) {
	names := in.Split(",")
	in.Line() // Skip blank line

	rules := make(map[string]Set[string])
	for in.HasNext() {
		lhs, rhs := in.CutS[string](" > ")
		rules[lhs.String()] = SetFrom(rhs.Split(",")...)
	}

	return names, rules
}
