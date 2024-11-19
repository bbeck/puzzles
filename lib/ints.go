package lib

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

// GCD returns the greatest common divisor of a series of integers.
func GCD(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	var gcd func(int, int) int
	gcd = func(a, b int) int {
		if a == 0 {
			return b
		}
		return gcd(b%a, a)
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = gcd(res, nums[i])
	}
	return res
}

// LCM returns the least common multiple of a series of integers.
func LCM(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = (res * nums[i]) / GCD(res, nums[i])
	}
	return res
}
