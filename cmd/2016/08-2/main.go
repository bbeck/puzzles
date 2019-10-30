package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var screen Screen
	for _, line := range aoc.InputToLines(2016, 8) {
		var width, height int
		var x, y, distance int
		if _, err := fmt.Sscanf(line, "rect %dx%d", &width, &height); err == nil {
			screen = screen.Rect(width, height)
		}

		if _, err := fmt.Sscanf(line, "rotate row y=%d by %d", &y, &distance); err == nil {
			screen = screen.RotateRow(y, distance)
		}

		if _, err := fmt.Sscanf(line, "rotate column x=%d by %d", &x, &distance); err == nil {
			screen = screen.RotateCol(x, distance)
		}
	}

	screen.Display()
}

type Screen [6][50]bool

func (s Screen) Rect(width, height int) Screen {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			s[y][x] = true
		}
	}

	return s
}

func (s Screen) RotateRow(y int, distance int) Screen {
	for d := 0; d < distance; d++ {
		last := s[y][len(s[y])-1]
		for x := len(s[y]) - 1; x >= 1; x-- {
			s[y][x] = s[y][x-1]
		}
		s[y][0] = last
	}

	return s
}

func (s Screen) RotateCol(x int, distance int) Screen {
	for d := 0; d < distance; d++ {
		last := s[len(s)-1][x]
		for y := len(s) - 1; y >= 1; y-- {
			s[y][x] = s[y-1][x]
		}
		s[0][x] = last
	}

	return s
}

func (s Screen) Display() {
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			if s[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
