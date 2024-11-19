package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
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
