package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points := InputToPoints(2018, 6)

	// Determine the bounds of the region we're working in.
	minX, minY, maxX, maxY := aoc.GetBounds(points)

	// Expand the bounds by 10k and figure out how many points are within a
	// distance of 10k from every point.
	var size int
	for x := minX - 10000; x < maxX+10000; x++ {
	nextPoint:
		for y := minY - 10000; y < maxY+10000; y++ {
			p := aoc.Point2D{X: x, Y: y}

			var total int
			for _, other := range points {
				total += p.ManhattanDistance(other)
				if total >= 10000 {
					continue nextPoint
				}
			}

			if total < 10000 {
				size++
			}
		}
	}

	fmt.Printf("size of the region: %d\n", size)
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
