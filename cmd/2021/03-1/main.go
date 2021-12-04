package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ns := InputToInts()

	// Determine how many bits are in the input.  Since the numbers are now sorted it's the
	// number of bits required to store the largest number.
	B := int(math.Ceil(math.Log2(float64(ns[len(ns)-1]))))

	// Build gamma taking the most popular digit of each bit from the input numbers
	var gamma int
	for bit := B - 1; bit >= 0; bit-- {
		mask := 1 << bit

		ones := 0
		for _, n := range ns {
			if n&mask != 0 {
				ones++
			}
		}

		gamma = gamma << 1
		if ones > len(ns)/2 {
			gamma = gamma + 1
		}
	}

	// Invert the bits of gamma to compute epsilon
	epsilon := (1<<B - 1) - gamma

	fmt.Println(gamma * epsilon)
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
