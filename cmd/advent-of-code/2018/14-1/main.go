package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	N := lib.InputToInt()
	recipes := []int{3, 7}

	elf1, elf2 := 0, 1
	for N+10-len(recipes) > 0 {
		sum := recipes[elf1] + recipes[elf2]
		recipes = append(recipes, lib.Digits(sum)...)
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}

	var scores int
	for _, d := range recipes[N : N+10] {
		scores = scores*10 + d
	}
	fmt.Println(scores)
}
