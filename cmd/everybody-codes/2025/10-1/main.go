package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var dragon Point2D
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "D" {
			dragon = Point2D{X: x, Y: y}
		}
		return s
	})

	moves := func(d Point2D) []Point2D {
		var moves []Point2D
		for dx := -2; dx <= 2; dx++ {
			for dy := -2; dy <= 2; dy++ {
				if Abs(dx)+Abs(dy) == 3 && grid.InBounds(d.X+dx, d.Y+dy) {
					moves = append(moves, Point2D{X: d.X + dx, Y: d.Y + dy})
				}
			}
		}
		return moves
	}

	var seen Set[Point2D]
	var current = SetFrom(dragon)
	for range 4 {
		var next Set[Point2D]
		for p := range current {
			for _, m := range moves(p) {
				next.Add(m)

				if grid.GetPoint(m) == "S" {
					seen.Add(m)
				}
			}
		}
		current = next
	}

	fmt.Println(len(seen))
}
