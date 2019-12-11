package main

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	locations := InputToAsteroidLocations(2019, 10)

	station := aoc.Point2D{30, 34}
	N := 200

	visible := Visible(station, locations)
	if len(visible) < N {
		log.Fatal("not enough visible stations, need to implement looping")
	}

	// Sort the visible asteroids by the angle they make with the y-axis
	sort.Slice(visible, func(i, j int) bool {
		return Angle(station, visible[i]) < Angle(station, visible[j])
	})

	nth := visible[N-1]
	fmt.Printf("the %d-th asteroid vaporized is %s: answer: %d\n", N, nth, nth.X*100+nth.Y)
}

func Angle(station aoc.Point2D, p aoc.Point2D) float64 {
	// We have 2 vectors, both are starting at station.  One goes straight up
	// parallel to the Y-axis, and the other goes to p.  We can then use the
	// dot product to determine the angle between them because we know:
	//
	//    u dot v = |u|*|v|*cos(theta)
	//
	// From this we can derive theta to be:
	//
	//                   / u dot v \
	//    theta = arccos| --------- |
	//                   \ |u|*|v| /
	//
	// Now we determine the values for u and v:
	//    u = <0, -1>  (because y increases as we go down)
	//    v = <p.X - station.X, p.Y - station.Y>
	//
	// Simplifying we see that:
	//
	//    theta = arccos(-v.Y / |v|)
	vx, vy := float64(p.X-station.X), float64(p.Y-station.Y)
	theta := math.Acos(-vy / math.Sqrt(vx*vx+vy*vy))

	// Since this just computes the angle between the vector and the Y-axis we
	// need to adjust a bit to make the angle relative to the Y-axis going
	// clockwise.
	if vx < 0 {
		theta = 2*math.Pi - theta
	}

	return theta
}

func Visible(p aoc.Point2D, locations []aoc.Point2D) []aoc.Point2D {
	// Our approach here will be to compute the slope of the line between the
	// asteroid we're interested in and every other asteroid.  Asteroids that have
	// the same slope have the ability to block one another but don't necessarily
	// block depending on which side of the line they are from our point of
	// interest.  We'll keep track of a list of these per slope/side and then
	// order them by distance to our point of interest to figure out which
	// asteroids are visible.
	byslope := make(map[Slope][]aoc.Point2D)
	for _, q := range locations {
		if p == q {
			continue
		}

		dy, dx := p.Slope(q)

		// Determine which side of p the location q is on.  If q is to the left
		// of p, or they're on a vertical line and q is above p, then we'll consider
		// that the true side.  Otherwise it's the false side.
		side := q.X < p.X || (q.X == p.X && q.Y < p.Y)

		slope := Slope{dy, dx, side}
		byslope[slope] = append(byslope[slope], q)
	}

	// Now sort each slope list by distance from p.  Manhattan distance works fine
	// for this.
	var visible []aoc.Point2D
	for _, qs := range byslope {
		sort.Slice(qs, func(i, j int) bool {
			return p.ManhattanDistance(qs[i]) < p.ManhattanDistance(qs[j])
		})

		visible = append(visible, qs[0])
	}

	return visible
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
