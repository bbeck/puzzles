package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	s := aoc.InputToString(2017, 1)
	N := len(s)

	var sum int
	for i := 0; i < N; i++ {
		j := (i + 1 + N) % N
		if s[i] == s[j] {
			sum += aoc.ParseInt(string(s[i]))
		}
	}
	fmt.Println(sum)
}
