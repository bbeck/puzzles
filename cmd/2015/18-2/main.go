package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	lights := InputToLights()
	TurnOnCorners(lights)

	for i := 0; i < 100; i++ {
		lights = Next(lights)
		TurnOnCorners(lights)
	}

	fmt.Println(len(lights))
}

func TurnOnCorners(lights aoc.Set[aoc.Point2D]) {
	lights.Add(
		aoc.Point2D{X: 0, Y: 0},
		aoc.Point2D{X: 99, Y: 0},
		aoc.Point2D{X: 0, Y: 99},
		aoc.Point2D{X: 99, Y: 99},
	)
}

func Next(lights aoc.Set[aoc.Point2D]) aoc.Set[aoc.Point2D] {
	var next aoc.Set[aoc.Point2D]
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			p := aoc.Point2D{X: x, Y: y}

			var count int
			for _, neighbor := range p.Neighbors() {
				if lights.Contains(neighbor) {
					count++
				}
			}

			// If light==on and count in (2, 3)
			// If light==off and count==3
			if count == 3 || (lights.Contains(p) && count == 2) {
				next.Add(p)
			}
		}
	}

	return next
}

func InputToLights() aoc.Set[aoc.Point2D] {
	var lights aoc.Set[aoc.Point2D]
	for y, line := range aoc.InputToLines(2015, 18) {
		for x, c := range line {
			if c == '#' {
				lights.Add(aoc.Point2D{X: x, Y: y})
			}
		}
	}

	return lights
}
