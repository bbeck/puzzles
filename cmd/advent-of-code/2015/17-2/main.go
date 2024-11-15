package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	sizes := puz.InputToInts()

	best := math.MaxInt
	ways := make(map[int]int)
	EnumerateWays(sizes, func(containers []bool) {
		var count int
		for _, container := range containers {
			if container {
				count++
			}
		}

		best = puz.Min(best, count)
		ways[count]++
	})

	fmt.Println(ways[best])
}

func EnumerateWays(sizes []int, fn func([]bool)) {
	containers := make([]bool, len(sizes))

	var helper func(index int, remaining int)
	helper = func(index int, remaining int) {
		// If we haven't used too much eggnog, and still have containers to
		// consider, then consider what happens if we do/don't use the
		// current container.
		if remaining >= 0 && index < len(sizes) {
			containers[index] = true
			helper(index+1, remaining-sizes[index])
			containers[index] = false
			helper(index+1, remaining)
			return
		}

		// If we've reached the last container, then check to see if we've
		// used exactly the amount of eggnog we needed to.  If we have then
		// let our caller know.
		if remaining == 0 {
			fn(containers)
		}
	}

	helper(0, 150)
}
