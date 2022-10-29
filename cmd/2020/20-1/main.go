package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	tiles := InputToTiles()

	product := 1
	for _, t := range tiles {
		if IsCorner(t, tiles) {
			product *= t.ID
		}
	}
	fmt.Println(product)
}

func IsCorner(t Tile, tiles []Tile) bool {
	// A tile is a corner piece if only 2 other tiles are compatible with it
	var count int
	for _, o := range tiles {
		if t.ID == o.ID {
			continue
		}

		for _, s := range o.Orientations() {
			if t.FitsOnTop(s) || t.FitsOnRight(s) || t.FitsOnBottom(s) || t.FitsOnLeft(s) {
				count++
			}
		}
	}

	return count == 2
}

type Tile struct {
	ID int
	aoc.Grid2D[bool]
}

func (t Tile) Orientations() []Tile {
	A := t.Rotate()
	B := A.Rotate()
	C := B.Rotate()
	D := t.Flip()
	E := D.Rotate()
	F := E.Rotate()
	G := F.Rotate()
	return []Tile{t, A, B, C, D, E, F, G}
}

func (t Tile) Flip() Tile {
	N := t.Width
	s := Tile{ID: t.ID, Grid2D: aoc.NewGrid2D[bool](N, N)}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			s.AddXY(x, y, t.GetXY(N-x-1, y))
		}
	}
	return s
}

func (t Tile) Rotate() Tile {
	N := t.Width
	s := Tile{ID: t.ID, Grid2D: aoc.NewGrid2D[bool](N, N)}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			s.AddXY(x, y, t.GetXY(N-y-1, x))
		}
	}
	return s
}

func (t Tile) FitsOnTop(s Tile) bool {
	N := t.Width
	for n := 0; n < N; n++ {
		if t.GetXY(n, 0) != s.GetXY(n, N-1) {
			return false
		}
	}
	return true
}

func (t Tile) FitsOnBottom(s Tile) bool {
	return s.FitsOnTop(t)
}

func (t Tile) FitsOnRight(s Tile) bool {
	N := t.Width
	for n := 0; n < N; n++ {
		if t.GetXY(N-1, n) != s.GetXY(0, n) {
			return false
		}
	}
	return true
}

func (t Tile) FitsOnLeft(s Tile) bool {
	return s.FitsOnRight(t)
}

func InputToTiles() []Tile {
	lines := aoc.InputToLines(2020, 20)
	N := len(lines[1])

	var tiles []Tile
	for base := 0; base < len(lines); base += 12 {
		var id int
		fmt.Sscanf(lines[base], "Tile %d:", &id)

		grid := aoc.NewGrid2D[bool](N, N)
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				grid.AddXY(x, y, lines[base+y+1][x] == '#')
			}
		}

		tiles = append(tiles, Tile{ID: id, Grid2D: grid})
	}

	return tiles
}
