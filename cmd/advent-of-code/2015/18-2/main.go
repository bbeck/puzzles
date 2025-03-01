package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lights := InputToLights()
	TurnOnCorners(lights)

	for range 100 {
		lights = Next(lights)
		TurnOnCorners(lights)
	}

	var count int
	for y := range lights.Height {
		for x := range lights.Width {
			if lights.Get(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func TurnOnCorners(lights Grid2D[bool]) {
	lights.Set(0, 0, true)
	lights.Set(lights.Width-1, 0, true)
	lights.Set(0, lights.Height-1, true)
	lights.Set(lights.Width-1, lights.Height-1, true)
}

func Next(lights Grid2D[bool]) Grid2D[bool] {
	next := NewGrid2D[bool](lights.Width, lights.Height)
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
			next.Set(x, y, true)
		}
	})

	return next
}

func InputToLights() Grid2D[bool] {
	return in.ToGrid2D(func(x, y int, s string) bool {
		return s == "#"
	})
}
