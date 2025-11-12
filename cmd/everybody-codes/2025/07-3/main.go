package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	names, rules := InputToNamesAndRules()

	var seen Set[string]
	for _, name := range names {
		if IsValid(name, rules) {
			Visit(name, rules, &seen)
		}
	}
	fmt.Println(len(seen))
}

func Visit(name string, rules map[string]Set[string], seen *Set[string]) {
	if 7 <= len(name) && len(name) <= 11 {
		if !seen.Add(name) {
			return
		}
	}

	if len(name) < 11 {
		for suffix := range rules[name[len(name)-1:]] {
			Visit(name+suffix, rules, seen)
		}
	}
}

func IsValid(name string, rules map[string]Set[string]) bool {
	for i := 0; i < len(name)-1; i++ {
		current, next := name[i:i+1], name[i+1:i+2]
		if !rules[current].Contains(next) {
			return false
		}
	}
	return true
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
