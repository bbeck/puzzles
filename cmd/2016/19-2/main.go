package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	n := aoc.InputToInt(2016, 19)

	// create each elf
	elves := make([]*Elf, n)
	for i := 0; i < n; i++ {
		elves[i] = &Elf{id: i + 1, presents: 1}
	}

	// now that all of the elves are created, link them together in a ring
	for i := 0; i < n; i++ {
		elves[i].prev = elves[(i-1+n)%n]
		elves[i].next = elves[(i+1+n)%n]
	}

	elf1, elf2 := elves[0], elves[n/2]
	for i := 0; i < n-1; i++ {
		// elf1 steals from elf2
		elf1.presents += elf2.presents

		// remove elf2
		elf2.prev.next = elf2.next
		elf2.next.prev = elf2.prev

		elf1 = elf1.next
		elf2 = elf2.next
		if i%2 == 1 {
			elf2 = elf2.next
		}
	}

	fmt.Printf("elf %d is left with %d presents\n", elf1.id, elf1.presents)
}

type Elf struct {
	id       int
	presents int

	prev, next *Elf
}
