package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var numbers [][]int
	for y := 0; y < grid.Height-1; y++ {
		var row, digits []int
		for x := 0; x < grid.Width; x++ {
			if v := grid.Get(x, y); strings.TrimSpace(v) != "" {
				digits = append(digits, ParseInt(v))
				if x < grid.Width-1 {
					continue
				}
			}

			if len(digits) > 0 {
				row = append(row, JoinDigits(digits))
				digits = nil
			}
		}

		numbers = append(numbers, row)
	}

	var ops []string
	for x := 0; x < grid.Width; x++ {
		if op := grid.Get(x, grid.Height-1); strings.TrimSpace(op) != "" {
			ops = append(ops, op)
		}
	}

	var sum int
	for i, op := range ops {
		var col []int
		for j := range numbers {
			col = append(col, numbers[j][i])
		}

		if op == "+" {
			sum += Sum(col...)
		}
		if op == "*" {
			sum += Product(col...)
		}
	}
	fmt.Println(sum)
}
