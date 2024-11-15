package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	points := InputToPoints()
	delta := 20000 / len(points)
	tl, br := puz.GetBounds(points)

	var count int
	for y := br.Y - delta; y <= tl.Y+delta; y++ {
		for x := br.X - delta; x <= tl.X+delta; x++ {
			cell := puz.Point2D{X: x, Y: y}

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

func InputToPoints() []puz.Point2D {
	return puz.InputLinesTo(func(line string) puz.Point2D {
		var p puz.Point2D
		fmt.Sscanf(line, "%d, %d", &p.X, &p.Y)
		return p
	})
}
