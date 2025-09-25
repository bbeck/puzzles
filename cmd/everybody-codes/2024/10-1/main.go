package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToGrid2D(func(_, _ int, s string) string {
		return s
	})

	grid.ForEach(func(x int, y int, s string) {
		if s != "." {
			return
		}

		row := Row(grid, 0, y)
		col := Col(grid, x, 0)
		ch := SetFrom(row...).IntersectElems(col...).DifferenceElems(".")
		grid.Set(x, y, ch.Entries()[0])
	})

	word := Word(grid, 0, 0)
	fmt.Println(strings.Join(word, ""))
}

func Row(grid Grid2D[string], x, y int) []string {
	var row []string
	for dx := 0; dx < 8; dx++ {
		row = append(row, grid.Get(x+dx, y))
	}
	return row
}

func Col(grid Grid2D[string], x, y int) []string {
	var col []string
	for dy := 0; dy < 8; dy++ {
		col = append(col, grid.Get(x, y+dy))
	}
	return col
}

func Word(grid Grid2D[string], x, y int) []string {
	var word []string
	for dy := 2; dy < 6; dy++ {
		for dx := 2; dx < 6; dx++ {
			word = append(word, grid.Get(x+dx, y+dy))
		}
	}
	return word
}
