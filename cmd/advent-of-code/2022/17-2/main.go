package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

const N = 1000000000000

func main() {
	input := puz.InputToString()

	var heights []int          // heights after each turn
	turns := make(map[Key]int) // memoization table
	var first, last int        // first and last turns in the cycle

	room := &Room{0x1FF}
	bindex, iindex := -1, -1
	for i := 0; ; i++ {
		key := Key{
			BlockIndex: bindex,
			InputIndex: iindex,
			Top:        room.Top(),
		}

		bindex, iindex = Drop(room, input, bindex, iindex)
		heights = append(heights, room.Height())

		if seen, ok := turns[key]; ok {
			first, last = seen, i
			break
		}
		turns[key] = i
	}

	// length of the cycle and how much height we gain per iteration
	length, delta := last-first, heights[last]-heights[first]

	// how many full iterations around the cycle we need to take
	iters := (N - first) / length

	// how many steps in the final partial iteration we need to take
	partial := N - iters*length - first

	// total height gained is a combination of:
	// - the height we enter the cycle at
	// - the height we gain through full iterations of the cycle
	// - the height we gain in the final partial iteration of the cycle
	fmt.Println(heights[first] + iters*delta + (heights[first+partial] - heights[first] - 1))
}

type Key struct {
	BlockIndex int
	InputIndex int
	Top        [30]uint16
}

func Drop(room *Room, input string, bindex, iindex int) (int, int) {
	bindex = (bindex + 1) % len(Blocks)
	x, y := 3, room.Height()+4 // +3 for gap, +1 for floor

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
	return bindex, iindex
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

func (r *Room) Top() [30]uint16 {
	var top [30]uint16

	y := r.Height()
	for i := 0; i < 30; i++ {
		y--
		if y < 0 {
			break
		}
		top[i] = (*r)[y]
	}

	return top
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
