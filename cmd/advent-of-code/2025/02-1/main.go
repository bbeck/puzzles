package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var sum int
	for _, in := range in.SplitS[any](",") {
		var start, end int
		in.Scanf("%d-%d", &start, &end)

		for n := start; n <= end; n++ {
			if !Valid(n) {
				sum += n
			}
		}
	}
	fmt.Println(sum)
}

func Valid(n int) bool {
	digits := Digits(n)
	if len(digits)%2 == 1 {
		return true
	}

	HL := len(digits) / 2
	for i := 0; i < HL; i++ {
		if digits[i] != digits[HL+i] {
			return true
		}
	}
	return false
}
