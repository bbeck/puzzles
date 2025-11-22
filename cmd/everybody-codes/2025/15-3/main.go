package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	start, end, segments := InputToSegments()

	// Build a list of the x and y values that we can travel to.  These will
	// be the points along the edges of segments, plus the values used by
	// our start and end points.
	var xSet = SetFrom(start.X, end.X)
	var ySet = SetFrom(start.Y, end.Y)
	for _, segment := range segments {
		xSet.Add(segment.P1.X-1, segment.P1.X+1, segment.P2.X-1, segment.P2.X+1)
		ySet.Add(segment.P1.Y-1, segment.P1.Y+1, segment.P2.Y-1, segment.P2.Y+1)
	}
	xs, ys := xSet.Entries(), ySet.Entries()
	sort.Ints(xs)
	sort.Ints(ys)

	// Now perform a search where at a given point we consider moving to the next
	// row or column in each direction.

	var endNeighbors = SetFrom(end.Up(), end.Right(), end.Down(), end.Left())
	children := func(p Point2D) []Point2D {
		// If we're one step away from the end then we're done
		if endNeighbors.Contains(p) {
			return []Point2D{end}
		}

		xIndex, yIndex := sort.SearchInts(xs, p.X), sort.SearchInts(ys, p.Y)

		var ts []Segment
		if xIndex > 0 {
			ts = append(ts, Segment{P1: p, P2: Point2D{X: xs[xIndex-1], Y: p.Y}})
		}
		if xIndex+1 < len(xs) {
			ts = append(ts, Segment{P1: p, P2: Point2D{X: xs[xIndex+1], Y: p.Y}})
		}
		if yIndex > 0 {
			ts = append(ts, Segment{P1: p, P2: Point2D{X: p.X, Y: ys[yIndex-1]}})
		}
		if yIndex+1 < len(ys) {
			ts = append(ts, Segment{P1: p, P2: Point2D{X: p.X, Y: ys[yIndex+1]}})
		}

		var children []Point2D
		for _, t := range ts {
			var intersects bool
			for _, s := range segments {
				if Intersects(s, t) {
					intersects = true
					break
				}
			}

			if !intersects {
				children = append(children, t.P2)
			}
		}

		return children
	}

	cost := func(from, to Point2D) int {
		return from.ManhattanDistance(to)
	}

	costs, _ := Dijkstra(start, children, cost)
	fmt.Println(costs[end])
}

type Segment struct {
	P1, P2 Point2D
}

func InputToSegments() (Point2D, Point2D, []Segment) {
	var segments []Segment

	var turtle Turtle
	for _, part := range in.Split(",") {
		dir, n := part[0], ParseInt(part[1:])
		if dir == 'L' {
			turtle.TurnLeft()
		}
		if dir == 'R' {
			turtle.TurnRight()
		}

		turtle.Forward(1)

		start := turtle.Location
		turtle.Forward(n - 1)
		segments = append(segments, Segment{P1: start, P2: turtle.Location})
	}

	return Origin2D, turtle.Location, segments
}

func Intersects(s, t Segment) bool {
	sx1, sx2 := min(s.P1.X, s.P2.X), max(s.P1.X, s.P2.X)
	sy1, sy2 := min(s.P1.Y, s.P2.Y), max(s.P1.Y, s.P2.Y)
	tx1, tx2 := min(t.P1.X, t.P2.X), max(t.P1.X, t.P2.X)
	ty1, ty2 := min(t.P1.Y, t.P2.Y), max(t.P1.Y, t.P2.Y)

	overlap := func(min1, max1, min2, max2 int) bool {
		return max(min1, min2) <= min(max1, max2)
	}

	return overlap(sx1, sx2, tx1, tx2) && overlap(sy1, sy2, ty1, ty2)
}
