package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	card, door := InputToKeys()

	var loop int
	for loop = 1; ; loop++ {
		if puz.ModPow(7, loop, 20201227) == card {
			break
		}
	}
	fmt.Println(puz.ModPow(door, loop, 20201227))
}

func InputToKeys() (int, int) {
	ns := puz.InputToInts(2020, 25)
	return ns[0], ns[1]
}
