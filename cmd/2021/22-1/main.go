package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var on aoc.Set[aoc.Point3D]
	for _, c := range InputToCubes() {
		if aoc.Min(c.MinX, c.MinY, c.MinZ) < -50 || aoc.Max(c.MaxX, c.MaxY, c.MaxZ) > 50 {
			continue
		}

		for x := c.MinX; x <= c.MaxX; x++ {
			for y := c.MinY; y <= c.MaxY; y++ {
				for z := c.MinZ; z <= c.MaxZ; z++ {
					p := aoc.Point3D{X: x, Y: y, Z: z}
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
	return aoc.InputLinesTo[Cube](2021, 22, func(line string) Cube {
		var c Cube
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &c.State, &c.MinX, &c.MaxX, &c.MinY, &c.MaxY, &c.MinZ, &c.MaxZ)
		return c
	})
}
