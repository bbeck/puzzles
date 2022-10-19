package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	area := InputToArea()
	for n := 0; n < 10; n++ {
		area = Next(area)
	}

	counts := make(map[string]int)
	area.ForEach(func(p aoc.Point2D, value string) {
		counts[value]++
	})

	fmt.Println(counts["|"] * counts["#"])
}

func Next(area Area) Area {
	next := Area{aoc.NewGrid2D[string](area.Width, area.Height)}
	area.ForEach(func(p aoc.Point2D, value string) {
		counts := make(map[string]int)
		for _, n := range p.Neighbors() {
			if area.InBounds(n) {
				counts[area.Get(n)]++
			}
		}

		if value == "." && counts["|"] >= 3 {
			value = "|"
		} else if value == "|" && counts["#"] >= 3 {
			value = "#"
		} else if value == "#" && (counts["#"] < 1 || counts["|"] < 1) {
			value = "."
		}
		next.Add(p, value)
	})

	return next
}

type Area struct{ aoc.Grid2D[string] }

func InputToArea() Area {
	return Area{aoc.InputToGrid2D(2018, 18, func(x int, y int, s string) string {
		return s
	})}
}
