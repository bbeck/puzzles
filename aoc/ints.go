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

// AbsInt returns the absolute value of the passed in integer.
func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// GCD returns the greatest common divisor of two integers.
func GCD(a, b int) int {
	if a == 0 {
		return b
	}
	return GCD(b%a, a)
}

// LCM returns the least common multiple of two integers.
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
