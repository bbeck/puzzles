package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	points := InputToPoints()
	delta := 20000 / len(points)
	tl, br := GetBounds(points)

	var count int
	for y := br.Y - delta; y <= tl.Y+delta; y++ {
		for x := br.X - delta; x <= tl.X+delta; x++ {
			cell := Point2D{X: x, Y: y}

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

func InputToPoints() []Point2D {
	return in.LinesToS[Point2D](func(in in.Scanner[Point2D]) Point2D {
		return Point2D{X: in.Int(), Y: in.Int()}
	})
}
