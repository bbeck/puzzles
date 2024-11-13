package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	area := Area{puz.InputToStringGrid2D(2018, 18)}
	for n := 0; n < 10; n++ {
		area = Next(area)
	}

	counts := make(map[string]int)
	area.ForEach(func(_, _ int, value string) {
		counts[value]++
	})

	fmt.Println(counts["|"] * counts["#"])
}

func Next(area Area) Area {
	next := Area{puz.NewGrid2D[string](area.Width, area.Height)}
	area.ForEach(func(x, y int, value string) {
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
		next.Set(x, y, value)
	})

	return next
}

type Area struct{ puz.Grid2D[string] }
