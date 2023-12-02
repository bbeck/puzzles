package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	digits := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	nums := aoc.InputLinesTo[int](2023, 1, func(line string) (int, error) {
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

		return 10*first + last, nil
	})

	fmt.Println(aoc.Sum(nums...))
}
