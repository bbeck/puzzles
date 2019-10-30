package main

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

var parents = make(map[State]State)

func main() {
	// 0 = promethium
	// 1 = cobalt
	// 2 = curium
	// 3 = ruthenium
	// 4 = plutonium
	// 5 = elerium
	// 6 = dilithium

	// first floor:
	//   promethium generator
	//   promethium microchip
	//   elerium generator
	//   elerium microchip
	//   dilithium generator
	//   dilithium microchip
	//
	//      0123456x 0123456x
	f0 := 0b10000110_10000110

	// second floor:
	//   cobalt generator
	//   curium generator
	//   ruthenium generator
	//   plutonium generator
	//
	//      0123456x 0123456x
	f1 := 0b01111000_00000000

	// third floor:
	//   cobalt microchip
	//   curium microchip
	//   ruthenium microchip
	//   plutonium microchip
	//
	//      0123456x 0123456x
	f2 := 0b00000000_01111000

	// fourth floor:
	//    <empty>
	f3 := 0b00000000_00000000

	// elevator on floor 0
	eh, el := 0, 0

	start := State(f0<<48 | f1<<32 | f2<<16 | f3<<0 | eh<<8 | el)
	fmt.Printf("start: %s\n", start)
	goal := State(f0 | f1 | f2 | f3 | 0b00000001_00000001)
	fmt.Printf("goal: %s\n", goal)
	fmt.Println()

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(node aoc.Node) int {
		s := node.(State)
		f0, f1, f2, _ := s.Floors()
		return 3*((bits.OnesCount64(f0)+1)/2) +
			2*((bits.OnesCount64(f1)+1)/2) +
			(bits.OnesCount64(f2)+1)/2
	}

	visit := func(node aoc.Node) bool {
		return node.(State) == goal
	}

	path, distance, found := aoc.AStarSearch(start, visit, cost, heuristic)
	if !found {
		fmt.Println("no path found")
		return
	}

	fmt.Printf("path (%v):\n", distance)
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Println("  ", path[i])
	}
}

//
// There are 14 distinct items plus the location of the elevator.  This means
// we need 56 bits to represent which items are on which floor, and an
// additional 2 bits to represent which floor the elevator is on.
//
// We'll encode things as follows:
//   0000000x 0000000x 1111111x 1111111x 2222222x 2222222x 3333333E 3333333E
//     gens    chips     gens    chips     gens    chips     gens    chips
//
//
type State uint64

// These masks extract just the bits corresponding to the items on a specific
// floor.
var FloorMasks = []uint64{
	0xFEFE_0000_0000_0000,
	0x0000_FEFE_0000_0000,
	0x0000_0000_FEFE_0000,
	0x0000_0000_0000_FEFE,
}

// These masks are the elevator bits for each level.
var LevelMasks = []uint64{
	0b00000000_00000000,
	0b00000000_00000001,
	0b00000001_00000000,
	0b00000001_00000001,
}

// These are the masks for each single item that we can move.
var Move1Masks = []uint64{
	0b10000000_00000000,
	0b01000000_00000000,
	0b00100000_00000000,
	0b00010000_00000000,
	0b00001000_00000000,
	0b00000100_00000000,
	0b00000010_00000000,
	0b00000000_10000000,
	0b00000000_01000000,
	0b00000000_00100000,
	0b00000000_00010000,
	0b00000000_00001000,
	0b00000000_00000100,
	0b00000000_00000010,
}

// These are the masks for each set of two items we can move together.  We can
// move 2 generators, 2 chips, or a matched chip and generator.
var Move2Masks = []uint64{
	// two generators
	0b11000000_00000000,
	0b10100000_00000000,
	0b10010000_00000000,
	0b10001000_00000000,
	0b10000100_00000000,
	0b10000010_00000000,
	0b01100000_00000000,
	0b01010000_00000000,
	0b01001000_00000000,
	0b01000100_00000000,
	0b01000010_00000000,
	0b00110000_00000000,
	0b00101000_00000000,
	0b00100100_00000000,
	0b00100010_00000000,
	0b00011000_00000000,
	0b00010100_00000000,
	0b00010010_00000000,
	0b00001100_00000000,
	0b00001010_00000000,
	0b00000110_00000000,

	// two chips
	0b00000000_11000000,
	0b00000000_10100000,
	0b00000000_10010000,
	0b00000000_10001000,
	0b00000000_10000100,
	0b00000000_10000010,
	0b00000000_01100000,
	0b00000000_01010000,
	0b00000000_01001000,
	0b00000000_01000100,
	0b00000000_01000010,
	0b00000000_00110000,
	0b00000000_00101000,
	0b00000000_00100100,
	0b00000000_00100010,
	0b00000000_00011000,
	0b00000000_00010100,
	0b00000000_00010010,
	0b00000000_00001100,
	0b00000000_00001010,
	0b00000000_00000110,

	// matched generator and chip
	0b10000000_10000000,
	0b01000000_01000000,
	0b00100000_00100000,
	0b00010000_00010000,
	0b00001000_00001000,
	0b00000100_00000100,
	0b00000010_00000010,
}

