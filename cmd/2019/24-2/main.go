package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = uint8(5)
const D = 225

func main() {
	state := InputToState(2019, 24)
	for tm := 0; tm < 200; tm++ {
		state = state.Next()
	}

	var count int
	for _, v := range state {
		if v {
			count++
		}
	}

	fmt.Printf("number of bugs: %d\n", count)
}

type Coordinate struct {
	id    uint8
	depth int
}

func (c Coordinate) Neighbors() []Coordinate {
	switch c.id {
	case 1:
		return []Coordinate{
			{8, c.depth - 1},
			{2, c.depth},
			{6, c.depth},
			{12, c.depth - 1},
		}
	case 2:
		return []Coordinate{
			{8, c.depth - 1},
			{3, c.depth},
			{7, c.depth},
			{1, c.depth},
		}
	case 3:
		return []Coordinate{
			{8, c.depth - 1},
			{4, c.depth},
			{8, c.depth},
			{2, c.depth},
		}
	case 4:
		return []Coordinate{
			{8, c.depth - 1},
			{5, c.depth},
			{9, c.depth},
			{3, c.depth},
		}
	case 5:
		return []Coordinate{
			{8, c.depth - 1},
			{14, c.depth - 1},
			{10, c.depth},
			{4, c.depth},
		}
	case 6:
		return []Coordinate{
			{1, c.depth},
			{7, c.depth},
			{11, c.depth},
			{12, c.depth - 1},
		}
	case 7:
		return []Coordinate{
			{2, c.depth},
			{8, c.depth},
			{12, c.depth},
			{6, c.depth},
		}
	case 8:
		return []Coordinate{
			{3, c.depth},
			{9, c.depth},
			{1, c.depth + 1},
			{2, c.depth + 1},
			{3, c.depth + 1},
			{4, c.depth + 1},
			{5, c.depth + 1},
			{7, c.depth},
		}
	case 9:
		return []Coordinate{
			{4, c.depth},
			{10, c.depth},
			{14, c.depth},
			{8, c.depth},
		}
	case 10:
		return []Coordinate{
			{5, c.depth},
			{14, c.depth - 1},
			{15, c.depth},
			{9, c.depth},
		}
	case 11:
		return []Coordinate{
			{6, c.depth},
			{12, c.depth},
			{16, c.depth},
			{12, c.depth - 1},
		}
	case 12:
		return []Coordinate{
			{7, c.depth},
			{1, c.depth + 1},
			{6, c.depth + 1},
			{11, c.depth + 1},
			{16, c.depth + 1},
			{21, c.depth + 1},
			{17, c.depth},
			{11, c.depth},
		}
	case 14:
		return []Coordinate{
			{9, c.depth},
			{15, c.depth},
			{19, c.depth},
			{5, c.depth + 1},
			{10, c.depth + 1},
			{15, c.depth + 1},
			{20, c.depth + 1},
			{25, c.depth + 1},
		}
	case 15:
		return []Coordinate{
			{10, c.depth},
			{14, c.depth - 1},
			{20, c.depth},
			{14, c.depth},
		}
	case 16:
		return []Coordinate{
			{11, c.depth},
			{17, c.depth},
			{21, c.depth},
			{12, c.depth - 1},
		}
	case 17:
		return []Coordinate{
			{12, c.depth},
			{18, c.depth},
			{22, c.depth},
			{16, c.depth},
		}
	case 18:
		return []Coordinate{
			{21, c.depth + 1},
			{22, c.depth + 1},
			{23, c.depth + 1},
			{24, c.depth + 1},
			{25, c.depth + 1},
			{19, c.depth},
			{23, c.depth},
			{17, c.depth},
		}
	case 19:
		return []Coordinate{
			{14, c.depth},
			{20, c.depth},
			{24, c.depth},
			{18, c.depth},
		}
	case 20:
		return []Coordinate{
			{15, c.depth},
			{14, c.depth - 1},
			{25, c.depth},
			{19, c.depth},
		}
	case 21:
		return []Coordinate{
			{16, c.depth},
			{22, c.depth},
			{18, c.depth - 1},
			{12, c.depth - 1},
		}
	case 22:
		return []Coordinate{
			{17, c.depth},
			{23, c.depth},
			{18, c.depth - 1},
			{21, c.depth},
		}
	case 23:
		return []Coordinate{
			{18, c.depth},
			{24, c.depth},
			{18, c.depth - 1},
			{22, c.depth},
		}
	case 24:
		return []Coordinate{
			{19, c.depth},
			{25, c.depth},
			{18, c.depth - 1},
			{23, c.depth},
		}
	case 25:
		return []Coordinate{
			{20, c.depth},
			{14, c.depth - 1},
			{18, c.depth - 1},
			{24, c.depth},
		}
	default:
		log.Fatalf("unable to determine neighbors for: %+v", c)
		return nil
	}
}

type State map[Coordinate]bool

func NewState() State {
	state := make(State)
	for depth := -D; depth < D; depth++ {
		for id := uint8(1); id <= N*N; id++ {
			if id == 13 {
				continue
			}

			state[Coordinate{id, depth}] = false
		}
	}

	return state
}

func (s State) Next() State {
	// Count the number of neighbors of this coordinate that are true
	neighbors := func(c Coordinate) int {
		var count int
		for _, neighbor := range c.Neighbors() {
			if s[neighbor] {
				count++
			}
		}
		return count
	}

	next := NewState()
	for depth := -D; depth < D; depth++ {
		for id := uint8(1); id <= N*N; id++ {
			if id == 13 {
				continue
			}

			coordinate := Coordinate{id, depth}
			count := neighbors(coordinate)

			if s[coordinate] {
				next[coordinate] = count == 1
			} else {
				next[coordinate] = count == 1 || count == 2
			}
		}
	}
	return next
}

func (s State) String() string {
	var builder strings.Builder
	for depth := -D; depth < D; depth++ {
		builder.WriteString(fmt.Sprintf("Depth %d:\n", depth))
		for id := uint8(1); id <= N*N; id++ {
			if id == 13 {
				builder.WriteString("?")
				continue
			}

			coordinate := Coordinate{id, depth}
			if s[coordinate] {
				builder.WriteString("#")
			} else {
				builder.WriteString(".")
			}
			if id%5 == 0 {
				builder.WriteString("\n")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func InputToState(year, day int) State {
	state := NewState()
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			if c == '#' {
				state[Coordinate{id: uint8(y)*N + uint8(x+1), depth: 0}] = true
			}
		}
	}
	return state
}
