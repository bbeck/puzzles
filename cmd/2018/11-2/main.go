package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	serial := aoc.InputToInt(2018, 11)
	N := 300

	cells := make(map[aoc.Point2D]int)
	for x := 1; x <= N; x++ {
		for y := 1; y <= N; y++ {
			rack := x + 10
			power := (((rack*y)+serial)*rack)/100%10 - 5
			cells[aoc.Point2D{X: x, Y: y}] = power
		}
	}

	// A series of partial sums that sum all of the cells to the right of the
	// current cell.
	var sums [][]int
	sums = append(sums, []int{})
	for y := 1; y <= N; y++ {
		sums = append(sums, make([]int, N+1))

		var sum int
		for x := 1; x <= N; x++ {
			sum += cells[aoc.Point2D{x, y}]
		}

		for x := 1; x <= N; x++ {
			v := cells[aoc.Point2D{x, y}]
			sums[y][x] = sum
			sum -= v
		}
	}

	var best = math.MinInt64
	var bestP aoc.Point2D
	var bestN int
	for n := 1; n <= N; n++ {
		for x := 1; x <= N-n; x++ {
			for y := 1; y <= N-n; y++ {
				p := aoc.Point2D{X: x, Y: y}
				if total := Total(sums, p, n); total > best {
					best = total
					bestP = p
					bestN = n
				}
			}
		}
	}

	fmt.Printf("square with top left %s and n of %d has best power level of %d\n", bestP, bestN, best)
	fmt.Printf("%d,%d,%d\n", bestP.X, bestP.Y, bestN)
}

func Total(sums [][]int, p aoc.Point2D, n int) int {
	var sum int
	for dy := 0; dy < n; dy++ {
		sum += sums[p.Y+dy][p.X] - sums[p.Y+dy][p.X+n]
	}

	return sum
}
