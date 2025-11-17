package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var dragon Point2D
	var sheep Set[Point2D]
	grid := in.ToGrid2D(func(x, y int, s string) string {
		if s == "D" {
			dragon = Point2D{X: x, Y: y}
		}
		if s == "S" {
			sheep.Add(Point2D{X: x, Y: y})
		}
		return s
	})

	var safe = Make2D[bool](grid.Width, grid.Height)
	for y := grid.Height - 1; y >= 0; y-- {
		for x := range grid.Width {
			if (y == grid.Height-1 || safe[x][y+1]) && grid.Get(x, y) == "#" {
				safe[x][y] = true
			}
		}
	}

	moves := func(x, y int) [][2]int {
		var moves [][2]int
		for dx := -2; dx <= 2; dx++ {
			for dy := -2; dy <= 2; dy++ {
				if Abs(dx)+Abs(dy) == 3 && grid.InBounds(x+dx, y+dy) {
					moves = append(moves, [2]int{x + dx, y + dy})
				}
			}
		}
		return moves
	}

	state := NewState(sheep, dragon)
	fmt.Println(Count(grid, state, "sheep", moves, safe))
}

var seen = map[string]map[State]int{
	"dragon": make(map[State]int),
	"sheep":  make(map[State]int),
}

func Count(grid Grid2D[string], s State, turn string, moves func(x, y int) [][2]int, safe [][]bool) int {
	if value, ok := seen[turn][s]; ok {
		return value
	}

	helper := func() int {
		var dragonX, dragonY = s.GetDragon()
		if turn == "sheep" {
			if !s.HasSheep() {
				return 1
			}

			var moved bool
			var count int
			for x := range grid.Width {
				var y = s.GetSheep(x)
				if y >= grid.Height {
					continue
				}
				if y+1 == grid.Height || safe[x][y+1] {
					moved = true
					continue
				}
				if grid.Get(x, y+1) == "#" || dragonX != x || dragonY != y+1 {
					moved = true
					count += Count(grid, s.MoveSheep(x), "dragon", moves, safe)
				}
			}

			if !moved {
				return Count(grid, s, "dragon", moves, safe)
			}
			return count
		}

		if turn == "dragon" {
			var dragonX, dragonY = s.GetDragon()

			var count int
			for _, d := range moves(dragonX, dragonY) {
				if grid.Get(d[0], d[1]) != "#" && s.GetSheep(d[0]) == d[1] {
					count += Count(grid, s.RemoveSheep(d[0]).MoveDragon(d[0], d[1]), "sheep", moves, safe)
				} else {
					count += Count(grid, s.MoveDragon(d[0], d[1]), "sheep", moves, safe)
				}
			}
			return count
		}

		panic("should never get here")
	}

	seen[turn][s] = helper()
	return seen[turn][s]
}

// State is a compact representation of the grid long with the locations of the
// sheep and dragon.  There are 7 columns that can contain sheep with 6 rows.
// For simplicity, we'll use 4-bits per sheep to represent its row with 0xFFFF
// meaning the row doesn't have a sheep in it.  We'll also represent the
// dragon's position using 4-bits for each coordinate.
type State uint64

func NewState(sheep Set[Point2D], dragon Point2D) State {
	var state uint64 = 0xFFFF_FFFA_AAAA_AA00
	for s := range sheep {
		var shift = 64 - 4*(s.X+1)
		state = (state &^ (0xF << shift)) | uint64(s.Y<<shift)
	}

	state |= uint64(dragon.X<<4) | uint64(dragon.Y)
	return State(state)
}

func (s State) HasSheep() bool {
	var mask uint64 = 0xFFFF_FFF0_0000_0000
	return (uint64(s) & mask) != mask
}

func (s State) GetSheep(col int) int {
	var shift = 64 - 4*(col+1)
	return int((s >> shift) & 0xF)
}

func (s State) MoveSheep(col int) State {
	var shift = 64 - 4*(col+1)
	var next = State(1 + ((s >> shift) & 0xF))
	return (s &^ (0xF << shift)) | (next << shift)
}

func (s State) RemoveSheep(col int) State {
	var shift = 64 - 4*(col+1)
	return s | (0xF << shift)
}

func (s State) GetDragon() (int, int) {
	var x = (s & 0xF0) >> 4
	var y = s & 0x0F
	return int(x), int(y)
}

func (s State) MoveDragon(x, y int) State {
	var next = State((x << 4) | y)
	return (s &^ 0xFF) | next
}
