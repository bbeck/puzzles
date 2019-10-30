package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var sum2, sum3 int
	for _, line := range aoc.InputToLines(2018, 2) {
		c2, c3 := Analyze(line)
		sum2 += c2
		sum3 += c3
	}

	fmt.Printf("checksum: %d\n", sum2*sum3)
}

func Analyze(s string) (int, int) {
	counts := make(map[string]int)
	for _, c := range s {
		counts[string(c)]++
	}

	var has2, has3 int
	for _, v := range counts {
		if v == 2 {
			has2 = 1
		}
		if v == 3 {
			has3 = 1
		}
	}

	return has2, has3
}
