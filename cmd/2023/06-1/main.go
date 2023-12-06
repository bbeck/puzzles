package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	times, distances := InputToTimesAndDistances()

	var wins []int
	for race := 0; race < len(times); race++ {
		time, distance := float64(times[race]), float64(distances[race])

		// We want distance - hold*(time - hold) < 0, thus we can solve for hold
		// using the quadratic formula. hold^2 - time*hold + distance < 0
		D := math.Sqrt(time*time - 4*distance)
		h1 := int(math.Ceil(time/2 - D/2))
		h2 := int(math.Floor(time/2 + D/2))

		// When the discriminant is an integer then we have a case where we tied
		// the record.  We need to adjust the roots in order to avoid the tie.
		if D == float64(int(D)) {
			h1 += 1
			h2 -= 1
		}

		wins = append(wins, h2-h1+1)
	}

	fmt.Println(aoc.Product(wins...))
}

func InputToTimesAndDistances() ([]int, []int) {
	var times, distances []int
	for _, line := range aoc.InputToLines(2023, 6) {
		var nums []int
		for _, field := range strings.Fields(line)[1:] {
			nums = append(nums, aoc.ParseInt(field))
		}

		if strings.HasPrefix(line, "Time") {
			times = nums
		} else {
			distances = nums
		}
	}

	return times, distances
}
