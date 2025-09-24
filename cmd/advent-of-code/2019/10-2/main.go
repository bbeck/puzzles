package main

import (
	"fmt"
	"math"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	locations := InputToAsteroidLocations()

	// This station location was determined from the solution to part 1.
	station := Point2D{X: 30, Y: 34}

	// Treat the station as the origin and determine the angle each location
	// makes with the positive y-axis.
	angles := make(map[Point2D]float64)
	for _, p := range locations {
		dx, dy := p.X-station.X, p.Y-station.Y

		// Add pi/2 to the angle to measure relative to positive y-axis.
		angle := math.Atan2(float64(dy), float64(dx)) + math.Pi/2
		if angle < 0 {
			angle += 2 * math.Pi
		}

		angles[p] = angle
	}

	// Sort the locations by angle.
	sort.Slice(locations, func(i, j int) bool {
		ai, aj := angles[locations[i]], angles[locations[j]]
		if ai != aj {
			return ai < aj
		}

		return station.ManhattanDistance(locations[i]) < station.ManhattanDistance(locations[j])
	})

	// Index the locations in a ring in their order around the station.
	var r Ring[Point2D]
	for _, p := range locations {
		r.InsertAfter(p)
	}
	r.Next()

	// We're now going to traverse around the ring removing elements.  Locations
	// that have the same angle will appear next to each other in the ring, so
	// after each removal we'll advance until we end up on a location with a
	// different angle.  This technically has a bug where it will enter an
	// infinite loop if all remaining locations have the same angle.  In practice
	// this doesn't happen.
	var last Point2D
	for range 200 {
		last = r.Remove()
		for angles[r.Current()] == angles[last] {
			r.Next()
		}
	}
	fmt.Println(last.X*100 + last.Y)
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
