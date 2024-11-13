package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz/cpus"
	"math"
	"sort"
)

func main() {
	// The rectangular region that's closest to the origin that contains will
	// always have it's lower left corner touching the left side of the tractor
	// beam.  Given this we'll use a binary search to choose the Y coordinate of
	// the lower left corner to find the smallest Y such that the entire box is
	// contained with the tractor beam.
	y := sort.Search(10000, func(y int) bool {
		x := FindMinimumX(y)
		return IsInTractorBeam(x+99, y-99)
	})

	x := FindMinimumX(y)
	fmt.Println(x*10000 + y - 99)
}

func FindMinimumX(y int) int {
	// Through some experimentation it seems that points along the line y = 1.3x
	// are always within the tractor beam.  So we'll perform a binary search
	// using that as our limit knowing that the minimum is always to the left.
	x0 := int(math.Ceil(float64(y) / 1.3))
	return sort.Search(x0, func(x int) bool {
		return IsInTractorBeam(x, y)
	})
}

func IsInTractorBeam(x, y int) bool {
	inputs := make(chan int, 2)
	inputs <- x
	inputs <- y

	var output bool
	cpu := cpus.IntcodeCPU{
		Memory: cpus.InputToIntcodeMemory(2019, 19),
		Input:  func() int { return <-inputs },
		Output: func(value int) { output = value == 1 },
	}
	cpu.Execute()

	return output
}
