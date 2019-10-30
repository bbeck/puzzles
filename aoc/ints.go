package aoc

import (
	"log"
	"strconv"
)

// ParseInt parses a string as an integer value.
func ParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("unable to parse integer: '%s'", s)
	}

	return n
}

// MinInt determines the minimum integer of a set of integers passed as
// arguments to the function.
func MinInt(i int, is ...int) int {
	min := i
	for _, i := range is {
		if i < min {
			min = i
		}
	}

	return min
}

// MaxInt determines the maximum integer of a set of integers passed as
// arguments to the function.
func MaxInt(i int, is ...int) int {
	max := i
	for _, i := range is {
		if i > max {
			max = i
		}
	}

	return max
}
