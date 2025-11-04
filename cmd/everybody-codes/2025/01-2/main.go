package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var names Ring[string]
	for _, name := range in.Split(",") {
		names.InsertAfter(name)
	}
	names.Next() // Get back to the first name

	in.Line() // blank line

	for _, instruction := range in.SplitS[string](",") {
		dir, n := instruction.Byte(), instruction.Int()
		if dir == 'R' {
			names.NextN(n)
		}
		if dir == 'L' {
			names.PrevN(n)
		}
	}

	fmt.Println(names.Current())
}