func (s State) IsValid() bool {
	invalid := func(floor uint64) bool {
		generators := uint8(floor & 0xFE00 >> 8)
		chips := uint8(floor & 0x00FE)

		// clear the bits for all chips that have their generator
		chips = chips & ^generators

		// we're invalid if we have generators and there are chips without their
		// generator left over.
		return generators > 0 && chips > 0
	}

	f0, f1, f2, f3 := s.Floors()
	return !invalid(f0) && !invalid(f1) && !invalid(f2) && !invalid(f3)
}

func (s State) ID() string {
	return fmt.Sprintf("%d", s)
}

func (s State) Elevator() int {
	return int(((s & 0b0000_0001_0000_0000) >> 7) | (s & 0b0000_0000_0000_0001))
}

func (s State) Floors() (uint64, uint64, uint64, uint64) {
	f0 := (uint64(s) & FloorMasks[0]) >> 48
	f1 := (uint64(s) & FloorMasks[1]) >> 32
	f2 := (uint64(s) & FloorMasks[2]) >> 16
	f3 := (uint64(s) & FloorMasks[3]) >> 0
	return f0, f1, f2, f3
}

func (s State) Children() []aoc.Node {
	var neighbors []aoc.Node

	e := s.Elevator()

	var floor uint64
	switch e {
	case 0:
		floor, _, _, _ = s.Floors()
	case 1:
		_, floor, _, _ = s.Floors()
	case 2:
		_, _, floor, _ = s.Floors()
	case 3:
		_, _, _, floor = s.Floors()
	}

	move := func(mask uint64, from, to int) bool {
		if to < 0 || to > 3 {
			return false
		}

		f0, f1, f2, f3 := s.Floors()

		floors := []uint64{f0, f1, f2, f3}
		floors[from] &= ^mask
		floors[to] |= mask
		floors[3] = (floors[3] & 0b11111110_11111110) | LevelMasks[to]

		child := State(floors[0]<<48 | floors[1]<<32 | floors[2]<<16 | floors[3])
		if child.IsValid() {
			if _, ok := parents[child]; !ok {
				parents[child] = s
			}
			neighbors = append(neighbors, child)
			return true
		}

		return false
	}

	// When moving up, try to take two items first.  If we can then there's no
	// need to try to take just one.
	var found bool
	for _, mask := range Move2Masks {
		if floor&mask == mask {
			found = move(mask, e, e+1) || found
		}
	}

	// Check moves with just a single item if we weren't able to move two at once.
	if !found {
		for _, mask := range Move1Masks {
			if floor&mask == mask {
				move(mask, e, e+1)
			}
		}
	}

	// When moving down, try to take just a single item first.  If we can then
	// there's no need to try to take two.
	found = false
	for _, mask := range Move1Masks {
		if floor&mask == mask {
			found = move(mask, e, e-1) || found
		}
	}

	if !found {
		for _, mask := range Move2Masks {
			if floor&mask == mask {
				move(mask, e, e-1)
			}
		}
	}

	return neighbors
}

func (s State) String() string {
	floor := func(f uint64) string {
		var items []string
		for i := 0; i < 7; i++ {
			if f&(1<<(16-i-1)) > 0 {
				items = append(items, fmt.Sprintf("G%d", i))
			}
			if f&(1<<(8-i-1)) > 0 {
				items = append(items, fmt.Sprintf("M%d", i))
			}
		}
		sort.Strings(items)

		return strings.Join(items, ",")
	}

	f0, f1, f2, f3 := s.Floors()

	return fmt.Sprintf("E:%d F0:[%s] F1:[%s] F2:[%s] F3:[%s]",
		s.Elevator(), floor(f0), floor(f1), floor(f2), floor(f3))
}
