package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2020, 1)
	for i, a := range ns {
		for j, b := range ns[i+1:] {
			for _, c := range ns[j+1:] {
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
