package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	locations := InputToAsteroidLocations()

	var best int
	for _, p := range locations {
		best = lib.Max(best, NumVisible(p, locations))
	}
	fmt.Println(best)
}

func NumVisible(p lib.Point2D, qs []lib.Point2D) int {
	// Consider p as the center of a coordinate system.  Within each quadrant of
	// the coordinate system only a single point with a slope can be seen by p.
	var slopes [4]lib.Set[Slope]
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

func InputToAsteroidLocations() []lib.Point2D {
	var locations []lib.Point2D
	for y, line := range lib.InputToLines() {
		for x, b := range line {
			if b == '#' {
				locations = append(locations, lib.Point2D{X: x, Y: y})
			}
		}
	}

	return locations
}
