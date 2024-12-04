package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

var Segments = map[string]int{"A": 1, "B": 2, "C": 3}

func main() {
	meteors := InputToMeteors()

	// The dimensions of the grid we're going to work on.
	_, br := GetBounds(meteors)
	W, H := br.X+1, br.Y+1

	// Transform the meteors to grid coordinates, keeping in mind that they need
	// to be flipped because x=0 is at the top not the bottom.
	for i, m := range meteors {
		meteors[i] = Point2D{X: m.X, Y: H - m.Y - 1}
	}

	// Place the catapults on the grid at the very bottom.
	catapults := map[string]Point2D{
		"A": {X: 0, Y: H - 1},
		"B": {X: 0, Y: H - 2},
		"C": {X: 0, Y: H - 3},
	}

	type Hit struct {
		Time  int
		Power int
	}
	hits := NewGrid2D[Hit](W, H)
	visit := func(p Point2D, tm, power int) {
		h := hits.GetPoint(p)
		if h.Power == 0 || power < h.Power {
			hits.SetPoint(p, Hit{Time: tm, Power: power})
		}
	}

	// We only have to explore half of the grid's width because the asteroids are
	// moving towards us, so anything we shoot will be met in the horizontal
	// middle along the way.
	for power := 1; power < W/2; power++ {
		for id, c := range catapults {
			tm := 0
			p := c

			// 45-degree angle up
			for n := 0; n < power; n++ {
				tm++
				p = Point2D{X: p.X + 1, Y: p.Y - 1}
				visit(p, tm, Segments[id]*power)
			}
			// Horizontal
			for n := 0; n < power; n++ {
				tm++
				p = Point2D{X: p.X + 1, Y: p.Y}
				visit(p, tm, Segments[id]*power)
			}
			// 45-degree angle down
			for p.X < W-1 && p.Y < H-1 {
				tm++
				p = Point2D{X: p.X + 1, Y: p.Y + 1}
				visit(p, tm, Segments[id]*power)
			}
		}
	}

	// Now move each meteor until it reaches a cell where there's a hit.
	var sum int
	for _, m := range meteors {
		tm := 0

		var found bool
		for {
			if m.X <= 0 || m.Y < 0 {
				break
			}

			if h := hits.GetPoint(m); h.Power > 0 && h.Time <= tm {
				sum += h.Power
				found = true
				break
			}

			tm++
			m = Point2D{X: m.X - 1, Y: m.Y + 1}
		}
		if !found {
			fmt.Println("  ", m, "*** WAS NOT SHOT DOWN ***")
		}
	}

	fmt.Println(sum)
}

func InputToMeteors() []Point2D {
	var ps []Point2D
	for _, line := range InputToLines() {
		fields := strings.Fields(line)
		dx, dy := ParseInt(fields[0]), ParseInt(fields[1])
		ps = append(ps, Point2D{X: dx, Y: dy})
	}
	return ps
}
