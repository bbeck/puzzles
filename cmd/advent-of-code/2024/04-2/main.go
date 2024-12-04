package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

var Goals = SetFrom("MAS", "SAM")

func main() {
	grid := InputToStringGrid2D()

	get := func(p Point2D, dx, dy int) string {
		var sb strings.Builder
		for n := -1; n <= 1; n++ {
			q := Point2D{X: p.X + n*dx, Y: p.Y + n*dy}
			if grid.InBoundsPoint(q) {
				sb.WriteString(grid.GetPoint(q))
			}
		}
		return sb.String()
	}

	var count int
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "A" {
			if Goals.Contains(get(p, 1, 1)) && Goals.Contains(get(p, 1, -1)) {
				count++
			}
		}
	})
	fmt.Println(count)
}
