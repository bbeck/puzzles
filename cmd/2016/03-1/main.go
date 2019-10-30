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
	var triangles []Triangle
	for _, line := range aoc.InputToLines(year, day) {
		line = strings.Trim(line, " ")
		for i := 0; i < 10; i++ {
			line = strings.ReplaceAll(line, "  ", " ")
		}

		parts := strings.Split(line, " ")
		triangles = append(triangles, Triangle{
			aoc.ParseInt(parts[0]),
			aoc.ParseInt(parts[1]),
			aoc.ParseInt(parts[2]),
		})
	}

	return triangles
}
