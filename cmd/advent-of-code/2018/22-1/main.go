package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	depth, target := InputToParameters()

	cave := puz.NewGrid2D[int](target.X+1, target.Y+1)
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

			cave.Set(x, y, (geologic+depth)%20183)
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

func InputToParameters() (int, puz.Point2D) {
	var depth int
	var target puz.Point2D

	for _, line := range puz.InputToLines() {
		k, v, _ := strings.Cut(line, ": ")
		if k == "depth" {
			depth = puz.ParseInt(v)
		} else if k == "target" {
			x, y, _ := strings.Cut(v, ",")
			target = puz.Point2D{X: puz.ParseInt(x), Y: puz.ParseInt(y)}
		}
	}

	return depth, target
}
