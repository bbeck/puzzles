package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	cube := InputToCube(2020, 17)
	for i := 0; i < 6; i++ {
		cube = Next(cube)
	}

	var sum int
	for _, v := range cube {
		if v {
			sum++
		}
	}
	fmt.Println(sum)
}

func Next(cube map[Point4D]bool) map[Point4D]bool {
	minX, maxX, minY, maxY, minZ, maxZ, minW, maxW := Bounds(cube)

	next := make(map[Point4D]bool)
	for x := minX - 1; x <= maxX+1; x++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for z := minZ - 1; z <= maxZ+1; z++ {
				for w := minW - 1; w <= maxW+1; w++ {
					p := Point4D{X: x, Y: y, Z: z, W: w}

					var sum int
					for _, n := range p.Neighbors() {
						if cube[n] {
							sum++
						}
					}

					// Only write trues into the cube.
					if (cube[p] && sum == 2) || sum == 3 {
						next[p] = true
					}
				}
			}
		}
	}

	return next
}

func Bounds(cube map[Point4D]bool) (int, int, int, int, int, int, int, int) {
	var minX, maxX, minY, maxY, minZ, maxZ, minW, maxW int
	for p := range cube {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
		if p.Z < minZ {
			minZ = p.Z
		}
		if p.Z > maxZ {
			maxZ = p.Z
		}
		if p.W < minW {
			minW = p.W
		}
		if p.W > maxW {
			maxW = p.W
		}
	}

	return minX, maxX, minY, maxY, minZ, maxZ, minW, maxW
}

type Point4D struct {
	X, Y, Z, W int
}

func (p Point4D) Neighbors() []Point4D {
	var neighbors []Point4D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}

					p := Point4D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz, W: p.W + dw}
					neighbors = append(neighbors, p)
				}
			}
		}
	}

	return neighbors
}

func InputToCube(year, day int) map[Point4D]bool {
	cube := make(map[Point4D]bool)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := Point4D{X: x, Y: y, Z: 0, W: 0}
			cube[p] = c == '#'
		}
	}

	return cube
}
