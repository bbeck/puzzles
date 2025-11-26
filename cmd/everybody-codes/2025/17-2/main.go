package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var vx, vy int
	grid := in.ToGrid2D(func(x, y int, s string) int {
		if s == "@" {
			vx, vy = x, y
			return 0
		}
		return ParseInt(s)
	})

	var largest, largestR int
	for r := range Min(grid.Width, grid.Height) {
		var sum int
		grid.ForEach(func(x int, y int, n int) {
			dx, dy := x-vx, y-vy
			if dx*dx+dy*dy <= r*r {
				sum += n
				grid.Set(x, y, 0)
			}
		})

		if sum > largest {
			largest = sum
			largestR = r
		}
	}
	fmt.Println(largest * largestR)
}
