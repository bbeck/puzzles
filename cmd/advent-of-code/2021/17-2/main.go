package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	minX, maxX, minY, maxY := InputToTargetArea()

	// Brute force a solution.  The x-velocity should be between 1 and maxX, anything
	// larger will skip over the target area entirely in the first second.  The y-velocity
	// is a bit trickier to bound, so we just use a hardcoded range.
	var solutions []puz.Point2D
	for vx0 := 1; vx0 <= maxX; vx0++ {
		for vy0 := -500; vy0 < 500; vy0++ {
			px, py := 0, 0
			vx, vy := vx0, vy0

			for px <= maxX && minY <= py {
				px, py = px+vx, py+vy
				vx, vy = puz.Max(0, vx-1), vy-1

				if minX <= px && px <= maxX && minY <= py && py <= maxY {
					solutions = append(solutions, puz.Point2D{X: vx0, Y: vy0})
					break
				}
			}
		}
	}

	fmt.Println(len(solutions))
}

func InputToTargetArea() (int, int, int, int) {
	s := puz.InputToString(2021, 17)

	var minX, maxX, minY, maxY int
	fmt.Sscanf(s, "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	return minX, maxX, minY, maxY
}
