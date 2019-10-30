package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	state := InputToState(2016, 16)
	length := 35651584

	for len(state) < length {
		state = state.Expand()
	}

	fmt.Printf("checksum: %s\n", state.Checksum(length))
}

type State []bool

func (s State) Expand() State {
	expanded := s
	expanded = append(expanded, false)
	for i := len(s) - 1; i >= 0; i-- {
		expanded = append(expanded, !s[i])
	}

	return expanded
}

func (s State) Checksum(n int) State {
	s = s[:n]

	var checksum State
	for {
		for i := 0; i < len(s)-1; i = i + 2 {
			a, b := s[i], s[i+1]
			if a == b {
				checksum = append(checksum, true)
			} else {
				checksum = append(checksum, false)
			}
		}

		if len(checksum)%2 == 1 {
			break
		}

		s = checksum
		checksum = nil
	}

	return checksum
}

func (s State) String() string {
	var builder strings.Builder
	for _, b := range s {
		if b {
			builder.WriteRune('1')
		} else {
			builder.WriteRune('0')
		}
	}

	return builder.String()
}

func InputToState(year, day int) State {
	var state State
	for _, c := range aoc.InputToString(year, day) {
		if c == '0' {
			state = append(state, false)
		} else {
			state = append(state, true)
		}
	}

	return state
}
