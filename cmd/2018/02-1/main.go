package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum2, sum3 int
	for _, line := range aoc.InputToLines(2018, 2) {
		var counter aoc.FrequencyCounter[rune]
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
