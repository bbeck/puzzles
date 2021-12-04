package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
	"sort"
)

func main() {
	ns := InputToInts()

	// Determine how many bits are in the input.  Since the numbers are now sorted it's the
	// number of bits required to store the largest number.
	B := int(math.Ceil(math.Log2(float64(ns[len(ns)-1]))))

	// Oxygen rating wants to build a new number using the most common bit value of all
	// numbers that share the current prefix.
	prefix := 0
	for b := B - 1; len(ns) > 1 && b >= 0; b-- {
		zeroes, ones := count(ns, b)

		var bit int
		if ones >= zeroes {
			bit = 1
		}

		prefix |= bit << b
		ns = filter(ns, func(n int) bool { return n&(1<<b)>>b != bit })
	}
	oxygen := ns[0]

	// CO2 rating wants to build a new number using the least common bit value of all
	// numbers that share the current prefix.
	ns = InputToInts()

	prefix = 0
	for b := B - 1; len(ns) > 1 && b >= 0; b-- {
		zeroes, ones := count(ns, b)

		var bit int
		if ones < zeroes {
			bit = 1
		}

		prefix |= bit << b
		ns = filter(ns, func(n int) bool { return n&(1<<b)>>b != bit })
	}
	co2 := ns[0]

	fmt.Println(oxygen * co2)
}

func InputToInts() []int {
	var ns []int
	for _, line := range aoc.InputToLines(2021, 3) {
		n := aoc.ParseIntWithBase(line, 2)
		ns = append(ns, n)
	}

	sort.Ints(ns)
	return ns
}

func count(ns []int, bit int) (int, int) {
	mask := 1 << bit

	var zeroes, ones int
	for _, n := range ns {
		if n&mask != 0 {
			ones++
		} else {
			zeroes++
		}
	}
	return zeroes, ones
}

func filter(ns []int, fn func(n int) bool) []int {
	f := make([]int, 0, len(ns))
	for _, n := range ns {
		if !fn(n) {
			f = append(f, n)
		}
	}
	return f
}
