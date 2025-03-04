package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	// The problem is asking us to solve a system of congruences:
	//   tm = offset_1 (mod positions_1)
	//   tm = offset_2 (mod positions_2)
	//   ...
	//
	// To do this we can use the Chinese Remainder Theorem as long as we adjust
	// each offset to account for the amount of time it takes for the ball to
	// reach that disc.
	discs := InputToDiscs()
	discs = append(discs, Disc{Size: 11, Offset: 0})

	var offsets, positions []int
	for i, disc := range discs {
		offsets = append(offsets, disc.Size-(disc.Offset+i+1))
		positions = append(positions, disc.Size)
	}

	tm := ChineseRemainderTheorem(offsets, positions)
	fmt.Println(tm)
}

type Disc struct {
	Size, Offset int
}

func InputToDiscs() []Disc {
	return in.LinesTo(func(in *in.Scanner[Disc]) Disc {
		var id, size, offset int
		in.Scanf("Disc #%d has %d positions; at time=0, it is at position %d.", &id, &size, &offset)

		// This only works because all of our disk offsets are specified at tm=0.
		return Disc{Size: size, Offset: offset}
	})
}
