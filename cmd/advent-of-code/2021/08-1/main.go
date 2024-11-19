package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
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
	return lib.InputLinesTo(func(line string) []string {
		_, rhs, _ := strings.Cut(line, " | ")
		return strings.Fields(rhs)
	})
}
