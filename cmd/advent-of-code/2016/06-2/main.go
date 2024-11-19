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

	var password []rune
	for _, counter := range counters {
		entries := counter.Entries()
		password = append(password, entries[len(entries)-1].Value)
	}
	fmt.Println(string(password))
}
