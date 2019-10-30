package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	N := aoc.InputToInt(2018, 14)
	recipes := []int{3, 7}
	limit := len(recipes) + N + 10
	elf1, elf2 := 0, 1

	for len(recipes) < limit {
		recipes = append(recipes, Digits(recipes[elf1]+recipes[elf2])...)
		elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
		elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)
	}

	for i := N; i < N+10; i++ {
		fmt.Print(recipes[i])
	}
	fmt.Println()
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

func ShowRecipes(recipes []int, elf1, elf2 int) string {
	var builder strings.Builder
	for i, r := range recipes {
		if i == elf1 || i == elf2 {
			builder.WriteString(fmt.Sprintf("[%d] ", r))
		} else {
			builder.WriteString(fmt.Sprintf(" %d  ", r))
		}
	}

	return builder.String()
}
