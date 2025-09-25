package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	dots := []int{1, 3, 5, 10, 15, 16, 20, 24, 25, 30, 37, 38, 49, 50, 74, 75, 100, 101}

	seen := map[int]int{0: 0}

	var sum int
	for in.HasNext() {
		beetle, best := in.Int(), math.MaxInt
		for b1 := beetle/2 - 100; b1 < beetle/2+100; b1++ {
			b2 := beetle - b1
			if Abs(b2-b1) > 100 {
				continue
			}
			c1 := FindMin(b1, dots, seen)
			c2 := FindMin(b2, dots, seen)
			best = Min(best, c1+c2)
		}
		sum += best
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
