package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	desired := puz.InputToInt()

	var house int
	for house = 1; ; house++ {
		if NumPresentsDelivered(house) >= desired {
			break
		}
	}
	fmt.Println(house)
}

func NumPresentsDelivered(n int) int {
	// The number of elves that deliver presents to a house is equal to the
	// sum of the divisors of that house number.
	sqrt := int(math.Sqrt(float64(n)))

	var sum int
	for divisor1 := 1; divisor1 <= sqrt+1; divisor1++ {
		if n%divisor1 == 0 {
			sum += divisor1

			if divisor2 := n / divisor1; divisor1 != divisor2 {
				sum += divisor2
			}
		}
	}
	return sum * 10
}
