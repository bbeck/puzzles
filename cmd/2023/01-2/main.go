package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	nums := puz.InputLinesTo[int](2023, 1, func(line string) int {
		L := len(line)

		var first, last int
		for i := range line {
			for prefix, num := range digits {
				if first == 0 && strings.HasPrefix(line[i:], prefix) {
					first = num
				}

				if last == 0 && strings.HasPrefix(line[L-i-1:], prefix) {
					last = num
				}
			}
		}

		return 10*first + last
	})

	fmt.Println(puz.Sum(nums...))
}
