package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var N int // The dimension of a tile, including the border (10)
var D int // The dimension, in number of tiles, of the image along one axis (12)
var B int // The dimension, in pixels, of the image along one axis (96)

func main() {
	tiles := InputToTiles(2020, 20)
	N = tiles[0].width
	D = int(math.Sqrt(float64(len(tiles))))
	B = (N - 2) * D

	image := BuildBitmap(BuildImage(tiles))

	// Attempt to find sea monsters in the image, keeping track of which cells
	// were occupied by a found sea monster.
	occupied := aoc.NewSet()
	for y := 0; y < B; y++ {
		for x := 0; x < B; x++ {
			for _, monster := range SeaMonster.Transformed() {
				var found = true
				var ps []aoc.Point2D

				for dy := 0; dy < monster.height; dy++ {
					for dx := 0; dx < monster.width; dx++ {
						p := aoc.Point2D{X: x + dx, Y: y + dy}
						if monster.Get(dx, dy) {
							ps = append(ps, p)
							if !image[p] {
								found = false
							}
						}
					}
				}
				if found {
					for _, p := range ps {
						occupied.Add(p)
					}
				}
			}
		}
	}

	// Now that we know where the sea monsters are, count the number of active
	// cells that don't have a sea monster in them.
	var roughness int
	for p, bit := range image {
		if bit && !occupied.Contains(p) {
			roughness++
		}
	}
	fmt.Println(roughness)
}

const (
	T = true
	F = false
)

var SeaMonster = Tile{
	id: "Monster",
	data: []bool{
		F, F, F, F, F, F, F, F, F, F, F, F, F, F, F, F, F, F, F, T, F,
		F, T, F, F, F, F, T, T, F, F, F, F, T, T, F, F, F, F, T, T, T,
		F, F, T, F, F, T, F, F, T, F, F, T, F, F, T, F, F, T, F, F, F,
	},
	width:  21,
	height: 3,
}

// Convert the bits of the image into a bitmap being sure to remove the bits
// corresponding to the tile borders.
func BuildBitmap(image map[aoc.Point2D]*Tile) map[aoc.Point2D]bool {
	bitmap := make(map[aoc.Point2D]bool)
	for y := 0; y < D; y++ {
		for x := 0; x < D; x++ {
			tile := image[aoc.Point2D{X: x, Y: y}]

			for dy := 1; dy < N-1; dy++ {
				for dx := 1; dx < N-1; dx++ {
					p := aoc.Point2D{
						X: (N-2)*x + dx - 1,
						Y: (N-2)*y + dy - 1,
					}

					bitmap[p] = tile.Get(dx, dy)
				}
			}
		}
	}

	return bitmap
}

func BuildImage(tiles []Tile) map[aoc.Point2D]*Tile {
	// First, find a corner of the image.
	var c Tile
	for _, tile := range tiles {
		if IsCorner(tile, tiles) {
			c = tile
			break
		}
	}

	// Now that we know a corner, use it to build up the rest of the image.  We'll
	// start with this being the top left corner, and then keep filling in
	// neighbors until we have a complete image.  We'll have to try each
	// transformed version of this corner piece though to ensure that we find the
	// version of it that works as the top left corner of the image.
outer:
	for _, corner := range c.Transformed() {
		image := map[aoc.Point2D]*Tile{
			aoc.Point2D{X: 0, Y: 0}: &corner,
		}

		used := aoc.NewSet()
		used.Add(corner.id)

		for y := 0; y < D; y++ {
			for x := 0; x < D; x++ {
				p := aoc.Point2D{X: x, Y: y}
				if image[p] != nil {
					continue
				}

				tile := FindTile(image[p.Left()], image[p.Up()], tiles, used)
				if tile == nil {
					continue outer
				}

				image[p] = tile
				used.Add(tile.id)
			}
		}

		return image
	}

	return nil
}

func FindTile(left, above *Tile, tiles []Tile, used aoc.Set) *Tile {
	for _, tile := range tiles {
		if used.Contains(tile.id) {
			continue
		}

		for _, transformed := range tile.Transformed() {
			if left != nil && !FitsHorizontally(*left, transformed) {
				continue
			}

			if above != nil && !FitsVertically(*above, transformed) {
				continue
			}

			return &transformed
		}
	}

	return nil
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
		if l.Get(N-1, y) != r.Get(0, y) {
			return false
		}
	}
	return true
}

func FitsVertically(t, b Tile) bool {
	for x := 0; x < N; x++ {
		if t.Get(x, N-1) != b.Get(x, 0) {
			return false
		}
	}
	return true
}

type Tile struct {
	id            string
	data          []bool
	width, height int
}

func (t Tile) Get(x, y int) bool {
	return t.data[y*t.width+x]
}

func (t Tile) RotateLeft() Tile {
	W, H := t.width, t.height

	transformed := make([]bool, len(t.data))
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			transformed[x*H+y] = t.Get(W-x-1, y)
		}
	}

	return Tile{
		id:     t.id,
		data:   transformed,
		width:  t.height,
		height: t.width,
	}
}

func (t Tile) FlipH() Tile {
	transformed := make([]bool, len(t.data))
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			transformed[y*t.width+x] = t.Get(t.width-x-1, y)
		}
	}

	return Tile{
		id:     t.id,
		data:   transformed,
		width:  t.width,
		height: t.height,
	}
}

func (t Tile) FlipV() Tile {
	transformed := make([]bool, len(t.data))
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			transformed[y*t.width+x] = t.Get(x, t.height-y-1)
		}
	}

	return Tile{
		id:     t.id,
		data:   transformed,
		width:  t.width,
		height: t.height,
	}
}

func (t Tile) Transformed() []Tile {
	A := t
	B := A.RotateLeft()
	C := B.RotateLeft()
	D := C.RotateLeft()
	E := A.FlipH()
	F := E.RotateLeft()
	G := F.RotateLeft()
	H := G.RotateLeft()
	I := A.FlipV()
	J := I.RotateLeft()
	K := J.RotateLeft()
	L := K.RotateLeft()
	return []Tile{A, B, C, D, E, F, G, H, I, J, K, L}
}

func (t Tile) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Tile: %s\n", t.id))
	for y := 0; y < t.height; y++ {
		for x := 0; x < t.width; x++ {
			if t.Get(x, y) {
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
	blocks := make(map[string][]string)

	var current string
	for _, line := range aoc.InputToLines(year, day) {
		if len(line) == 0 {
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			current = strings.ReplaceAll(strings.ReplaceAll(line, "Tile ", ""), ":", "")
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
			id:     id,
			data:   data,
			width:  len(lines[0]),
			height: len(lines),
		})
	}

	return tiles
}
