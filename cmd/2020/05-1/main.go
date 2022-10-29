package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var max int
	for _, line := range aoc.InputToLines(2020, 5) {
		max = aoc.Max(max, aoc.ParseIntWithBase(replacer.Replace(line), 2))
	}

	fmt.Println(max)
}
