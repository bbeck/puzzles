package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
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

	tm := ChineseRemainderTheorem(offsets, positions)
	fmt.Println(tm)
}

func ChineseRemainderTheorem(as, ns []int) int {
	var prod = 1
	for _, n := range ns {
		prod *= n
	}

	var sum int
	for i := 0; i < len(as); i++ {
		p := prod / ns[i]
		sum += as[i] * MulInv(p, ns[i]) * p
	}

	for sum < 0 {
		sum += prod
	}

	return sum % prod
}

func MulInv(a, b int) int {
	if b == 1 {
		return 1
	}

	x0, x1 := 0, 1
	for a > 1 {
		q := a / b
		a, b = b, a%b
		x0, x1 = x1-q*x0, x0
	}

	return x1
}

type Disc struct {
	Size, Offset int
}

func InputToDiscs() []Disc {
	return aoc.InputLinesTo(2016, 15, func(line string) (Disc, error) {
		line = strings.ReplaceAll(line, "Disc #", "")
		line = strings.ReplaceAll(line, "has ", "")
		line = strings.ReplaceAll(line, "positions; at time=", "")
		line = strings.ReplaceAll(line, ", it is at position", "")
		line = strings.ReplaceAll(line, ".", "")

		var id, size, tm, offset int
		fmt.Sscanf(line, "%d %d %d %d", &id, &size, &tm, &offset)

		// This only works because all of our disk offsets are specified at tm=0.
		return Disc{Size: size, Offset: offset}, nil
	})
}
