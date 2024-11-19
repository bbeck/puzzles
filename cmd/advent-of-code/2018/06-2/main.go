package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	points := InputToPoints()
	delta := 20000 / len(points)
	tl, br := lib.GetBounds(points)

	var count int
	for y := br.Y - delta; y <= tl.Y+delta; y++ {
		for x := br.X - delta; x <= tl.X+delta; x++ {
			cell := lib.Point2D{X: x, Y: y}

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

func InputToPoints() []lib.Point2D {
	return lib.InputLinesTo(func(line string) lib.Point2D {
		var p lib.Point2D
		fmt.Sscanf(line, "%d, %d", &p.X, &p.Y)
		return p
	})
}
