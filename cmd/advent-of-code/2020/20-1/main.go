package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
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
	puz.Grid2D[bool]
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
	s := Tile{ID: t.ID, Grid2D: puz.NewGrid2D[bool](N, N)}
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			s.Set(x, y, t.Get(N-x-1, y))
		}
	}
	return s
}

func (t Tile) Rotate() Tile {
	return Tile{ID: t.ID, Grid2D: t.RotateRight()}
}

func (t Tile) FitsOnTop(s Tile) bool {
	N := t.Width
	for n := 0; n < N; n++ {
		if t.Get(n, 0) != s.Get(n, N-1) {
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
		if t.Get(N-1, n) != s.Get(0, n) {
			return false
		}
	}
	return true
}

func (t Tile) FitsOnLeft(s Tile) bool {
	return s.FitsOnRight(t)
}

func InputToTiles() []Tile {
	lines := puz.InputToLines()
	N := len(lines[1])

	var tiles []Tile
	for base := 0; base < len(lines); base += 12 {
		var id int
		fmt.Sscanf(lines[base], "Tile %d:", &id)

		grid := puz.NewGrid2D[bool](N, N)
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				grid.Set(x, y, lines[base+y+1][x] == '#')
			}
		}

		tiles = append(tiles, Tile{ID: id, Grid2D: grid})
	}

	return tiles
}
