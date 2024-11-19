package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
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
	return lib.InputLinesTo(func(line string) int {
		var id string
		var value int
		fmt.Sscanf(line, "Generator %s starts with %d", &id, &value)
		return value
	})
}
