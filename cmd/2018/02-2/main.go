package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	report := func(a, b string) {
		for i := 0; i < len(a); i++ {
			if a[i] == b[i] {
				fmt.Print(string(a[i]))
			}
		}
		fmt.Println()
	}

	lines := aoc.InputToLines(2018, 2)
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if Distance(lines[i], lines[j]) == 1 {
				report(lines[i], lines[j])
			}
		}
	}
}

func Distance(a, b string) int {
	var dist int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist
}
