package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var counters [8]FrequencyCounter[rune]
	for in.HasNext() {
		for i, b := range in.String() {
			counters[i].Add(b)
		}
	}

	var password []rune
	for _, counter := range counters {
		entries := counter.Entries()
		password = append(password, entries[len(entries)-1].Value)
	}
	fmt.Println(string(password))
}
