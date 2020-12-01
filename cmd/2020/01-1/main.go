package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 1)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			if ns[i]+ns[j] == 2020 {
				fmt.Println(ns[i] * ns[j])
			}
		}
	}
}
