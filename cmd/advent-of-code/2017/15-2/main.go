package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	previous := InputToPreviousValues()
	factors := []int{16807, 48271}

	var count int
	for n := 0; n < 5_000_000; n++ {
		previous[0] = Next(previous[0], factors[0], 4)
		previous[1] = Next(previous[1], factors[1], 8)

		if previous[0]&0b11111111_11111111 == previous[1]&0b11111111_11111111 {
			count++
		}
	}
	fmt.Println(count)
}

func Next(previous, factor, mod int) int {
	for {
		previous = (previous * factor) % 2147483647
		if previous%mod == 0 {
			return previous
		}
	}
}

func InputToPreviousValues() []int {
	return in.LinesToS(func(in in.Scanner[int]) int {
		var id string
		var value int
		in.Scanf("Generator %s starts with %d", &id, &value)
		return value
	})
}
