package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	grid := InputToGrid()
	grid.ForEach(func(x int, y int, s string) {
		if s != "." {
			return
		}

		row := Row(grid, 8*(x/8), y)
		col := Col(grid, x, 8*(y/8))
		ch := SetFrom(row...).IntersectElems(col...).DifferenceElems(".")
		grid.Set(x, y, ch.Entries()[0])
	})

	var sum int
	for y := 0; y < grid.Height/8; y++ {
		for x := 0; x < grid.Width/8; x++ {
			word := Word(grid, 8*x, 8*y)
			sum += Power(word)
		}
	}
	fmt.Println(sum)
}

func InputToGrid() Grid2D[string] {
	var lines []string
	for _, line := range InputToLines() {
		if line == "" {
			continue
		}
		lines = append(lines, strings.ReplaceAll(line, " ", ""))
	}

	grid := NewGrid2D[string](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, ch := range line {
			grid.Set(x, y, string(ch))
		}
	}
	return grid
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

func Power(word []string) int {
	powers := map[string]int{
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
		"J": 10, "K": 11, "L": 12, "M": 13, "N": 14, "O": 15, "P": 16, "Q": 17,
		"R": 18, "S": 19, "T": 20, "U": 21, "V": 22, "W": 23, "X": 24, "Y": 25,
		"Z": 26,
	}

	var power int
	for i, s := range word {
		power += (i + 1) * powers[s]
	}
	return power
}
