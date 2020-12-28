package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = 9 // The largest cup value
const MOVES = 100

func main() {
	cups := aoc.NewRing()
	for _, c := range aoc.InputToString(2020, 23) {
		digit := aoc.ParseInt(string(c))
		cups.InsertAfter(digit)
	}

	// Move back to the beginning of the ring
	cups.Next()

	for move := 1; move <= MOVES; move++ {
		current := cups.Current().(int)

		// Remove the next 3 elements.
		cups.Next()
		c1 := cups.Remove().(int)
		c2 := cups.Remove().(int)
		c3 := cups.Remove().(int)

		// Remember where to jump back to after we finish our insertion.
		next := cups.Current().(int)

		// Determine where we're going to add the removed elements back.
		destination := current - 1
		for destination == c1 || destination == c2 || destination == c3 || destination < 1 {
			destination--
			if destination < 1 {
				destination = N
			}
		}

		// Go to the destination and add back our removed elements.
		cups.JumpTo(destination)
		cups.InsertAfter(c1)
		cups.InsertAfter(c2)
		cups.InsertAfter(c3)

		// Return back to where we started.
		cups.JumpTo(next)
	}

	// Gather the values of the cups starting after the 1.
	cups.JumpTo(1)
	for cup := cups.Next().(int); cup != 1; cup = cups.Next().(int) {
		fmt.Print(cup)
	}
	fmt.Println()
}
