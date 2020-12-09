package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

var grid map[aoc.Point2D]bool

func main() {
	grid = make(map[aoc.Point2D]bool)
	position := aoc.Point2D{}

	output := func(value int) {
		switch value {
		case '.':
			position = position.Right()

		case '#':
			grid[position] = true
			position = position.Right()

		case '\n':
			position = aoc.Point2D{0, position.Y + 1}
		}
	}

	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 17),
		Output: output,
	}
	cpu.Execute()

	var sum int
	for p := range grid {
		if IsIntersection(p) {
			sum += p.X * p.Y
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func IsIntersection(p aoc.Point2D) bool {
	return grid[p.Up()] && grid[p.Right()] && grid[p.Down()] && grid[p.Left()]
}
