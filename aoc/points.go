package aoc

import (
	"fmt"
	"math"
	"math/big"
)

type Point2D struct {
	X, Y int
}

func (p Point2D) Up() Point2D {
	return Point2D{X: p.X, Y: p.Y - 1}
}

func (p Point2D) Down() Point2D {
	return Point2D{X: p.X, Y: p.Y + 1}
}

func (p Point2D) Left() Point2D {
	return Point2D{p.X - 1, p.Y}
}

func (p Point2D) Right() Point2D {
	return Point2D{p.X + 1, p.Y}
}

func (p Point2D) South() Point2D {
	return Point2D{p.X, p.Y - 1}
}

func (p Point2D) West() Point2D {
	return Point2D{p.X - 1, p.Y}
}

func (p Point2D) North() Point2D {
	return Point2D{p.X, p.Y + 1}
}

func (p Point2D) East() Point2D {
	return Point2D{p.X + 1, p.Y}
}

func (p Point2D) ManhattanDistance(q Point2D) int {
	dx := p.X - q.X
	if dx < 0 {
		dx = -dx
	}

	dy := p.Y - q.Y
	if dy < 0 {
		dy = -dy
	}

	return dx + dy
}

// Slope computes the slope of the line between two points.  The slope is
// returned as two integers, the rise (dy) and the run (dx).  The returned
// slope will be reduced so there are no common factors between the rise and the
// run other than 1.
func (p Point2D) Slope(q Point2D) (int, int) {
	dx := int64(q.X - p.X)
	if dx == 0 {
		return 1, 0
	}

	dy := int64(q.Y - p.Y)
	slope := big.NewRat(dy, dx)
	return int(slope.Num().Int64()), int(slope.Denom().Int64())
}

func (p Point2D) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func GetBounds(points []Point2D) (int, int, int, int) {
	minX := math.MaxInt64
	maxX := 0
	minY := math.MaxInt64
	maxY := 0
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	return minX, minY, maxX, maxY
}
