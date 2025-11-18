package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	ducks := in.Ints()

	var increment = map[bool]int{true: 1}

	var rounds int
	var changed = true
	for ; changed; rounds += increment[changed] {
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
	for ; changed; rounds += increment[changed] {
		changed = false
		for col := 0; col < len(ducks)-1; col++ {
			if ducks[col+1] > ducks[col] {
				ducks[col+1]--
				ducks[col]++
				changed = true
			}
		}
	}

	fmt.Println(rounds)
}
