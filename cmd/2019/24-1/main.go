package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = 5

func main() {
	state := InputToState(2019, 24)

	seen := map[int]bool{
		state.Hash(): true,
	}
	for tm := 1; ; tm++ {
		state = state.Next()
		hash := state.Hash()

		if seen[hash] {
			fmt.Println("biodiversity:", hash)
			break
		}
		seen[hash] = true
	}
}

type State map[aoc.Point2D]bool

func (s State) Next() State {
	neighbors := func(p aoc.Point2D) int {
		var count int
		if s[p.Up()] {
			count++
		}
		if s[p.Down()] {
			count++
		}
		if s[p.Left()] {
			count++
		}
		if s[p.Right()] {
			count++
		}
		return count
	}

	next := make(State)
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			p := aoc.Point2D{X: x, Y: y}
			ns := neighbors(p)
			switch {
			case s[p] && ns != 1:
				next[p] = false
			case !s[p] && ns == 1:
				next[p] = true
			case !s[p] && ns == 2:
				next[p] = true
			default:
				next[p] = s[p]
			}
		}
	}
	return next
}

func (s State) Hash() int {
	hash := 0
	for p, value := range s {
		index := p.Y*N + p.X
		if value {
			hash |= 1 << index
		}
	}
	return hash
}

func (s State) String() string {
	var builder strings.Builder
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if s[aoc.Point2D{X: x, Y: y}] {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func InputToState(year, day int) State {
	state := make(State)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			if c == '#' {
				state[aoc.Point2D{X: x, Y: y}] = true
			}
		}
	}
	return state
}
