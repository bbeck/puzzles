package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	N := len(aoc.InputToLines(2021, 3))
	ones := make(map[int]int)
	for _, line := range aoc.InputToLines(2021, 3) {
		for i, c := range line {
			if c == '1' {
				ones[i] += 1
			}
		}
	}

	var gamma, epsilon int
	for i := 0; i < len(ones); i++ {
		if ones[i] >= N/2 {
			gamma = gamma<<1 + 1
			epsilon = epsilon << 1
		} else {
			gamma = gamma << 1
			epsilon = epsilon<<1 + 1
		}
	}

	fmt.Println(gamma * epsilon)
}
