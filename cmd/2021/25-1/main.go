package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	width, height, m := InputToMap()

	var step int
	for step = 1; ; step++ {
		var moved bool

		// right facing moves first
		for y := 0; y < height; y++ {
			// Keep track of which moves to make for this row.
			moves := make(map[int]int)

			for x := 0; x < width; x++ {
				if m[y][x] != ">" {
					continue
				}

				nx := (x + 1) % width
				if m[y][nx] != "." {
					continue
				}

				moves[x] = nx
			}

			// Make the moves
			for x, nx := range moves {
				m[y][x] = "."
				m[y][nx] = ">"
				moved = true
			}
		}

		// down facing move next
		for x := 0; x < width; x++ {
			// Keep track of which moves to make for this column.
			moves := make(map[int]int)

			for y := 0; y < height; y++ {
				if m[y][x] != "v" {
					continue
				}

				ny := (y + 1) % height
				if m[ny][x] != "." {
					continue
				}

				moves[y] = ny
			}

			// Make the moves
			for y, ny := range moves {
				m[y][x] = "."
				m[ny][x] = "v"
				moved = true
			}
		}

		if !moved {
			break
		}
	}

	fmt.Println(step)
}

func InputToMap() (int, int, [][]string) {
	var m [][]string
	for y, line := range aoc.InputToLines(2021, 25) {
		m = append(m, make([]string, len(line)))

		for x, c := range line {
			m[y][x] = string(c)
		}
	}

	return len(m[0]), len(m), m
}
