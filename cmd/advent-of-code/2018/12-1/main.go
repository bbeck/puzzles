package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	state, rules := InputToStateAndRules()
	for i := 0; i < 20; i++ {
		state = Next(state, rules)
	}

	var sum int
	for i := range state {
		sum += i
	}
	fmt.Println(sum)
}

func Next(state Set[int], rules map[string]bool) Set[int] {
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

	min, max := Min(state.Entries()...), Max(state.Entries()...)

	var next Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[key(i)] {
			next.Add(i)
		}
	}
	return next
}

func InputToStateAndRules() (Set[int], map[string]bool) {
	var state Set[int]

	var s string
	in.Scanf("initial state: %s", &s)
	for i, ch := range s {
		if ch == '#' {
			state.Add(i)
		}
	}

	// Skip blank line
	in.Line()

	var rules = make(map[string]bool)
	for in.HasNext() {
		var lhs, rhs string
		in.Scanf("%s => %s", &lhs, &rhs)
		rules[lhs] = rhs == "#"
	}

	return state, rules
}
