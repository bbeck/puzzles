package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	state, rules := InputToState(), InputToRules()

	// We can observe by looking at the first set of iterations that the state
	// eventually gets into a repeating pattern.  The number of pots doesn't
	// change, but the pattern slides to the right a step each time.  We'll
	// iterate until we determine we've repeated a state and then from there
	// we'll calculate how many more steps we need to move to finish the
	// evolution.
	var seen puz.Set[string]
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

func Next(state puz.Set[int], rules map[string]bool) puz.Set[int] {
	min, max := puz.Min(state.Entries()...), puz.Max(state.Entries()...)

	var next puz.Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[RangeKey(state, i-2, i+2)] {
			next.Add(i)
		}
	}
	return next
}

func RangeKey(state puz.Set[int], min, max int) string {
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

func Key(state puz.Set[int]) string {
	min, max := puz.Min[int](state.Entries()...), puz.Max[int](state.Entries()...)
	return RangeKey(state, min, max)
}

func InputToState() puz.Set[int] {
	line := puz.InputToLines()[0]
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
	for _, line := range puz.InputToLines()[2:] {
		lhs, rhs, _ := strings.Cut(line, " => ")
		rules[lhs] = rhs == "#"
	}

	return rules
}
