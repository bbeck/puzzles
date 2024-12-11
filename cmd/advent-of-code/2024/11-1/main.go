package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	stones := InputToStones()
	for n := 0; n < 25; n++ {
		stones = Next(stones)
	}

	fmt.Println(len(stones))
}

func Next(stones []int) []int {
	var next []int
	for _, stone := range stones {
		if stone == 0 {
			next = append(next, 1)
			continue
		}

		if digits := Digits(stone); len(digits)%2 == 0 {
			lhs := digits[:len(digits)/2]
			rhs := digits[len(digits)/2:]
			next = append(next, JoinDigits(lhs))
			next = append(next, JoinDigits(rhs))
			continue
		}

		next = append(next, stone*2024)
	}

	return next
}

func InputToStones() []int {
	var ns []int
	for _, s := range strings.Fields(InputToString()) {
		ns = append(ns, ParseInt(s))
	}

	return ns
}
