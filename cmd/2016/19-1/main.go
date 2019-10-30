package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	n := aoc.InputToInt(2016, 19)

	elves := aoc.NewRing()
	for i := 0; i < n; i++ {
		elves.InsertAfter(&Elf{id: i + 1, presents: 1})
	}
	elves.Next()

	for i := 0; i < n-1; i++ {
		current := elves.Current().(*Elf)
		next := elves.Next().(*Elf)
		current.presents += next.presents
		elves.Remove()
	}

	current := elves.Current().(*Elf)
	fmt.Printf("elf %d has all of the presents (%d)\n", current.id, current.presents)
}

type Elf struct {
	id       int
	presents int
}
