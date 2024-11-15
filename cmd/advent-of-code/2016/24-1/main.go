package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	var waypoints []puz.Point2D
	grid := puz.InputToGrid2D(func(x, y int, value string) bool {
		if value == "#" {
			return false
		}

		if value != "." {
			p := puz.Point2D{X: x, Y: y}

			if value == "0" {
				// This is our starting location, make sure it's always at the front of
				// the list.
				waypoints = append([]puz.Point2D{p}, waypoints...)
			} else {
				waypoints = append(waypoints, p)
			}
		}

		return true
	})

	// Determine the shortest distance between each of the waypoints.
	distances := puz.Make2D[int](len(waypoints), len(waypoints))
	for i := 0; i < len(waypoints); i++ {
		distances[i][i] = 0
		for j := i + 1; j < len(waypoints); j++ {
			distances[i][j] = distance(waypoints[i], waypoints[j], grid)
			distances[j][i] = distances[i][j]
		}
	}

	// Consider each ordering of the waypoints and determine the one that results
	// in the shortest distance traveled.  Keeping in mind that we always start
	// at the first waypoint in the list.
	var best = math.MaxInt
	puz.EnumeratePermutations(len(waypoints)-1, func(perm []int) bool {
		path := []int{0}
		for _, index := range perm {
			// We add 1 to the index since we're enumerating permutations skipping
			// over the first element which is our current location.  This ensures
			// that we're looking up the correct distances.
			path = append(path, index+1)
		}

		var sum int
		for i := 0; i < len(path)-1; i++ {
			sum += distances[path[i]][path[i+1]]
		}
		best = puz.Min(best, sum)
		return false
	})

	fmt.Println(best)
}

func distance(p1, p2 puz.Point2D, g puz.Grid2D[bool]) int {
	children := func(p puz.Point2D) []puz.Point2D {
		var children []puz.Point2D
		for _, c := range p.OrthogonalNeighbors() {
			if g.GetPoint(c) {
				children = append(children, c)
			}
		}
		return children
	}

	goal := func(p puz.Point2D) bool {
		return p == p2
	}

	path, found := puz.BreadthFirstSearch(p1, children, goal)
	if !found {
		return math.MaxInt
	}
	return len(path) - 1
}
