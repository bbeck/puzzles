package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var names = in.Split(",")
	in.Line() // blank line
	var instructions = in.SplitS[string](",")

	var current int
	for _, instruction := range instructions {
		dir, n := instruction.Byte(), instruction.Int()
		if dir == 'R' {
			current = Clamp(current+n, 0, len(names)-1)
		}
		if dir == 'L' {
			current = Clamp(current-n, 0, len(names)-1)
		}
	}

	fmt.Println(names[current])
}
