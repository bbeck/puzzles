package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	matrix := InputToMatrix(2017, 2)

	var sum int
	for _, row := range matrix {
		min, max := minmax(row)
		sum += max - min
	}

	fmt.Printf("checksum: %d\n", sum)
}

func minmax(row []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64

	for _, n := range row {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	return min, max
}

func InputToMatrix(year, day int) [][]int {
	var matrix [][]int
	for _, line := range aoc.InputToLines(2017, 2) {
		var row []int
		for _, s := range strings.Split(line, "\t") {
			row = append(row, aoc.ParseInt(s))
		}

		matrix = append(matrix, row)
	}

	return matrix
}
