package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ps := lib.SetFrom(InputToPoints()...)

	var count int
	for p := range ps {
		for _, n := range p.OrthogonalNeighbors() {
			if !ps.Contains(n) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func InputToPoints() []lib.Point3D {
	return lib.InputLinesTo(func(line string) lib.Point3D {
		var p lib.Point3D
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		return p
	})
}
