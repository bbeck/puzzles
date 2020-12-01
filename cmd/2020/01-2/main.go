package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 1)
	for i := 0; i < len(ns); i++ {
		for j := i + 1; j < len(ns); j++ {
			for k := j + 1; k < len(ns); k++ {
				if ns[i]+ns[j]+ns[k] == 2020 {
					fmt.Println(ns[i] * ns[j] * ns[k])
				}
			}
		}
	}
}
