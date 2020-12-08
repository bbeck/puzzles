package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m, width, height := InputToMap(2020, 3)

	product := CountTrees(m, width, height, 1, 1) *
		CountTrees(m, width, height, 3, 1) *
		CountTrees(m, width, height, 5, 1) *
		CountTrees(m, width, height, 7, 1) *
		CountTrees(m, width, height, 1, 2)

	fmt.Println(product)
}

func CountTrees(m map[aoc.Point2D]int, width, height int, dx, dy int) int {
	next := func(p aoc.Point2D) aoc.Point2D {
		return aoc.Point2D{X: (p.X + dx) % (width + 1), Y: p.Y + dy}
	}

	origin := func() aoc.Point2D {
		return aoc.Point2D{}
	}

	var count int
	for p := origin(); p.Y <= height; p = next(p) {
		count += m[p]
	}

	return count
}

func InputToMap(year, day int) (map[aoc.Point2D]int, int, int) {
	m := make(map[aoc.Point2D]int)
	var width, height int
	for y, line := range aoc.InputToLines(year, day) {
		height = y
		for x, c := range line {
			if c == '#' {
				m[aoc.Point2D{X: x, Y: y}] = 1
			}
			width = x
		}
	}
	return m, width, height
}
