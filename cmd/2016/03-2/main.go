package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, triangle := range InputToTriangles(2016, 3) {
		if triangle.s1+triangle.s2 <= triangle.s3 {
			continue
		}
		if triangle.s1+triangle.s3 <= triangle.s2 {
			continue
		}
		if triangle.s2+triangle.s3 <= triangle.s1 {
			continue
		}

		count++
	}

	fmt.Printf("possible: %d\n", count)
}

type Triangle struct {
	s1, s2, s3 int
}

func InputToTriangles(year, day int) []Triangle {
	var rows [][]int
	for _, line := range aoc.InputToLines(year, day) {
		line = strings.Trim(line, " ")
		for i := 0; i < 10; i++ {
			line = strings.ReplaceAll(line, "  ", " ")
		}

		parts := strings.Split(line, " ")
		rows = append(rows, []int{
			aoc.ParseInt(parts[0]),
			aoc.ParseInt(parts[1]),
			aoc.ParseInt(parts[2]),
		})
	}

	var triangles []Triangle
	for i := 0; i < len(rows); i += 3 {
		for col := 0; col < 3; col++ {
			triangles = append(triangles, Triangle{
				rows[i][col],
				rows[i+1][col],
				rows[i+2][col],
			})
		}
	}

	return triangles
}
