package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	grid := lib.InputToStringGrid2D()

	gears := make(map[lib.Point2D][]int)
	ForEachNumber(grid, func(x int, y int, n int) {
		digits := lib.Digits(n)

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= len(digits); dx++ {
				p := lib.Point2D{X: x + dx, Y: y + dy}

				if grid.InBoundsPoint(p) && grid.GetPoint(p) == "*" {
					gears[p] = append(gears[p], n)
				}
			}
		}
	})

	var prod int
	for _, nums := range gears {
		if len(nums) == 2 {
			prod += lib.Product(nums...)
		}
	}
	fmt.Println(prod)
}

func ForEachNumber(g lib.Grid2D[string], fn func(int, int, int)) {
	g.ForEach(func(x0 int, y int, s string) {
		// Check if this is the beginning of a number
		if IsDigit(s) && (x0 == 0 || !IsDigit(g.Get(x0-1, y))) {
			var digits strings.Builder
			for x := x0; x < g.Width && IsDigit(g.Get(x, y)); x++ {
				digits.WriteString(g.Get(x, y))
			}

			fn(x0, y, lib.ParseInt(digits.String()))
		}
	})
}

func IsDigit(s string) bool {
	return strings.Contains("0123456789", s)
}
