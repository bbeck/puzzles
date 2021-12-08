package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	var count int
	for _, digit := range InputToOutputDigits() {
		n := len(digit)
		if n == 2 || n == 3 || n == 4 || n == 7 {
			count++
		}
	}
	fmt.Println(count)
}

func InputToOutputDigits() []string {
	var digits []string
	for _, line := range aoc.InputToLines(2021, 8) {
		fields := strings.Split(line, " | ")
		outputs := strings.Split(fields[1], " ")

		digits = append(digits, outputs...)
	}
	return digits
}
