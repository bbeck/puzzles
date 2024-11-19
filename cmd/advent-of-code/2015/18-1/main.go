package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	lights := InputToLights()
	for i := 0; i < 100; i++ {
		lights = Next(lights)
	}

	var count int
	lights.ForEach(func(x, y int, value bool) {
		if value {
			count++
		}
	})
	fmt.Println(count)
}

func Next(lights lib.Grid2D[bool]) lib.Grid2D[bool] {
	next := lib.NewGrid2D[bool](lights.Width, lights.Height)
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

func InputToLights() lib.Grid2D[bool] {
	return lib.InputToGrid2D(func(x, y int, s string) bool {
		return s == "#"
	})
}
