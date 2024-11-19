package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	depth, target := InputToParameters()

	cave := lib.NewGrid2D[int](target.X+1, target.Y+1)
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

func InputToParameters() (int, lib.Point2D) {
	var depth int
	var target lib.Point2D

	for _, line := range lib.InputToLines() {
		k, v, _ := strings.Cut(line, ": ")
		if k == "depth" {
			depth = lib.ParseInt(v)
		} else if k == "target" {
			x, y, _ := strings.Cut(v, ",")
			target = lib.Point2D{X: lib.ParseInt(x), Y: lib.ParseInt(y)}
		}
	}

	return depth, target
}
