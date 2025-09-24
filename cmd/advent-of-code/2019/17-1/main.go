package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	var grid Set[Point2D]

	// Build the grid.
	var current Point2D
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(),
		Output: func(value int) {
			switch value {
			case '.':
				current = current.Right()
			case '#':
				grid.Add(current)
				current = current.Right()
			case '\n':
				current = Point2D{X: 0, Y: current.Y + 1}
			}
		},
	}
	cpu.Execute()

	// Compute the alignment parameters.
	var sum int
	for p := range grid {
		neighbors := SetFrom(p.OrthogonalNeighbors()...)

		if len(grid.Intersect(neighbors)) == 4 {
			sum += p.X * p.Y
		}
	}
	fmt.Println(sum)
}
