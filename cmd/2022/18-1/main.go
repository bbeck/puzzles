package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ps := aoc.SetFrom(InputToPoints()...)

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

func InputToPoints() []aoc.Point3D {
	return aoc.InputLinesTo(2022, 18, func(line string) (aoc.Point3D, error) {
		var p aoc.Point3D
		_, err := fmt.Sscanf(line, "%d,%d,%d", &p.X, &p.Y, &p.Z)
		return p, err
	})
}
