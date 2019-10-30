package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var digits []int
	for _, d := range aoc.InputToString(2018, 14) {
		digits = append(digits, aoc.ParseInt(string(d)))
	}
	D := len(digits)

	same := func(d1, d2 []int) bool {
		if len(d1) != len(d2) {
			return false
		}

		for i := 0; i < len(d1); i++ {
			if d1[i] != d2[i] {
				return false
			}
		}

		return true
	}

	recipes := []int{3, 7}
	elf1, elf2 := 0, 1

	for {
		ds := Digits(recipes[elf1] + recipes[elf2])
		recipes = append(recipes, ds...)
		N := len(recipes)

		if N >= D && same(recipes[N-D:N], digits) {
			fmt.Printf("idx: %d\n", N-D)
			break
		}

		if N >= D+1 && same(recipes[N-D-1:N-1], digits) {
			fmt.Printf("idx: %d\n", N-D-1)
			break
		}

		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}
}

func Digits(n int) []int {
	if n == 0 {
		return []int{0}
	}

	var digits []int
	for n != 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}

	return digits
}
