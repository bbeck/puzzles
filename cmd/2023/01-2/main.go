package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	nums := aoc.InputLinesTo[int](2023, 1, func(line string) (int, error) {
		numbers := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
			"0":     0,
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

		var nums []int
		for i := range line {
			for prefix, num := range numbers {
				if strings.HasPrefix(line[i:], prefix) {
					nums = append(nums, num)
				}
			}
		}

		return nums[0]*10 + nums[len(nums)-1], nil
	})

	fmt.Println(aoc.Sum(nums...))
}
