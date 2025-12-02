package main

import (
	"fmt"
	"math"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	gaps := InputToGaps()

	var xs = []int{0}
	for x := range gaps {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	// Go from left to right keeping track of what gaps are reachable.
	var reachable = []Gap{{Start: 0, End: 0}}
	for i := 1; i < len(xs); i++ {
		var next []Gap // The set of reachable gaps at the next x-coordinate.
		x, nx := xs[i-1], xs[i]

		for _, current := range reachable {
			for _, gap := range gaps[nx] {
				g := Gap{
					// The lowest point we can get to is either the start of the new gap
					// or the point we can reach by not flapping at all.
					Start: Max(current.Start-(nx-x), gap.Start),

					// The highest point we can get to is either the end of the new gap
					// or the point we can reach by constantly flapping.
					End: Min(current.End+(nx-x), gap.End),
				}

				// Make sure we ended up with a non-empty gap.
				if g.End < g.Start {
					continue
				}

				next = append(next, g)
			}
		}

		reachable = next
	}

	// Investigate each reachable point and compute the minimum cost to reach it.
	var best = math.MaxInt
	for _, gap := range reachable {
		x := xs[len(xs)-1]
		for y := gap.Start; y <= gap.End; y++ {
			// A point is reachable only if it is below the y=x diagonal and if it has
			// the same parity as the start point (the origin in this case).
			if y > x || (x+y)%2 != 0 {
				continue
			}

			// The cost to get to a point is y flaps to reach altitude, then (x-y)/2
			// flaps to maintain it the rest of the way.
			cost := (x + y) / 2
			best = Min(best, cost)
		}
	}
	fmt.Println(best)
}

type Gap struct {
	Start, End int // Inclusive
}

func InputToGaps() map[int][]Gap {
	var gaps = make(map[int][]Gap)
	for in.HasNext() {
		x, y0, height := in.Int(), in.Int(), in.Int()
		gaps[x] = append(gaps[x], Gap{Start: y0, End: y0 + height - 1})
	}

	return gaps
}
