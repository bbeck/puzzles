package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	crt := aoc.NewGrid2D[bool](40, 6)
	cycle, x := 0, 1

	draw := func() {
		value := x-1 <= cycle%40 && cycle%40 <= x+1
		crt.Add(cycle%40, cycle/40, value)
		cycle++
	}

	for _, line := range aoc.InputToLines(2022, 10) {
		op, arg, _ := strings.Cut(line, " ")
		switch op {
		case "addx":
			draw()
			draw()
			x += aoc.ParseInt(arg)
		case "noop":
			draw()
		}
	}

	for y := 0; y < crt.Height; y++ {
		for x := 0; x < crt.Width; x++ {
			if crt.Get(x, y) {
				fmt.Print("█")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
