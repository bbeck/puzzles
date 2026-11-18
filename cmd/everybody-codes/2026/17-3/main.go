package main

import (
	"fmt"
	"math"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	grid := in.ToStringGrid2D()

	var solid Set[Point2D]
	grid.ForEachPoint(func(p Point2D, s string) {
		if s == "#" {
			solid.Add(p)
		}
	})

	// Assume we're traveling along the left side of a cell P in the downward
	// direction.  There end up being 3 cases we need to consider involving the
	// cell below and the cell below and to the left.
	//
	//       │              │              │
	//       ▼ P=#          ▼ P=#          ▼ P=#
	//       └────          │         ─────┘
	//  BL=?   B=.     BL=. │ B=#     BL=#   B=#
	//
	// The other directions follow from this.  The important parts are the forward
	// direction, the cell along the direction as well as the cell diagonal to the
	// current cell.  With those all the decisions can be made.
	children := func(s Pose2D) []Pose2D {
		p, heading := s.Point, s.Heading

		// We don't know which direction we're traveling, so we'll test both.
		var children []Pose2D
		for _, forward := range []Heading{HeadingTurnLeft[heading], HeadingTurnRight[heading]} {
			ahead := p.Move(forward)
			diagonal := ahead.Move(heading)

			if !solid.Contains(ahead) {
				children = append(children, Pose2D{Point: p, Heading: forward})
			}
			if solid.Contains(ahead) && !solid.Contains(diagonal) {
				children = append(children, Pose2D{Point: ahead, Heading: heading})
			}
			if solid.Contains(ahead) && solid.Contains(diagonal) {
				children = append(children, Pose2D{Point: diagonal, Heading: HeadingTurnAround[forward]})
			}
		}

		return children
	}

	goal := func(s Pose2D) bool {
		return s.Point.X == grid.Width-1 && s.Point.Y == grid.Height-1 &&
			(s.Heading == Right || s.Heading == Down)
	}

	var best = math.MaxInt64
	for _, side := range []Heading{Up, Left} {
		steps, _ := BreadthFirstSearch(Pose2D{Point: Origin2D, Heading: side}, children, goal)
		best = Min(best, len(steps))
	}
	fmt.Println(best)
}
