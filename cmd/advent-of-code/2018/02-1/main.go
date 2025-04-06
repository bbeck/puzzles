package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum2, sum3 int

	for in.HasNext() {
		var line = in.Line()

		var counter FrequencyCounter[rune]
		for _, c := range line {
			counter.Add(c)
		}

		var has2, has3 bool
		for _, entry := range counter.Entries() {
			has2 = has2 || entry.Count == 2
			has3 = has3 || entry.Count == 3
		}

		if has2 {
			sum2++
		}
		if has3 {
			sum3++
		}
	}

	fmt.Println(sum2 * sum3)
}
