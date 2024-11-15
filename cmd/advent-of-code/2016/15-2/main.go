package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
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

	tm := puz.ChineseRemainderTheorem(offsets, positions)
	fmt.Println(tm)
}

type Disc struct {
	Size, Offset int
}

func InputToDiscs() []Disc {
	return puz.InputLinesTo(2016, 15, func(line string) Disc {
		line = strings.ReplaceAll(line, "Disc #", "")
		line = strings.ReplaceAll(line, "has ", "")
		line = strings.ReplaceAll(line, "positions; at time=", "")
		line = strings.ReplaceAll(line, ", it is at position", "")
		line = strings.ReplaceAll(line, ".", "")

		var id, size, tm, offset int
		fmt.Sscanf(line, "%d %d %d %d", &id, &size, &tm, &offset)

		// This only works because all of our disk offsets are specified at tm=0.
		return Disc{Size: size, Offset: offset}
	})
}
