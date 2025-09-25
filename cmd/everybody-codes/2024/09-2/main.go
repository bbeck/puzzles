package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dots := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30}

	var sum int
	for in.HasNext() {
		sum += FindMin(in.Int(), dots, map[int]int{0: 0})
	}
	fmt.Println(sum)
}

func FindMin(beetle int, dots []int, seen map[int]int) int {
	if value, ok := seen[beetle]; ok {
		return value
	}

	best := math.MaxInt
	for _, dot := range dots {
		if dot > beetle {
			continue
		}

		if count := FindMin(beetle-dot, dots, seen); count < math.MaxInt {
			best = Min(best, count+1)
		}
	}

	seen[beetle] = best
	return best
}
