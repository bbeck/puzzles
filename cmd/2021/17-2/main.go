package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	minX, maxX, minY, maxY := InputToTargetArea()

	// Brute force a solution.  The x-velocity should be between 1 and maxX, anything
	// larger will skip over the target area entirely in the first second.  The y-velocity
	// is a bit trickier to bound, so we just use a hardcoded range.
	var solutions []aoc.Point2D
	for vx0 := 1; vx0 <= maxX; vx0++ {
		for vy0 := -500; vy0 < 500; vy0++ {
			p := aoc.Point2D{X: 0, Y: 0}
			v := aoc.Point2D{X: vx0, Y: vy0}

			// Keep simulating as long as the probe hasn't moved beyond the target in one axis,
			// keeping in mind that the y-coordinate is falling, so we compare to minY.
			for p.X <= maxX && p.Y >= minY {
				p = aoc.Point2D{X: p.X + v.X, Y: p.Y + v.Y}
				v = aoc.Point2D{X: aoc.MaxInt(v.X-1, 0), Y: v.Y - 1}

				if minX <= p.X && p.X <= maxX && minY <= p.Y && p.Y <= maxY {
					solutions = append(solutions, aoc.Point2D{X: vx0, Y: vy0})
					break
				}
			}
		}
	}

	fmt.Println(len(solutions))
}

func InputToTargetArea() (int, int, int, int) {
	s := aoc.InputToString(2021, 17)

	var minX, maxX, minY, maxY int
	if _, err := fmt.Sscanf(s, "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY); err != nil {
		log.Fatal(err)
	}

	return minX, maxX, minY, maxY
}
