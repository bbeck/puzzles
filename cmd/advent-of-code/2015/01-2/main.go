package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var position, floor int
	for in.HasNext() {
		position++

		switch in.Byte() {
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
