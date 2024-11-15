package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	floor := 0
	for _, c := range puz.InputToString() {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}
	}

	fmt.Println(floor)
}
