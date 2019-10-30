package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	packages := aoc.InputToInts(2015, 24)
	total := Sum(packages)

	var best []int
	for size := 1; best == nil && size < len(packages)-2; size++ {
		aoc.EnumerateCombinations(len(packages), size, func(indices []int) {
			group := make([]int, size)
			for i := 0; i < size; i++ {
				group[i] = packages[indices[i]]
			}

			if Sum(group) == total/3 {
				if best == nil || Entanglement(group) < Entanglement(best) {
					best = group
				}
			}
		})
	}

	fmt.Printf("best: %+v\n", best)
	fmt.Printf("entanglement: %d\n", Entanglement(best))
}

func Sum(ns []int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}

	return sum
}

func Entanglement(group []int) int {
	product := 1
	for _, pkg := range group {
		product *= pkg
	}

	return product
}
