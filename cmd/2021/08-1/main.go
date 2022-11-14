package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var count int
	for _, digits := range InputToOutputDigits() {
		for _, digit := range digits {
			n := len(digit)
			if n == 2 || n == 3 || n == 4 || n == 7 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func InputToOutputDigits() [][]string {
	return aoc.InputLinesTo(2021, 8, func(line string) ([]string, error) {
		_, rhs, _ := strings.Cut(line, " | ")
		return strings.Fields(rhs), nil
	})
}
