package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	lights := make(map[Point3D]bool)
	for _, i := range InputToCubes() {
		if i.MinX < -50 || i.MinX > 50 || i.MaxX < -50 || i.MaxX > 50 {
			continue
		}
		if i.MinY < -50 || i.MinY > 50 || i.MaxY < -50 || i.MaxY > 50 {
			continue
		}
		if i.MinZ < -50 || i.MinZ > 50 || i.MaxZ < -50 || i.MaxZ > 50 {
			continue
		}

		for x := i.MinX; x <= i.MaxX; x++ {
			for y := i.MinY; y <= i.MaxY; y++ {
				for z := i.MinZ; z <= i.MaxZ; z++ {
					p := Point3D{X: x, Y: y, Z: z}

					if i.On {
						lights[p] = i.On
					} else {
						delete(lights, p)
					}
				}
			}
		}
	}

	fmt.Println(len(lights))
}

type Point3D struct {
	X, Y, Z int
}

type Cube struct {
	MinX, MaxX int
	MinY, MaxY int
	MinZ, MaxZ int
	On         bool
}

func InputToCubes() []Cube {
	var cubes []Cube
	for _, line := range aoc.InputToLines(2021, 22) {
		var state string
		var minX, maxX, minY, maxY, minZ, maxZ int

		if _, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &minX, &maxX, &minY, &maxY, &minZ, &maxZ); err != nil {
			log.Fatal(err)
		}

		cubes = append(cubes, Cube{
			On:   state == "on",
			MinX: aoc.MinInt(minX, maxX),
			MaxX: aoc.MaxInt(minX, maxX),
			MinY: aoc.MinInt(minY, maxY),
			MaxY: aoc.MaxInt(minY, maxY),
			MinZ: aoc.MinInt(minZ, maxZ),
			MaxZ: aoc.MaxInt(minZ, maxZ),
		})
	}

	return cubes
}
