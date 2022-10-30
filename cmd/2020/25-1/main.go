package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	card, door := InputToKeys()

	var loop int
	for loop = 1; ; loop++ {
		if aoc.ModPow(7, loop, 20201227) == card {
			break
		}
	}
	fmt.Println(aoc.ModPow(door, loop, 20201227))
}

func InputToKeys() (int, int) {
	ns := aoc.InputToInts(2020, 25)
	return ns[0], ns[1]
}
