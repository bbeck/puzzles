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

outer:
	for length := 1; length <= len(digits)/2; length++ {
		if len(digits)%length != 0 {
			continue
		}

		for i := 0; i < len(digits); i++ {
			for j := i; j < len(digits); j += length {
				if digits[i] != digits[j] {
					continue outer
				}
			}
		}

		return false
	}

	return true
}
