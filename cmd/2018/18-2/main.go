package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	area := InputToArea()

	// Keep track of each area we've seen along with the time we saw it at.  If
	// we repeat then we know there was a cycle.
	var previous aoc.Deque[Area]
	seen := make(map[string]int) // mapping of area key to the time we last saw it

	var tm int
	var key string
	for tm = 1; ; tm++ {
		area = Next(area)
		key = Key(area)
		if _, found := seen[key]; found {
			break
		}

		previous.PushBack(area)
		seen[key] = tm
	}

	// Move the entry we care about to the front of the deque.
	cycle := tm - seen[key]
	remainder := (1_000_000_000 - tm) % cycle
	previous.Rotate(cycle - remainder)

	counts := make(map[string]int)
	previous.PeekFront().ForEach(func(p aoc.Point2D, value string) {
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

func Key(area Area) string {
	var sb strings.Builder
	area.ForEach(func(p aoc.Point2D, value string) { sb.WriteString(value) })
	return sb.String()
}

type Area struct{ aoc.Grid2D[string] }

func InputToArea() Area {
	grid := aoc.NewGrid2D[string](50, 50)
	for y, line := range aoc.InputToLines(2018, 18) {
		for x, c := range line {
			grid.AddXY(x, y, string(c))
		}
	}

	return Area{grid}
}
