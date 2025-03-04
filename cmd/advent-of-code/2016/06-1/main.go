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
		password = append(password, counter.Entries()[0].Value)
	}
	fmt.Println(string(password))
}
