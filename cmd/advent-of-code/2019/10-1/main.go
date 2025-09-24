package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	locations := InputToAsteroidLocations()

	var best int
	for _, p := range locations {
		best = Max(best, NumVisible(p, locations))
	}
	fmt.Println(best)
}

func NumVisible(p Point2D, qs []Point2D) int {
	// Consider p as the center of a coordinate system.  Within each quadrant of
	// the coordinate system only a single point with a slope can be seen by p.
	var slopes [4]Set[Slope]
	for _, q := range qs {
		if p == q {
			continue
		}

		var quadrant int
		if q.X < p.X {
			quadrant += 1
		}
		if q.Y > p.Y {
			quadrant += 2
		}

		dy, dx := p.Slope(q)
		slopes[quadrant].Add(Slope{dx: dx, dy: dy})
	}

	return len(slopes[0]) + len(slopes[1]) + len(slopes[2]) + len(slopes[3])
}

type Slope struct {
	dy, dx int
}

func InputToAsteroidLocations() []Point2D {
	var locations []Point2D
	in.ToGrid2D(func(x, y int, s string) string {
		if s == "#" {
			locations = append(locations, Point2D{X: x, Y: y})
		}
		return s
	})
	return locations
}
