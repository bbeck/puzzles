package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ducks := in.Ints()

	var rounds = 10
	var changed = true
	for ; changed && rounds >= 0; rounds-- {
		changed = false
		for col := 0; col < len(ducks)-1; col++ {
			if ducks[col] > ducks[col+1] {
				ducks[col]--
				ducks[col+1]++
				changed = true
			}
		}
	}

	changed = true
	for ; changed && rounds >= 0; rounds-- {
		changed = false
		for col := 0; col < len(ducks)-1; col++ {
			if ducks[col+1] > ducks[col] {
				ducks[col+1]--
				ducks[col]++
				changed = true
			}
		}
	}

	var checksum int
	for i, c := range ducks {
		checksum += (i + 1) * c
	}
	fmt.Println(checksum)
}
