package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	minX, maxX, minY, maxY := InputToTargetArea()

	// Brute force a solution.  The x-velocity should be between 1 and maxX, anything
	// larger will skip over the target area entirely in the first second.  The y-velocity
	// is a bit trickier to bound, so we just use a hardcoded range.
	solutions := make(map[lib.Point2D]int)
	for vx0 := 1; vx0 <= maxX; vx0++ {
		for vy0 := -500; vy0 < 500; vy0++ {
			px, py := 0, 0
			vx, vy := vx0, vy0

			var maxH int
			for px <= maxX && minY <= py {
				px, py = px+vx, py+vy
				vx, vy = lib.Max(0, vx-1), vy-1
				maxH = lib.Max(maxH, py)

				if minX <= px && px <= maxX && minY <= py && py <= maxY {
					solutions[lib.Point2D{X: vx0, Y: vy0}] = maxH
					break
				}
			}
		}
	}

	fmt.Println(lib.Max(lib.GetMapValues(solutions)...))
}

func InputToTargetArea() (int, int, int, int) {
	s := lib.InputToString()

	var minX, maxX, minY, maxY int
	fmt.Sscanf(s, "target area: x=%d..%d, y=%d..%d", &minX, &maxX, &minY, &maxY)
	return minX, maxX, minY, maxY
}
