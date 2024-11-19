package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	rows := []Tiles{InputToTiles()}
	for row := 0; len(rows) < 40; row++ {
		rows = append(rows, Next(rows[row]))
	}

	var count int
	for row := 0; row < len(rows); row++ {
		for _, tile := range rows[row] {
			if tile == Safe {
				count++
			}
		}
	}
	fmt.Println(count)
}

func Next(tiles Tiles) Tiles {
	traps := map[[3]bool]bool{
		{Trap, Trap, Safe}: Trap,
		{Safe, Trap, Trap}: Trap,
		{Trap, Safe, Safe}: Trap,
		{Safe, Safe, Trap}: Trap,
	}

	var next Tiles
	for col := 0; col < len(tiles); col++ {
		var L, C, R bool
		if col > 0 {
			L = tiles[col-1]
		}
		C = tiles[col]
		if col < len(tiles)-1 {
			R = tiles[col+1]
		}

		next = append(next, traps[[3]bool{L, C, R}])
	}
	return next
}

const (
	Safe = false
	Trap = true
)

type Tiles []bool

func InputToTiles() Tiles {
	var tiles Tiles
	for _, c := range lib.InputToString() {
		tiles = append(tiles, c == '^')
	}

	return tiles
}
