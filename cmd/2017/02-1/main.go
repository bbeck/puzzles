package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var sum int
	for _, row := range InputToRows() {
		sum += aoc.Max(row...) - aoc.Min(row...)
	}
	fmt.Println(sum)
}

func InputToRows() [][]int {
	return aoc.InputLinesTo(2017, 2, func(line string) ([]int, error) {
		var row []int
		for _, s := range strings.Fields(line) {
			row = append(row, aoc.ParseInt(s))
		}

		return row, nil
	})
}
