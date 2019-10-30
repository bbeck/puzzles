package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	discs := InputToDiscs(2016, 15)
	discs = append(discs, Disc{11, 0})

	for start := 0; ; start++ {
		success := true
		for offset := 1; offset <= len(discs); offset++ {
			if !discs[offset-1].IsOpen(start + offset) {
				success = false
				break
			}
		}

		if success {
			fmt.Printf("success at tm: %d\n", start)
			break
		}
	}
}

type Disc struct {
	positions int
	position  int
}

func (d Disc) IsOpen(tm int) bool {
	return (d.position+tm)%d.positions == 0
}

func InputToDiscs(year, day int) []Disc {
	var discs []Disc
	for _, line := range aoc.InputToLines(2016, 15) {
		var id, positions, position int
		_, err := fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &id, &positions, &position)
		if err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		discs = append(discs, Disc{positions, position})
	}

	return discs
}
