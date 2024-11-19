package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	state, rules := InputToState(), InputToRules()

	// We can observe by looking at the first set of iterations that the state
	// eventually gets into a repeating pattern.  The number of pots doesn't
	// change, but the pattern slides to the right a step each time.  We'll
	// iterate until we determine we've repeated a state and then from there
	// we'll calculate how many more steps we need to move to finish the
	// evolution.
	var seen lib.Set[string]
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

func Next(state lib.Set[int], rules map[string]bool) lib.Set[int] {
	min, max := lib.Min(state.Entries()...), lib.Max(state.Entries()...)

	var next lib.Set[int]
	for i := min - 4; i <= max+4; i++ {
		if rules[RangeKey(state, i-2, i+2)] {
			next.Add(i)
		}
	}
	return next
}

func RangeKey(state lib.Set[int], min, max int) string {
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

func Key(state lib.Set[int]) string {
	min, max := lib.Min[int](state.Entries()...), lib.Max[int](state.Entries()...)
	return RangeKey(state, min, max)
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
