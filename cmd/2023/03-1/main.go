package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	grid := puz.InputToStringGrid2D(2023, 3)

	var sum int
	ForEachNumber(grid, func(x int, y int, num int) {
		for n := num; n > 0; n = n / 10 {
			grid.ForEachNeighbor(x, y, func(_ int, _ int, s string) {
				if IsSymbol(s) {
					sum += num
					n = 0
				}
			})

			x++
		}
	})

	fmt.Println(sum)
}

func ForEachNumber(g puz.Grid2D[string], fn func(int, int, int)) {
	g.ForEach(func(x0 int, y int, s string) {
		// Check if this is the beginning of a number
		if IsDigit(s) && (x0 == 0 || !IsDigit(g.Get(x0-1, y))) {
			var digits strings.Builder
			for x := x0; x < g.Width && IsDigit(g.Get(x, y)); x++ {
				digits.WriteString(g.Get(x, y))
			}

			fn(x0, y, puz.ParseInt(digits.String()))
		}
	})
}

func IsDigit(s string) bool {
	return strings.Contains("0123456789", s)
}

func IsSymbol(s string) bool {
	return s != "." && !IsDigit(s)
}
