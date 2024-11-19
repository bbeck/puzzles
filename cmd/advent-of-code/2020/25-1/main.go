package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	card, door := InputToKeys()

	var loop int
	for loop = 1; ; loop++ {
		if lib.ModPow(7, loop, 20201227) == card {
			break
		}
	}
	fmt.Println(lib.ModPow(door, loop, 20201227))
}

func InputToKeys() (int, int) {
	ns := lib.InputToInts()
	return ns[0], ns[1]
}
