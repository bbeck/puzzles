package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	n := in.Int()

	var elves Ring[*Elf]
	for i := range n {
		elves.InsertAfter(&Elf{ID: i + 1, Presents: 1})
	}
	elves.Next()

	for range n - 1 {
		current := elves.Current()
		next := elves.Next()
		current.Presents += next.Presents
		elves.Remove()
	}

	fmt.Println(elves.Current().ID)
}

type Elf struct {
	ID       int
	Presents int
}
