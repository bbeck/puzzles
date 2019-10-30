package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	state, rules := InputToStateAndRules(2018, 12)
	N := 50000000000

	// We can observe by looking at the first set of iterations that the state
	// eventually gets into a repeating pattern.  The number of pots doesn't
	// change, but the pattern slides to the right a step each time.  So keep
	// iterating until we determine that we're repeating, and from there we'll
	// perform a calculation of how many more steps to the right we need to take
	// to finish the evolution.
	seen := make(map[string]bool)

	var i int
	for i = 0; i <= N; {
		state = state.Evolve(rules)
		i++

		hash := state.String()
		if seen[hash] {
			break
		}

		seen[hash] = true
	}

	// Compute the delta that we need to add to each pot's value
	delta := N - i

	var sum int
	for _, pot := range state.Pots() {
		sum += pot + delta
	}
	fmt.Printf("sum: %d\n", sum)
}

// State is a mapping of position to whether or not there's a plant there
type State map[int]bool

func (s *State) Evolve(rules Rules) *State {
	min, max := s.Bounds()

	bit := func(i int) int {
		if (*s)[i] {
			return 1
		}
		return 0
	}

	// We loop over the full bounds + 2 extra pots on each side because each
	// generation of is influenced by 2 pots to the left and 2 pots to the right.
	next := make(State)
	for i := min - 2; i <= max+2; i++ {
		n := (bit(i-2) << 4) | (bit(i-1) << 3) | (bit(i) << 2) | (bit(i+1) << 1) | bit(i+2)

		// Only store pots with plants -- empty pots don't carry any information.
		if rules[n] {
			next[i] = true
		}
	}

	return &next
}

func (s *State) Pots() []int {
	min, max := s.Bounds()

	var indices []int
	for i := min; i <= max; i++ {
		if (*s)[i] {
			indices = append(indices, i)
		}
	}

	return indices
}

func (s *State) Bounds() (int, int) {
	min := math.MaxInt64
	max := math.MinInt64
	for x := range *s {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}

	return min, max
}

func (s *State) String() string {
	min, max := s.Bounds()

	var builder strings.Builder
	for i := min; i <= max; i++ {
		if (*s)[i] {
			builder.WriteRune('#')
		} else {
			builder.WriteRune('.')
		}
	}

	return builder.String()
}

type Rules map[int]bool

func InputToStateAndRules(year, day int) (*State, Rules) {
	lines := aoc.InputToLines(year, day)

	var initial string
	if _, err := fmt.Sscanf(lines[0], "initial state: %s", &initial); err != nil {
		log.Fatalf("unable to parse initial state: %s", lines[0])
	}

	state := make(State)
	for i := 0; i < len(initial); i++ {
		state[i] = initial[i] == '#'
	}

	toN := func(bs ...byte) int {
		var n int
		for _, b := range bs {
			n <<= 1

			if b == '#' {
				n |= 1
			}
		}

		return n
	}

	rules := make(Rules)
	for _, line := range lines[2:] {
		var lhs1, lhs2, lhs3, lhs4, lhs5, rhs byte
		if _, err := fmt.Sscanf(line, "%c%c%c%c%c => %c", &lhs1, &lhs2, &lhs3, &lhs4, &lhs5, &rhs); err != nil {
			log.Fatalf("unable to parse rule: %s", line)
		}

		rules[toN(lhs1, lhs2, lhs3, lhs4, lhs5)] = rhs == '#'
	}

	return &state, rules
}
