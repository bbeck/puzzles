package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points := InputToPoints(2018, 6)

	// Determine the bounds of the region we're working in.
	minX, minY, maxX, maxY := aoc.GetBounds(points)

	// Within our bounding box, determine the coordinate closest to each point.
	// If there are multiple, then don't assign a closest point.
	closestTo := make(map[aoc.Point2D]aoc.Point2D)
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			cell := aoc.Point2D{X: x, Y: y}

			var closest = -1
			var distance = math.MaxInt64
			for i, p := range points {
				d := cell.ManhattanDistance(p)
				if d == distance {
					// We just found a 2nd point that has the same distance as our best
					// so far, so we have to clear the point since neither can win.  We
					// leave the distance alone though, because we can still do better.
					closest = -1
				}

				if d < distance {
					closest = i
					distance = d
				}
			}

			if closest != -1 {
				closestTo[cell] = points[closest]
			}
		}
	}

	// For all non-infinite points determine how many cells are closest to them.
	// An infinite point is one that has an x value equal to the minimum or
	// maximum x or a y value equal to the minimum or maximum y.
	var best aoc.Point2D
	var bestCount int
	for _, p := range points {
		if p.X == minX || p.X == maxX || p.Y == minY || p.Y == maxY {
			continue
		}

		var count int
		for _, other := range closestTo {
			if p == other {
				count++
			}
		}

		if count > bestCount {
			best = p
			bestCount = count
		}
	}

	fmt.Printf("point %s has an area of %d\n", best, bestCount)
}

func InputToPoints(year, day int) []aoc.Point2D {
	var points []aoc.Point2D
	for _, line := range aoc.InputToLines(year, day) {
		var x, y int
		if _, err := fmt.Sscanf(line, "%d, %d", &x, &y); err != nil {
			log.Fatalf("unable to parse point: %s", line)
		}

		points = append(points, aoc.Point2D{X: x, Y: y})
	}

	return points
}
