package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string { return s })

	var A, B, C, S Point2D
	grid.ForEachPoint(func(p Point2D, s string) {
		switch s {
		case "A":
			A = p
		case "B":
			B = p
		case "C":
			C = p
		case "S":
			S = p
		}
	})

	isOkToVisit := func(s State, x, y int) (bool, string) {
		if !grid.InBounds(x, y) {
			return false, ""
		}

		ch := grid.Get(x, y)
		switch {
		case ch == "#":
			return false, ""
		case ch == "B" && s.Visited() < 1:
			return false, ""
		case ch == "C" && s.Visited() < 2:
			return false, ""
		case ch == "S" && s.Visited() < 3:
			return false, ""
		case ch == "S" && s.Altitude() < 10000:
			return false, ""
		default:
			return true, ch
		}
	}

	var seenAtAltitude = make(map[State]int)
	children := func(s State) []State {
		x, y, d, a, v := s.X(), s.Y(), s.Dir(), s.Altitude(), s.Visited()

		// See if we've been to this state before, but at a higher altitude.
		// If so then there's no reason to explore.
		other := s & 0xFFFF_0000_FFFF_FFFF
		if oa, ok := seenAtAltitude[other]; ok && oa > a {
			return nil
		}
		seenAtAltitude[other] = a

		var children []State
		for _, step := range STEPS[d] {
			nx, ny, nd := x+step[0], y+step[1], step[2]
			if ok, ch := isOkToVisit(s, nx, ny); ok {
				children = append(children, NewState(nx, ny, nd, a+dA[ch], max(v, V[ch])))
			}
		}

		return children
	}

	goal := func(s State) bool {
		return s.X() == S.X && s.Y() == S.Y && s.Visited() == 3 && s.Altitude() >= 10000
	}

	cost := func(State, State) int {
		return 1
	}

	// As a heuristic we'll compute the Manhattan distance between the current
	// location and all remaining locations.  First precompute the static
	// distances.
	dAB := A.ManhattanDistance(B)
	dBC := B.ManhattanDistance(C)
	dCS := C.ManhattanDistance(S)
	heuristic := func(s State) int {
		x, y := s.X(), s.Y()

		switch s.Visited() {
		case 0:
			return Abs(x-A.X) + Abs(y-A.Y) + dAB + dBC + dCS
		case 1:
			return Abs(x-B.X) + Abs(y-B.Y) + dBC + dCS
		case 2:
			return Abs(x-C.X) + Abs(y-C.Y) + dCS
		default:
			return Abs(x-S.X) + Abs(y-S.Y)
		}
	}

	var start = NewState(S.X, S.Y, 2, 10000, 0)
	_, duration, _ := AStarSearch(start, children, goal, cost, heuristic)
	fmt.Println(duration)
}

// State is a representation of where the glider is packed into a single 64-bit
// integer.
//
//	X: 8-bits
//	Y: 6-bits
//	Dir: 2-bits
//	Altitude: 16-bits
//	Visited: 2-bits
type State uint64

func NewState(x, y, dir, altitude, visited int) State {
	return State(
		((x & 0xFF) << 56) |
			((y & 0x3F) << 50) |
			((dir & 0x03) << 48) |
			((altitude & 0xFFFF) << 32) |
			((visited & 0x03) << 30),
	)
}

func (s State) X() int        { return int((uint64(s) & 0xFF00_0000_0000_0000) >> 56) }
func (s State) Y() int        { return int((uint64(s) & 0x00FC_0000_0000_0000) >> 50) }
func (s State) Dir() int      { return int((uint64(s) & 0x0003_0000_0000_0000) >> 48) }
func (s State) Altitude() int { return int((uint64(s) & 0x0000_FFFF_0000_0000) >> 32) }
func (s State) Visited() int  { return int((uint64(s) & 0x0000_0000_C000_0000) >> 30) }

var STEPS = map[int][][3]int{
	0: {{-1, 0, 3}, {0, -1, 0}, {1, 0, 1}},
	1: {{0, -1, 0}, {1, 0, 1}, {0, 1, 2}},
	2: {{1, 0, 1}, {0, 1, 2}, {-1, 0, 3}},
	3: {{0, 1, 2}, {-1, 0, 3}, {0, -1, 0}},
}

var dA = map[string]int{".": -1, "+": 1, "-": -2, "A": -1, "B": -1, "C": -1, "S": -1}
var V = map[string]int{"A": 1, "B": 2, "C": 3}
