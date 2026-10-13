package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var snails = InputToSnails()

	// Set up a system of equations that look like:
	//   1 ≡ pos_i + tm (mod len_i)
	//
	// The idea is that the snails move in a cycle so we want the y coordinate to
	// be equal to 1 at some future time relative to their current position.  Once
	// we have the system setup we can use the chinese remainder theorem to solve
	// it.
	var as, ns []int
	for _, snail := range snails {
		pos, length := snail.X, snail.X+snail.Y-1
		as = append(as, 1-pos)
		ns = append(ns, length)
	}

	x := ChineseRemainderTheorem(as, ns)
	fmt.Println(x - 1)
}

func InputToSnails() []Point2D {
	return in.LinesToS(func(in in.Scanner[Point2D]) Point2D {
		return Point2D{X: in.Int(), Y: in.Int()}
	})
}
