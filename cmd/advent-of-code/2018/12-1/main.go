package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	state, rules := InputToState(), InputToRules()
	for i := 0; i < 20; i++ {
		state = Next(state, rules)
	}

	var sum int
	for i := range state {
		sum += i
	}
	fmt.Println(sum)
}

func Next(state lib.Set[int], rules map[string]bool) lib.Set[int] {
	// Determine the rule key for the pot at position n (looks at n-2 to n+2).
	key := func(n int) string {
		var sb strings.Builder
		for i := n - 2; i <= n+2; i++ {
			if state.Contains(i) {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		return sb.String()
	}

	min, max := lib.Min(state.Entries()...), lib.Max(state.Entries()...)

	var next lib.Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[key(i)] {
			next.Add(i)
		}
	}
	return next
}

func InputToState() lib.Set[int] {
	line := lib.InputToLines()[0]
	line = strings.ReplaceAll(line, "initial state: ", "")

	var state lib.Set[int]
	for i, c := range line {
		if c == '#' {
			state.Add(i)
		}
	}
	return state
}

func InputToRules() map[string]bool {
	rules := make(map[string]bool)
	for _, line := range lib.InputToLines()[2:] {
		lhs, rhs, _ := strings.Cut(line, " => ")
		rules[lhs] = rhs == "#"
	}

	return rules
}
