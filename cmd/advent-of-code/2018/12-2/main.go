package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	state, rules := InputToStateAndRules()

	// We can observe by looking at the first set of iterations that the state
	// eventually gets into a repeating pattern.  The number of pots doesn't
	// change, but the pattern slides to the right a step each time.  We'll
	// iterate until we determine we've repeated a state and then from there
	// we'll calculate how many more steps we need to move to finish the
	// evolution.
	var seen Set[string]
	var n int
	for {
		state = Next(state, rules)
		n++

		if !seen.Add(Key(state)) {
			break
		}
	}

	var sum int
	for i := range state {
		sum += i + (50_000_000_000 - n)
	}
	fmt.Println(sum)
}

func Next(state Set[int], rules map[string]bool) Set[int] {
	min, max := Min(state.Entries()...), Max(state.Entries()...)

	var next Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[RangeKey(state, i-2, i+2)] {
			next.Add(i)
		}
	}
	return next
}

func RangeKey(state Set[int], min, max int) string {
	var sb strings.Builder
	for i := min; i <= max; i++ {
		if state.Contains(i) {
			sb.WriteRune('#')
		} else {
			sb.WriteRune('.')
		}
	}
	return sb.String()
}

func Key(state Set[int]) string {
	min, max := Min[int](state.Entries()...), Max[int](state.Entries()...)
	return RangeKey(state, min, max)
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
