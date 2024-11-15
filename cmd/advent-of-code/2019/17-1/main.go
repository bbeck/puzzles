package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	var grid puz.Set[puz.Point2D]

	// Build the grid.
	var current puz.Point2D
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 17),
		Output: func(value int) {
			switch value {
			case '.':
				current = current.Right()
			case '#':
				grid.Add(current)
				current = current.Right()
			case '\n':
				current = puz.Point2D{X: 0, Y: current.Y + 1}
			}
		},
	}
	cpu.Execute()

	// Compute the alignment parameters.
	var sum int
	for p := range grid {
		neighbors := puz.SetFrom(p.OrthogonalNeighbors()...)

		if len(grid.Intersect(neighbors)) == 4 {
			sum += p.X * p.Y
		}
	}
	fmt.Println(sum)
}
