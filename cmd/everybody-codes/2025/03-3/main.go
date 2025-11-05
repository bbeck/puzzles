package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var fc FrequencyCounter[int]
	for _, n := range in.Ints() {
		fc.Add(n)
	}

	first := fc.Entries()[0]
	fmt.Println(first.Count)
}
