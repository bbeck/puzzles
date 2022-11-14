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
