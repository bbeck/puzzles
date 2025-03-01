package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	digits := Digits(in.Int())
	for range 50 {
		digits = LookAndSay(digits)
	}

	fmt.Println(len(digits))
}

func LookAndSay(s []int) []int {
	var output []int

	last, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
			continue
		}

		output = append(output, []int{count, last}...)
		last = s[i]
		count = 1
	}

	output = append(output, []int{count, last}...)
	return output
}
