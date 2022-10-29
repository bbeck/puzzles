package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	card, door := InputToKeys()

	var loop int
	for loop = 1; ; loop++ {
		if ModPow(7, loop, 20201227) == card {
			break
		}
	}
	fmt.Println(ModPow(door, loop, 20201227))
}

func ModPow(b, e, m int) int {
	b = b % m

	result := 1
	for e > 0 {
		if e%2 == 1 {
			result = (result * b) % m
		}
		e >>= 1
		b = (b * b) % m
	}

	return result
}

func InputToKeys() (int, int) {
	ns := aoc.InputToInts(2020, 25)
	return ns[0], ns[1]
}
