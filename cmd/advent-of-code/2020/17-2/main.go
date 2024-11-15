package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	cube := InputToCube()
	for i := 0; i < 6; i++ {
		cube = Next(cube)
	}

	fmt.Println(len(cube))
}

func Next(cube puz.Set[Point4D]) puz.Set[Point4D] {
	var next puz.Set[Point4D]

	minW, minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
	maxW, maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt, math.MinInt
	for p := range cube {
		minW, maxW = puz.Min(minW, p.W), puz.Max(maxW, p.W)
		minX, maxX = puz.Min(minX, p.X), puz.Max(maxX, p.X)
		minY, maxY = puz.Min(minY, p.Y), puz.Max(maxY, p.Y)
		minZ, maxZ = puz.Min(minZ, p.Z), puz.Max(maxZ, p.Z)
	}

	for w := minW - 1; w <= maxW+1; w++ {
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					p := Point4D{W: w, X: x, Y: y, Z: z}

					var active int
					for _, n := range p.Neighbors() {
						if cube.Contains(n) {
							active++
						}
					}

					if cube.Contains(p) && (active == 2 || active == 3) {
						next.Add(p)
					} else if !cube.Contains(p) && (active == 3) {
						next.Add(p)
					}
				}
			}
		}
	}

	return next
}

func InputToCube() puz.Set[Point4D] {
	var cube puz.Set[Point4D]
	for y, line := range puz.InputToLines(2020, 17) {
		for x, c := range line {
			if c == '#' {
				cube.Add(Point4D{X: x, Y: y})
			}
		}
	}

	return cube
}

type Point4D struct {
	W, X, Y, Z int
}

func (p Point4D) Neighbors() []Point4D {
	var neighbors []Point4D
	for dw := -1; dw <= 1; dw++ {
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				for dz := -1; dz <= 1; dz++ {
					if dw == 0 && dx == 0 && dy == 0 && dz == 0 {
						continue
					}
					neighbors = append(neighbors, Point4D{W: p.W + dw, X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz})
				}
			}
		}
	}

	return neighbors
}
