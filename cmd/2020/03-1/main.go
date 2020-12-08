package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m, width, height := InputToMap(2020, 3)

	p := aoc.Point2D{}
	dx := 3
	dy := 1

	var count int
	for p.Y <= height {
		if m[p] {
			count++
		}

		p = aoc.Point2D{X: (p.X + dx) % (width + 1), Y: p.Y + dy}
	}

	fmt.Println(count)
}

func Dump(m map[aoc.Point2D]bool, width, height int, current aoc.Point2D) {
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			p := aoc.Point2D{X: x, Y: y}
			if p == current {
				fmt.Print("*")
				continue
			}
			if m[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func InputToMap(year, day int) (map[aoc.Point2D]bool, int, int) {
	m := make(map[aoc.Point2D]bool)
	var width, height int
	for y, line := range aoc.InputToLines(year, day) {
		height = y
		for x, c := range line {
			m[aoc.Point2D{X: x, Y: y}] = c == '#'
			width = x
		}
	}
	return m, width, height
}
