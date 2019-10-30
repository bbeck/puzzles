package aoc

import (
	"fmt"
	"math"
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
