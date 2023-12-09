package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points := InputToPoints()
	delta := 20000 / len(points)
	tl, br := aoc.GetBounds(points)

	var count int
	for y := br.Y - delta; y <= tl.Y+delta; y++ {
		for x := br.X - delta; x <= tl.X+delta; x++ {
			cell := aoc.Point2D{X: x, Y: y}

			var total int
			for i := 0; i < len(points) && total < 10000; i++ {
				total += cell.ManhattanDistance(points[i])
			}

			if total < 10000 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func InputToPoints() []aoc.Point2D {
	return aoc.InputLinesTo(2018, 6, func(line string) aoc.Point2D {
		var p aoc.Point2D
		fmt.Sscanf(line, "%d, %d", &p.X, &p.Y)
		return p
	})
}
