package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := aoc.InputToInts(2021, 1)

	window := func(n int) int {
		return ns[n] + ns[n+1] + ns[n+2]
	}

	var count int
	for i := 1; i < len(ns)-2; i++ {
		if window(i) > window(i-1) {
			count++
		}
	}
	fmt.Println(count)
}
