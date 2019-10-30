package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	matrix := InputToMatrix(2017, 2)

	var sum int
	for _, row := range matrix {
		a, b := divides(row)
		sum += a / b
	}

	fmt.Printf("checksum: %d\n", sum)
}

func divides(row []int) (int, int) {
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(row); j++ {
			if i == j {
				continue
			}

			if row[i]%row[j] == 0 {
				return row[i], row[j]
			}
		}
	}

	log.Fatal("unable to find the two numbers that divide each other")
	return 0, 0
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
