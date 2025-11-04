package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var names = in.Split(",")
	var N = len(names)
	in.Line() // blank line
	var instructions = in.SplitS[string](",")

	for _, instruction := range instructions {
		dir, n := instruction.Byte(), instruction.Int()

		var other int
		if dir == 'R' {
			other = Modulo(n, N)
		}
		if dir == 'L' {
			other = Modulo(N-n, N)
		}

		names[0], names[other] = names[other], names[0]
	}
	fmt.Println(names[0])
}
