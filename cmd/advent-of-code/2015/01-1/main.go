package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var floor int
	for in.HasNext() {
		switch in.Byte() {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
}
