package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
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

func Next(state puz.Set[int], rules map[string]bool) puz.Set[int] {
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

	min, max := puz.Min(state.Entries()...), puz.Max(state.Entries()...)

	var next puz.Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[key(i)] {
			next.Add(i)
		}
	}
	return next
}

func InputToState() puz.Set[int] {
	line := puz.InputToLines(2018, 12)[0]
	line = strings.ReplaceAll(line, "initial state: ", "")

	var state puz.Set[int]
	for i, c := range line {
		if c == '#' {
			state.Add(i)
		}
	}
	return state
}

func InputToRules() map[string]bool {
	rules := make(map[string]bool)
	for _, line := range puz.InputToLines(2018, 12)[2:] {
		lhs, rhs, _ := strings.Cut(line, " => ")
		rules[lhs] = rhs == "#"
	}

	return rules
}
