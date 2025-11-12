package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	names, rules := InputToNamesAndRules()
	for _, name := range names {
		if IsValid(name, rules) {
			fmt.Println(name)
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
