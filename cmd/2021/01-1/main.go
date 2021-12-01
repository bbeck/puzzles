package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2021, 1)

	var count int
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
