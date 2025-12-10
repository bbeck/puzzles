package main

import (
	"fmt"
	"maps"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ps []Point2D
	for in.HasNext() {
		ps = append(ps, Point2D{X: in.Int(), Y: in.Int()})
	}

	var areas Set[int]
	var coordinates = make(map[int][][2]int)
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			dx, dy := Abs(ps[i].X-ps[j].X)+1, Abs(ps[i].Y-ps[j].Y)+1
			area := dx * dy

			areas.Add(area)
			coordinates[area] = append(coordinates[area], [2]int{i, j})
		}
	}

outer:
	for _, area := range Reversed(slices.Sorted(maps.Keys(areas))) {
		for _, c := range coordinates[area] {
			p, q := ps[c[0]], ps[c[1]]

			if IsInsidePolygon(p, q, ps) {
				fmt.Println(area)
				break outer
			}
		}
	}
}

func IsInsidePolygon(a, b Point2D, ps []Point2D) bool {
	// For a rectangle to be fully contained within a polygon a couple of
	// conditions must be true:
	//
	// 1. The edges of the rectangle must not intersect with the edges of the
	//    polygon.  In this particular problem they can touch but not intersect.
	//
	// 2. An interior point of the rectangle must be contained within the polygon.
	//    Only one point needs to be checked since the rectangle doesn't interset
	//    the polygon anywhere.
	//
	// We'll do both of these checks simultaneously since they both require
	// iterating over all edges of the polygon.
	//
	// For the intersection constraint we'll use a standard line segment
	// intersection algorithm.  For the point inside of polygon we'll use a
	// ray-casting algorithm that envisions a ray from the point going off
	// infinitely to the left.  If the ray intersects the polygon an odd number of
	// times then our point is inside the polygon.

	// Determine the edges of the rectangle
	left, right := Min(a.X, b.X), Max(a.X, b.X)
	top, bottom := Min(a.Y, b.Y), Max(a.Y, b.Y)

	// Determine the midpoint of the rectangle, this is the point we'll require to
	// be inside the polygon.
	midpoint := Point2D{X: (a.X + b.X) / 2, Y: (a.Y + b.Y) / 2}

	var count int
	for i := 0; i < len(ps); i++ {
		p, q := ps[i], ps[(i+1)%len(ps)]
		isVert, isHoriz := p.X == q.X, p.Y == q.Y

		// Check for intersection with the rectangle
		if isVert && left < p.X && p.X < right && Overlaps(top, bottom, p.Y, q.Y) {
			return false
		}
		if isHoriz && top < p.Y && p.Y < bottom && Overlaps(left, right, p.X, q.X) {
			return false
		}

		// Check for intersection with the ray.  Since the ray is horizontal we only
		// need to check vertical edges of the polygon.
		if isVert && p.X < midpoint.X && Overlaps(midpoint.Y, midpoint.Y, p.Y, q.Y) {
			count++
		}
	}

	return count%2 == 1
}

func Overlaps(a1, a2, b1, b2 int) bool {
	a1, a2 = Min(a1, a2), Max(a1, a2)
	b1, b2 = Min(b1, b2), Max(b1, b2)
	return a1 < b2 && b1 < a2
}
