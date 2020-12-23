package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

const N = 9 // The largest cup value
const MOVES = 100

func main() {
	var digits []int
	for n := aoc.InputToInt(2020, 23); n > 0; n /= 10 {
		digits = append(digits, n%10)
	}

	cups := aoc.NewRing()
	for i := len(digits) - 1; i >= 0; i-- {
		cups.InsertAfter(digits[i])
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
