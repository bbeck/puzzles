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
		c1 := string(s[i])
		c2 := string(s[(i+1+N)%N])

		if c1 == c2 {
			sum += aoc.ParseInt(c1)
		}
	}

	fmt.Printf("sum: %d\n", sum)
}
