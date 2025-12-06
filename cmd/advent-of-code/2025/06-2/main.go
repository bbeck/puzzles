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
	var group []int
	for x := 0; x < grid.Width; x++ {
		var digits []int
		for y := 0; y < grid.Height-1; y++ {
			if v := grid.Get(x, y); strings.TrimSpace(v) != "" {
				digits = append(digits, ParseInt(v))
			}
		}

		if len(digits) > 0 {
			group = append(group, JoinDigits(digits))
		}

		if len(digits) == 0 || x == grid.Width-1 {
			numbers = append(numbers, group)
			group = nil
		}
	}

	var ops []string
	for x := 0; x < grid.Width; x++ {
		if op := grid.Get(x, grid.Height-1); strings.TrimSpace(op) != "" {
			ops = append(ops, op)
		}
	}

	var sum int
	for i, op := range ops {
		if op == "+" {
			sum += Sum(numbers[i]...)
		}
		if op == "*" {
			sum += Product(numbers[i]...)
		}
	}
	fmt.Println(sum)
}
