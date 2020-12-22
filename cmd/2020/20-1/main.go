package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var N int

func main() {
	tiles := InputToTiles(2020, 20)
	N = int(math.Sqrt(float64(len(tiles[0].data))))

	product := 1
	for _, tile := range tiles {
		if IsCorner(tile, tiles) {
			product *= tile.id
		}
	}
	fmt.Println(product)
}

func IsCorner(tile Tile, tiles []Tile) bool {
	// A tile is a corner only if we can only find 2 other tiles to fit adjacent
	// to it.
	neighbors := aoc.NewSet()
	for _, t := range tiles {
		if tile.id == t.id {
			continue
		}

		for _, other := range t.Transformed() {
			if FitsHorizontally(tile, other) || FitsHorizontally(other, tile) ||
				FitsVertically(tile, other) || FitsVertically(other, tile) {
				neighbors.Add(other.id)
			}
		}
	}

	return neighbors.Size() == 2
}

func FitsHorizontally(l, r Tile) bool {
	for y := 0; y < N; y++ {
		if l.data[y*N+N-1] != r.data[y*N] {
			return false
		}
	}
	return true
}

func FitsVertically(t, b Tile) bool {
	for x := 0; x < N; x++ {
		if t.data[(N-1)*N+x] != b.data[x] {
			return false
		}
	}
	return true
}

type Tile struct {
	id   int
	data []bool
}

func (t Tile) RotateRight() Tile {
	transformed := make([]bool, len(t.data))
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			transformed[x*N+N-y-1] = t.data[y*N+x]
		}
	}

	return Tile{
		id:   t.id,
		data: transformed,
	}
}

func (t Tile) FlipH() Tile {
	transformed := make([]bool, len(t.data))
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			transformed[y*N+x] = t.data[y*N+N-x-1]
		}
	}

	return Tile{
		id:   t.id,
		data: transformed,
	}
}

func (t Tile) FlipV() Tile {
	transformed := make([]bool, len(t.data))
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			transformed[y*N+x] = t.data[(N-y-1)*N+x]
		}
	}

	return Tile{
		id:   t.id,
		data: transformed,
	}
}

func (t Tile) Transformed() []Tile {
	A := t
	B := A.RotateRight()
	C := B.RotateRight()
	D := C.RotateRight()
	E := A.FlipH()
	F := E.RotateRight()
	G := F.RotateRight()
	H := G.RotateRight()
	I := A.FlipV()
	J := I.RotateRight() // Probably not needed from here on
	K := J.RotateRight()
	L := K.RotateRight()
	return []Tile{A, B, C, D, E, F, G, H, I, J, K, L}
}

func (t Tile) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Tile: %d\n", t.id))
	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if t.data[y*N+x] {
				sb.WriteRune('#')
			} else {
				sb.WriteRune('.')
			}
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func InputToTiles(year, day int) []Tile {
	blocks := make(map[int][]string)

	var current int
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			continue
		}

		var id int
		if _, err := fmt.Sscanf(line, "Tile %d:", &id); err == nil {
			current = id
			continue
		}

		blocks[current] = append(blocks[current], line)
	}

	var tiles []Tile
	for id, lines := range blocks {
		var data []bool
		for _, line := range lines {
			for _, c := range line {
				data = append(data, c == '#')
			}
		}

		tiles = append(tiles, Tile{
			id:   id,
			data: data,
		})
	}

	return tiles
}
