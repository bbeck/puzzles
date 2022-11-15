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

	var count int
	for y := 0; y < lights.Height; y++ {
		for x := 0; x < lights.Width; x++ {
			if lights.Get(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func TurnOnCorners(lights aoc.Grid2D[bool]) {
	lights.Add(0, 0, true)
	lights.Add(lights.Width-1, 0, true)
	lights.Add(0, lights.Height-1, true)
	lights.Add(lights.Width-1, lights.Height-1, true)
}

func Next(lights aoc.Grid2D[bool]) aoc.Grid2D[bool] {
	next := aoc.NewGrid2D[bool](lights.Width, lights.Height)
	lights.ForEach(func(x, y int, value bool) {
		var count int
		lights.ForEachNeighbor(x, y, func(x, y int, value bool) {
			if value {
				count++
			}
		})

		// If light==on and count in (2, 3)
		// If light==off and count==3
		if count == 3 || (lights.Get(x, y) && count == 2) {
			next.Add(x, y, true)
		}
	})

	return next
}

func InputToLights() aoc.Grid2D[bool] {
	return aoc.InputToGrid2D(2015, 18, func(x, y int, s string) bool {
		return s == "#"
	})
}
