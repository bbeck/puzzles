package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

const p = 20201227 // prime

func main() {
	cardKey, doorKey := InputToKeys(2020, 25)

	// Compute the card's loop size
	var cardLoop int
	for loop := 1; ; loop++ {
		if PowMod(7, loop, p) == cardKey {
			cardLoop = loop
			break
		}
	}

	// Use the card's loop size with the door key to compute the secret
	// encryption key
	fmt.Println(PowMod(doorKey, cardLoop, p))
}

// Compute x^y mod p using Fermat's Little Theorem.
func PowMod(x, y, p int) int {
	value := 1
	for y > 0 {
		if y%2 == 1 {
			value = (value * x) % p
		}

		y = y >> 1
		x = (x * x) % p
	}

	return value
}

func InputToKeys(year, day int) (int, int) {
	ns := aoc.InputToInts(year, day)
	return ns[0], ns[1]
}
