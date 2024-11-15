package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	ps := puz.SetFrom(InputToPoints()...)

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

func InputToPoints() []puz.Point3D {
	return puz.InputLinesTo(func(line string) puz.Point3D {
		var p puz.Point3D
		fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		return p
	})
}
