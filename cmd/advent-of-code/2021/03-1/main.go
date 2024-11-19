package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	lines := lib.InputToLines()

	counters := make([]lib.FrequencyCounter[rune], len(lines[0]))
	for _, line := range lines {
		for i, c := range line {
			counters[i].Add(c)
		}
	}

	var gamma, epsilon int
	for _, counter := range counters {
		entries := counter.Entries()
		gamma = (gamma << 1) + int(entries[0].Value-'0')
		epsilon = (epsilon << 1) + int(entries[1].Value-'0')
	}
	fmt.Println(gamma * epsilon)
}
