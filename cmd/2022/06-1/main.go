package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

const N = 4

func main() {
	s := aoc.InputToString(2022, 6)
	for end := N; end < len(s); end++ {
		bs := []byte(s[end-N : end])
		if len(aoc.SetFrom(bs...)) == N {
			fmt.Println(end)
			break
		}
	}
}
