package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var digits []int
	for _, c := range puz.InputToString(2015, 10) {
		digits = append(digits, puz.ParseInt(string(c)))
	}

	for i := 0; i < 50; i++ {
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
