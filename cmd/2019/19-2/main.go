package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	W, H := 100, 100

	// Find some point in this row that's affected by the tractor beam.  Generally
	// speaking the points around the line y=1.1365x seem to be the ones affected
	// so we'll start there and then broaden our search outwards.
	find := func(y int) int {
		x := int(float64(y) / 1.1365)
		for dx := 0; ; dx++ {
			if Affected(x+dx, y) {
				return x + dx
			}
			if Affected(x-dx, y) {
				return x - dx
			}
		}
	}

	// Determine the x coordinates of the first and last point in this row that's
	// affected by the tractor beam.  We'll do this by performing two binary
	// searches.
	bounds := func(y int) (int, int) {
		mid := find(y)

		// Search for left most point
		L, R := 0, mid
		for L <= R {
			x := (L + R) / 2
			if Affected(x, y) {
				R = x - 1
			} else {
				L = x + 1
			}
		}
		min := L

		// Search for right most point
		L, R = mid, y*1000
		for L <= R {
			x := (L + R) / 2
			if Affected(x, y) {
				L = x + 1
			} else {
				R = x - 1
			}
		}
		max := R

		return min, max
	}

	// Determine whether or not a box with dimensions WxH fits within the tractor
	// beam at the given top left coordinate.
	fits := func(x, y int, w, h int) bool {
		for dy := 0; dy < h; dy++ {
			minX, maxX := bounds(y + dy)

			if x < minX {
				return false
			}

			if x > maxX {
				return false
			}

			if x+w-1 < minX {
				return false
			}

			if x+w-1 > maxX {
				return false
			}
		}

		return true
	}

	var foundX, foundY int
outer:
	for y := H - 10; ; y++ {
		minX, maxX := bounds(y)
		for x := minX; x < maxX; x++ {
			if fits(x, y, W, H) {
				foundX, foundY = x, y
				break outer
			}
		}
	}

	fmt.Println("value:", foundX*10000+foundY)
}

var memo = make(map[aoc.Point2D]bool)

func Affected(x, y int) bool {
	if x < 0 {
		return false
	}
	if y < 0 {
		return false
	}

	p := aoc.Point2D{X: x, Y: y}
	if value, ok := memo[p]; ok {
		return value
	}

	inputs := make(chan int, 2)
	inputs <- x
	inputs <- y
	close(inputs)

	var output int
	cpu := &CPU{
		memory: InputToMemory(2019, 19),
		input:  func(addr int) int { return <-inputs },
		output: func(value int) { output = value },
	}
	cpu.Execute()

	value := output == 1
	memo[p] = value
	return value
}
