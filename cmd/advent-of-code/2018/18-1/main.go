package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	area := in.ToGrid2D(func(x, y int, s string) string { return s })
	for n := 0; n < 10; n++ {
		area = Next(area)
	}

	counts := make(map[string]int)
	area.ForEach(func(_, _ int, value string) {
		counts[value]++
	})

	fmt.Println(counts["|"] * counts["#"])
}

func Next(area Grid2D[string]) Grid2D[string] {
	return area.Map(func(x int, y int, value string) string {
		counts := make(map[string]int)
		area.ForEachNeighbor(x, y, func(_, _ int, value string) {
			counts[value]++
		})

		if value == "." && counts["|"] >= 3 {
			value = "|"
		} else if value == "|" && counts["#"] >= 3 {
			value = "#"
		} else if value == "#" && (counts["#"] < 1 || counts["|"] < 1) {
			value = "."
		}

		return value
	})
}
