package aoc

import (
	"log"
	"strconv"
)

// ParseIntWithBase parses a string as an integer value in a specific base.
func ParseIntWithBase(s string, base int) int {
	n, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		log.Fatalf("unable to parse integer (base %d): '%s'", base, s)
	}
	return int(n)
}

// ParseInt parses a string as an integer value.
func ParseInt(s string) int {
	return ParseIntWithBase(s, 10)
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
