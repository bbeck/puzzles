package aoc

import (
	"fmt"
	"math"
	"math/big"
)

// Point2D is a rectangular representation of a point in 2D space.
type Point2D struct {
	X, Y int
}

// Origin2D is the point that lies on the origin.
var Origin2D = Point2D{X: 0, Y: 0}

// Up returns the point above the current one when the points are from a
// screen based coordinate system with the origin in the top left.
func (p Point2D) Up() Point2D {
	return Point2D{X: p.X, Y: p.Y - 1}
}

// Down returns the point below the current one when the points are from
// a screen based coordinate system with the origin in the top left.
func (p Point2D) Down() Point2D {
	return Point2D{X: p.X, Y: p.Y + 1}
}

// Left returns the point to the left of the current one when the points
// are from a screen based coordinate system with the origin in the top
// left.
func (p Point2D) Left() Point2D {
	return Point2D{p.X - 1, p.Y}
}

// Right returns the point to the right of the current one when the points
// are from a screen based coordinate system with the origin in the top left.
func (p Point2D) Right() Point2D {
	return Point2D{p.X + 1, p.Y}
}

// Move returns the adjacent point along the specified heading.
func (p Point2D) Move(h Heading) Point2D {
	switch h {
	case Up:
		return p.Up()
	case Down:
		return p.Down()
	case Left:
		return p.Left()
	case Right:
		return p.Right()
	default:
		return p
	}
}

// MoveN returns the point after moving along the specified heading n times.
func (p Point2D) MoveN(h Heading, n int) Point2D {
	for n > 0 {
		p = p.Move(h)
		n--
	}
	return p
}

// OrthogonalNeighbors returns a slice of neighbors that are orthogonal to
// the current point.
func (p Point2D) OrthogonalNeighbors() []Point2D {
	return []Point2D{
		p.Down(),
		p.Left(),
		p.Right(),
		p.Up(),
	}
}

// Neighbors returns a slice of all neighbors to the current point.
func (p Point2D) Neighbors() []Point2D {
	return []Point2D{
		p.Up().Left(), p.Up(), p.Up().Right(),
		p.Left(), p.Right(),
		p.Down().Left(), p.Down(), p.Down().Right(),
	}
}

// ManhattanDistance computes the distance between the current point and the
// provided point when traveling along a rectilinear path between the points.
func (p Point2D) ManhattanDistance(q Point2D) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y)
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

// GetBounds determines the corners of the bounding box that contains all
// specified points.  The bounding box corners returned are the top left and
// bottom right corners.
func GetBounds(ps []Point2D) (Point2D, Point2D) {
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt
	for _, p := range ps {
		minX, maxX = Min(minX, p.X), Max(maxX, p.X)
		minY, maxY = Min(minY, p.Y), Max(maxY, p.Y)
	}

	return Point2D{X: minX, Y: minY}, Point2D{X: maxX, Y: maxY}
}

// Point3D is a rectangular representation of a point in 3D space.
type Point3D struct {
	X, Y, Z int
}

// Origin3D is the point that lies on the origin.
var Origin3D = Point3D{X: 0, Y: 0, Z: 0}

// OrthogonalNeighbors returns a slice of neighbors that are orthogonal to
// the current point.
func (p Point3D) OrthogonalNeighbors() []Point3D {
	return []Point3D{
		{p.X - 1, p.Y, p.Z},
		{p.X + 1, p.Y, p.Z},
		{p.X, p.Y - 1, p.Z},
		{p.X, p.Y + 1, p.Z},
		{p.X, p.Y, p.Z - 1},
		{p.X, p.Y, p.Z + 1},
	}
}

// Neighbors returns a slice of all neighbors to the current point.
func (p Point3D) Neighbors() []Point3D {
	var neighbors []Point3D
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}

				n := Point3D{X: p.X + dx, Y: p.Y + dy, Z: p.Z + dz}
				neighbors = append(neighbors, n)
			}
		}
	}

	return neighbors
}

// ManhattanDistance computes the distance between the current point and the
// provided point when traveling along a rectilinear path between the points.
func (p Point3D) ManhattanDistance(q Point3D) int {
	return Abs(p.X-q.X) + Abs(p.Y-q.Y) + Abs(p.Z-q.Z)
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

// GetBounds3D determines the corners of the bounding box that contains all
// specified points.  The bounding box corners returned are the top left and
// bottom right corners.
func GetBounds3D(ps []Point3D) (Point3D, Point3D) {
	minX, minY, minZ, maxX, maxY, maxZ := math.MaxInt, math.MaxInt, math.MaxInt, math.MinInt, math.MinInt, math.MinInt
	for _, p := range ps {
		minX, maxX = Min(minX, p.X), Max(maxX, p.X)
		minY, maxY = Min(minY, p.Y), Max(maxY, p.Y)
		minZ, maxZ = Min(minZ, p.Z), Max(maxZ, p.Z)
	}

	return Point3D{X: minX, Y: minY, Z: minZ}, Point3D{X: maxX, Y: maxY, Z: maxZ}
}
