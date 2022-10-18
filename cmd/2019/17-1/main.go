package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	var grid aoc.Set[aoc.Point2D]

	// Build the grid.
	var current aoc.Point2D
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
				current = aoc.Point2D{X: 0, Y: current.Y + 1}
			}
		},
	}
	cpu.Execute()

	// Compute the alignment parameters.
	var sum int
	for _, p := range grid.Entries() {
		var neighbors aoc.Set[aoc.Point2D]
		neighbors.Add(p.OrthogonalNeighbors()...)

		if len(grid.Intersect(neighbors)) == 4 {
			sum += p.X * p.Y
		}
	}
	fmt.Println(sum)
}
