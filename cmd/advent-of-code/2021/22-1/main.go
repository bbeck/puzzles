package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var on lib.Set[lib.Point3D]
	for _, c := range InputToCubes() {
		if lib.Min(c.MinX, c.MinY, c.MinZ) < -50 || lib.Max(c.MaxX, c.MaxY, c.MaxZ) > 50 {
			continue
		}

		for x := c.MinX; x <= c.MaxX; x++ {
			for y := c.MinY; y <= c.MaxY; y++ {
				for z := c.MinZ; z <= c.MaxZ; z++ {
					p := lib.Point3D{X: x, Y: y, Z: z}
					if c.State == "on" {
						on.Add(p)
					} else {
						on.Remove(p)
					}
				}
			}
		}
	}

	fmt.Println(len(on))
}

type Cube struct {
	MinX, MaxX int
	MinY, MaxY int
	MinZ, MaxZ int
	State      string
}

func InputToCubes() []Cube {
	return lib.InputLinesTo[Cube](func(line string) Cube {
		var c Cube
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &c.State, &c.MinX, &c.MaxX, &c.MinY, &c.MaxY, &c.MinZ, &c.MaxZ)
		return c
	})
}
