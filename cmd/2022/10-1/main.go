package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var Track = aoc.SetFrom(20, 60, 100, 140, 180, 220)

func main() {
	var sum int
	cycle, x := 1, 1

	update := func() {
		cycle++
		if Track.Contains(cycle) {
			sum += cycle * x
		}
	}

	for _, line := range aoc.InputToLines(2022, 10) {
		op, arg, _ := strings.Cut(line, " ")
		switch op {
		case "addx":
			update()
			x += aoc.ParseInt(arg)
			update()
		case "noop":
			update()
		}
	}

	fmt.Println(sum)
}
