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
			if lights.GetXY(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func TurnOnCorners(lights aoc.Grid2D[bool]) {
	lights.AddXY(0, 0, true)
	lights.AddXY(lights.Width-1, 0, true)
	lights.AddXY(0, lights.Height-1, true)
	lights.AddXY(lights.Width-1, lights.Height-1, true)
}

func Next(lights aoc.Grid2D[bool]) aoc.Grid2D[bool] {
	next := aoc.NewGrid2D[bool](lights.Width, lights.Height)
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			p := aoc.Point2D{X: x, Y: y}

			var count int
			for _, neighbor := range p.Neighbors() {
				if lights.InBounds(neighbor) && lights.Get(neighbor) {
					count++
				}
			}

			// If light==on and count in (2, 3)
			// If light==off and count==3
			if count == 3 || (lights.Get(p) && count == 2) {
				next.Add(p, true)
			}
		}
	}

	return next
}

func InputToLights() aoc.Grid2D[bool] {
	return aoc.InputToGrid2D(2015, 18, func(x, y int, s string) bool {
		return s == "#"
	})
}
