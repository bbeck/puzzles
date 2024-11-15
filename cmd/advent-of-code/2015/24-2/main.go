package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
)

func main() {
	packages := puz.InputToInts()
	target := puz.Sum(packages...) / 4

	// Find all candidates for the passenger compartment.  This partition should
	// contain the fewest packages possible, so once a partition is found we'll
	// reject any later partitions that have more packages.
	//
	// Once we find a partition we'll consider its entanglement relative to the
	// best entanglement we've already identified.  If it has the possibility of
	// improving our entanglement then we'll check to make sure it's a valid
	// partition.  It's a valid partition if we are able to find 2 subpartitions
	// of the remaining packages that also sum to our target weight.
	best := math.MaxInt
	size := math.MaxInt
	FindPartitions(target, packages, func(partition, remaining []int) bool {
		if len(partition) > size {
			return true
		}

		size = len(partition)
		entanglement := puz.Product(partition...)
		if entanglement >= best {
			return true
		}

		// Check to see if we can find another partition in the remaining packages.
		// If we can then this is a candidate that we should consider.
		var found bool
		FindPartitions(target, remaining, func(partition, remainder []int) bool {
			// We need one more subpartition.
			FindPartitions(target, remaining, func(partition, remaining []int) bool {
				found = true
				return found
			})

			return found
		})

		// We didn't find valid subpartitions, keep searching.
		if !found {
			return false
		}

		// This is a valid partition, and we already know it has a better
		// entanglement than our current best.
		best = entanglement
		return true
	})

	fmt.Println(best)
}

func FindPartitions(weight int, packages []int, fn func(partition, remaining []int) bool) {
	all := puz.SetFrom(packages...)

	var done bool
	for size := 1; !done && size < len(packages); size++ {
		puz.EnumerateCombinations(len(packages), size, func(indices []int) bool {
			partition := make([]int, 0, size)
			for _, index := range indices {
				partition = append(partition, packages[index])
			}

			if puz.Sum(partition...) == weight {
				remaining := all.DifferenceElems(partition...)
				done = fn(partition, remaining.Entries())
			}

			return done
		})
	}
}
