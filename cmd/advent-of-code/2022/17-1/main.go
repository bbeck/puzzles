package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	input := lib.InputToString()

	var room = Room{0x1FF}
	var bindex, iindex = -1, -1
	for n := 0; n < 2022; n++ {
		x, y := 3, room.Height()+4 // +3 for gap, +1 for floor
		bindex = (bindex + 1) % len(Blocks)

		for {
			iindex = (iindex + 1) % len(input)
			switch input[iindex] {
			case '<':
				if !room.Overlaps(Blocks[bindex], x-1, y) {
					x--
				}

			case '>':
				if !room.Overlaps(Blocks[bindex], x+1, y) {
					x++
				}
			}

			if room.Overlaps(Blocks[bindex], x, y-1) {
				break
			}

			y--
		}

		room.Add(Blocks[bindex], x, y)
	}

	fmt.Println(room.Height())
}

type Room []uint16

func (r *Room) Height() int {
	height := len(*r) - 1
	for y := len(*r) - 1; y >= 0; y-- {
		if (*r)[y] != 0x101 {
			break
		}
		height--
	}
	return height
}

// Add adds a block to the room at the given coordinate.  The y coordinate is
// the *bottom* of the block.
func (r *Room) Add(b Block, x, y int) {
	r.Grow()
	for by := 0; by < len(b); by++ {
		(*r)[y+by] |= b[by] >> x
	}
	r.Grow()
}

func (r *Room) Overlaps(b Block, x, y int) bool {
	r.Grow()
	for by := 0; by < len(b); by++ {
		if (*r)[y+by]&(b[by]>>x) != 0 {
			return true
		}
	}
	return false
}

func (r *Room) Grow() {
	height := r.Height()
	for len(*r)-height < 8 {
		*r = append(*r, 0x101)
	}
}

type Block []uint16

var Blocks = []Block{
	// These are "upside down" because they're added from bottom to top.  They are
	// also padded to 9 bits so an add just shifts right by the x-coordinate.
	{0b111100000},
	{0b010000000, 0b111000000, 0b010000000},
	{0b111000000, 0b001000000, 0b001000000},
	{0b100000000, 0b100000000, 0b100000000, 0b100000000},
	{0b110000000, 0b110000000},
}
