package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	calories := lib.InputLinesTo(func(line string) int {
		if line == "" {
			return 0
		}
		return lib.ParseInt(line)
	})

	var best int
	for _, group := range lib.Split(calories, func(n int) bool { return n != 0 }) {
		best = lib.Max(best, lib.Sum(group...))
	}
	fmt.Println(best)
}
