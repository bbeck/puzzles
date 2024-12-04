package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	grid := InputToStringGrid2D()

	get := func(p Point2D, dx, dy int) string {
		var sb strings.Builder
		for n := 0; n < 4; n++ {
			q := Point2D{X: p.X + n*dx, Y: p.Y + n*dy}
			if grid.InBoundsPoint(q) {
				sb.WriteString(grid.GetPoint(q))
			}
		}
		return sb.String()
	}

	var count int
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "X" {
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					if get(p, dx, dy) == "XMAS" {
						count++
					}
				}
			}
		}
	})
	fmt.Println(count)
}
