package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	locations := InputToAsteroidLocations()

	var best int
	for _, p := range locations {
		best = aoc.Max(best, NumVisible(p, locations))
	}
	fmt.Println(best)
}

func NumVisible(p aoc.Point2D, qs []aoc.Point2D) int {
	// Consider p as the center of a coordinate system.  Within each quadrant of
	// the coordinate system only a single point with a slope can be seen by p.
	var slopes [4]aoc.Set[Slope]
	for _, q := range qs {
		var quadrant int
		if q.X < p.X {
			quadrant += 1
		}
		if q.Y > p.Y {
			quadrant += 2
		}

		dx, dy := p.Slope(q)
		slopes[quadrant].Add(Slope{dx: dx, dy: dy})
	}

	return len(slopes[0]) + len(slopes[1]) + len(slopes[2]) + len(slopes[3])
}

type Slope struct {
	dy, dx int
}

func InputToAsteroidLocations() []aoc.Point2D {
	var locations []aoc.Point2D
	for y, line := range aoc.InputToLines(2019, 10) {
		for x, b := range line {
			if b == '#' {
				locations = append(locations, aoc.Point2D{X: x, Y: y})
			}
		}
	}

	return locations
}
