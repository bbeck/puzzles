package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid, instructions := InputToGridAndInstructions()

	var sum int
	for i := range instructions {
		var best int
		for toss := 1; toss <= grid.Width/2+1; toss++ {
			final := Drop(grid, instructions[i], toss)
			best = Max(best, 2*final-toss)
		}
		sum += best
	}
	fmt.Println(sum)
}

func Drop(grid Grid2D[string], instructions string, toss int) int {
	s2p := func(s int) Point2D { return Point2D{X: 2 * (s - 1), Y: -1} }
	p2s := func(p Point2D) int { return p.X/2 + 1 }

	var p = s2p(toss)
	for p.Y < grid.Height-1 {
		switch grid.GetPoint(p.Down()) {
		case ".":
			p = p.Down()

		case "*":
			var op byte
			op, instructions = instructions[0], instructions[1:]

			if op == 'L' {
				if p.X > 0 {
					p = p.Left()
				} else {
					p = p.Right() // Bounce off the left wall
				}
			} else if op == 'R' {
				if p.X < grid.Width-1 {
					p = p.Right()
				} else {
					p = p.Left() // Bounce off the right wall
				}
			}
		}
	}

	return p2s(p)
}

func InputToGridAndInstructions() (Grid2D[string], []string) {
	chunk := in.ChunkS[string]()
	return chunk.StringGrid2D(), in.Lines()
}
