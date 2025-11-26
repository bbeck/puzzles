package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

const R = 10

func main() {
	var vx, vy int
	grid := in.ToGrid2D(func(x, y int, s string) int {
		if s == "@" {
			vx, vy = x, y
			return 0
		}
		return ParseInt(s)
	})

	var sum int
	grid.ForEach(func(x, y, n int) {
		dx, dy := x-vx, y-vy
		if dx*dx+dy*dy <= R*R {
			sum += n
		}
	})

	fmt.Println(sum)
}
