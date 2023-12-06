package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	times, distances := InputToTimesAndDistances()

	wins := make([]int, len(times))
	for race := 0; race < len(times); race++ {
		time, distance := times[race], distances[race]

		for hold := 1; hold < time; hold++ {
			remaining := distance - hold*(time-hold)
			if remaining < 0 {
				wins[race]++
			}
		}
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
