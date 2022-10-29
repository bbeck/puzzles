package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
)

func main() {
	tiles := InputToTiles()
	image := AssembleImage(tiles)

	var in aoc.Set[aoc.Point2D] // The points in the image in a monster

	for _, m := range GetSeaMonster().Orientations() {
		for y := 0; y < image.Height-m.Height; y++ {
		next:
			for x := 0; x < image.Width-m.Width; x++ {
				var ps []aoc.Point2D // Points in the image for this monster

				for my := 0; my < m.Height; my++ {
					for mx := 0; mx < m.Width; mx++ {
						if m.GetXY(mx, my) {
							p := aoc.Point2D{X: x + mx, Y: y + my}
							ps = append(ps, p)

							if !image.Get(p) {
								continue next
							}
						}
					}
				}

				in.Add(ps...)
			}
		}
	}

	var count int
	image.ForEach(func(p aoc.Point2D, value bool) {
		if value && !in.Contains(p) {
			count++
		}
	})

	fmt.Println(count)
}

func AssembleImage(ts []Tile) Tile {
	var corner Tile
	for _, t := range ts {
		if IsCorner(t, ts) {
			corner = t
			break
		}
	}

	// The resulting image will be DxD tiles.
	D := int(math.Sqrt(float64(len(ts))))
	tiles := aoc.Make2D[Tile](D, D)

	// Helper to determine if a tile fits in a specific location in the grid.  If
	// the tile fits, then a copy of the tile in the proper orientation is
	// returned.
	fits := func(x, y int, t Tile) (Tile, bool) {
		for _, o := range t.Orientations() {
			if x > 0 && !tiles[y][x-1].FitsOnRight(o) {
				continue
			}
			if y > 0 && !tiles[y-1][x].FitsOnBottom(o) {
				continue
			}
			return o, true
		}

		return t, false
	}

	// The corner we found might not be in the right orientation to be the top
	// left corner of the image, so we'll have to try all possible orientations.
outer:
	for _, tl := range corner.Orientations() {
		used := aoc.SingletonSet(tl.ID)
		tiles[0][0] = tl

		for y := 0; y < D; y++ {
			for x := 0; x < D; x++ {
				if x == 0 && y == 0 {
					continue
				}

				var found bool
				for _, t := range ts {
					if used.Contains(t.ID) {
						continue
					}

					if s, ok := fits(x, y, t); ok {
						tiles[y][x] = s
						used.Add(s.ID)
						found = true
						break
					}
				}

				if !found {
					// We weren't able to find a tile for this location, the orientation
					// of the corner piece must be at fault.  Try another orientation.
					continue outer
				}
			}
		}

		break
	}

	// Now that we have the arrangement of tiles, piece them together into one
	// big tile ignoring the borders of each tile.
	N := corner.Width
	I := D * (N - 2)
	image := aoc.NewGrid2D[bool](I, I)
	for ty := 0; ty < D; ty++ {
		for tx := 0; tx < D; tx++ {
			tile := tiles[ty][tx]

			for y := 1; y < N-1; y++ {
				for x := 1; x < N-1; x++ {
					image.AddXY(tx*(N-2)+x-1, ty*(N-2)+y-1, tile.GetXY(x, y))
				}
			}
		}
	}
	return Tile{Grid2D: image}
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

func GetSeaMonster() Tile {
	lines := []string{
		"                  # ",
		"#    ##    ##    ###",
		" #  #  #  #  #  #   ",
	}

	grid := aoc.NewGrid2D[bool](len(lines[0]), len(lines))
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				grid.AddXY(x, y, true)
			}
		}
	}
	return Tile{Grid2D: grid}
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
	W, H := t.Width, t.Height
	s := Tile{ID: t.ID, Grid2D: aoc.NewGrid2D[bool](W, H)}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			s.AddXY(x, y, t.GetXY(W-x-1, y))
		}
	}
	return s
}

func (t Tile) Rotate() Tile {
	W, H := t.Width, t.Height
	s := Tile{ID: t.ID, Grid2D: aoc.NewGrid2D[bool](H, W)}
	for y := 0; y < W; y++ {
		for x := 0; x < H; x++ {
			s.AddXY(x, y, t.GetXY(W-y-1, x))
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
