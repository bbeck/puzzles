package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	lines := aoc.InputToLines(2016, 6)

	var entries []Entry
	for i := 0; i < len(lines[0]); i++ {
		entries = append(entries, make(Entry))
	}

	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			entries[i][line[i]]++
		}
	}

	fmt.Print("message: ")
	for i := 0; i < len(entries); i++ {
		fmt.Printf("%c", entries[i].LeastFrequent())
	}
	fmt.Println()
}

type Entry map[byte]int

func (e Entry) LeastFrequent() byte {
	var best byte
	var bestCount int = math.MaxInt64
	for b, count := range e {
		if count < bestCount {
			best = b
			bestCount = count
		}
	}

	return best
}
