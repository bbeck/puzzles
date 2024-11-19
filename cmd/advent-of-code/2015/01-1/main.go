package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	floor := 0
	for _, c := range lib.InputToString() {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
}
