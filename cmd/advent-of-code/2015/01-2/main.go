package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	position := 0
	floor := 0
	for p, c := range lib.InputToString() {
		position = p + 1

		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			break
		}
	}

	fmt.Println(position)
}
