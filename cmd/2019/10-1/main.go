package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	locations := InputToAsteroidLocations(2019, 10)

	var bestVisible int
	var bestLocation aoc.Point2D
	for _, p := range locations {
		visible := NumVisible(p, locations)
		if visible > bestVisible {
			bestVisible = visible
			bestLocation = p
		}
	}

	fmt.Printf("asteroid %+v is the best with %d other asteroids visible\n", bestLocation, bestVisible)
}

func NumVisible(p aoc.Point2D, locations []aoc.Point2D) int {
	// Our approach here will be to compute the slope of the line between the
	// asteroid we're interested in and every other asteroid.  Asteroids that have
	// the same slope have the ability to block one another but don't necessarily
	// block depending on which side of the line they are from our point of
	// interest.
	slopes := make(map[Slope]bool)
	for _, q := range locations {
		if p == q {
			continue
		}

		dy, dx := p.Slope(q)

		// Determine which side of p the location q is on.  If q is to the left
		// of p, or they're on a vertical line and q is above p, then we'll consider
		// that the true side.  Otherwise it's the false side.
		side := q.X < p.X || (q.X == p.X && q.Y < p.Y)

		slopes[Slope{dy, dx, side}] = true
	}

	return len(slopes)
}

type Slope struct {
	dy, dx int
	side   bool
}

func InputToAsteroidLocations(year, day int) []aoc.Point2D {
	var locations []aoc.Point2D
	for y, line := range aoc.InputToLines(year, day) {
		for x, b := range line {
			if b == '#' {
				locations = append(locations, aoc.Point2D{x, y})
			}
		}
	}

	return locations
}
