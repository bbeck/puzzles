package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	grid := puz.InputToIntGrid2D()

	var count int
	grid.ForEachPoint(func(p puz.Point2D, _ int) {
		visible := []bool{
			IsVisible(grid, p, puz.Left),
			IsVisible(grid, p, puz.Right),
			IsVisible(grid, p, puz.Up),
			IsVisible(grid, p, puz.Down),
		}
		if puz.Any(visible, puz.Identity[bool]) {
			count++
		}
	})
	fmt.Println(count)
}

func IsVisible(g puz.Grid2D[int], p puz.Point2D, h puz.Heading) bool {
	height := g.GetPoint(p)
	for p = p.Move(h); g.InBoundsPoint(p); p = p.Move(h) {
		if g.GetPoint(p) >= height {
			return false
		}
	}
	return true
}
