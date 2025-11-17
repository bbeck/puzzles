package main

import "github.com/bbeck/puzzles/lib/in"
import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
)

const TURNS = 20

func main() {
	// Locations of the dragons and sheep after each turn
	var dragons, sheep [TURNS + 1]Set[Point2D]

	// Static locations of the hideouts
	var hideouts Set[Point2D]
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "D" {
			dragons[0].Add(Point2D{X: x, Y: y})
		}
		if s == "S" {
			sheep[0].Add(Point2D{X: x, Y: y})
		}
		if s == "#" {
			hideouts.Add(Point2D{X: x, Y: y})
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

	var count int
	for turn := 1; turn <= TURNS; turn++ {
		// Dragons move
		for d := range dragons[turn-1] {
			dragons[turn].Add(moves(d)...)
		}

		// See which sheep they eat before the sheep move
		eaten1 := sheep[turn-1].Intersect(dragons[turn].Difference(hideouts))

		// Now move the sheep that haven't been eaten
		for s := range sheep[turn-1].Difference(eaten1) {
			sheep[turn].Add(s.Down())
		}

		// Now see which sheep have moved into dragons
		eaten2 := sheep[turn].Intersect(dragons[turn].Difference(hideouts))
		sheep[turn] = sheep[turn].Difference(eaten2)

		count += len(eaten1) + len(eaten2)
	}
	fmt.Println(count)
}
