package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	depth, target := InputToParameters()

	cave := aoc.NewGrid2D[int](target.X+1, target.Y+1)
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			var geologic int

			switch {
			case (x == 0 && y == 0) || (x == target.X && y == target.Y):
				geologic = 0
			case x == 0:
				geologic = y * 48271
			case y == 0:
				geologic = x * 16807
			default:
				geologic = cave.Get(x-1, y) * cave.Get(x, y-1)
			}

			cave.Add(x, y, (geologic+depth)%20183)
		}
	}

	var risk int
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			risk += cave.Get(x, y) % 3
		}
	}
	fmt.Println(risk)
}

func InputToParameters() (int, aoc.Point2D) {
	var depth int
	var target aoc.Point2D

	for _, line := range aoc.InputToLines(2018, 22) {
		k, v, _ := strings.Cut(line, ": ")
		if k == "depth" {
			depth = aoc.ParseInt(v)
		} else if k == "target" {
			x, y, _ := strings.Cut(v, ",")
			target = aoc.Point2D{X: aoc.ParseInt(x), Y: aoc.ParseInt(y)}
		}
	}

	return depth, target
}
