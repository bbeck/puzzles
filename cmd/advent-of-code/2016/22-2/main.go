package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var width, height int
	var hole Point2D
	for _, node := range InputToNodes() {
		width = Max(width, node.X)
		height = Max(height, node.Y)

		if node.Used == 0 {
			hole = node.Point2D
		}
	}

	// This solution relies on the layout of the puzzle input.  It's possible
	// that this won't work on every input.
	//
	// Visualizing the nodes yields a picture that looks like this:
	//
	// T........G
	// ..........
	// ..........
	// ..........
	// .#########
	// .....H....
	// ..........
	// ..........
	// ..........
	// ..........
	//
	// T = The target node to transfer the data into
	// G = The node that contains the goal data
	// H = The node that is empty of data
	// # = A node that has more data than fits on the empty node
	//
	// This can be viewed as a type of sliding block puzzle where the goal is to
	// via data transfers move the empty node (hole) next to the node containing
	// the goal data and then to use the hole to move the goal data closer to the
	// target.
	//
	// First, move the hole to the left of goal.  This requires
	// H.x + H.y + G.x - 1 moves.
	//
	// Next, we use the hole to move the goal data over to column 0.  Moving the
	// goal to the left one space requires 1 move, and resetting the hole to the
	// left of the goal requires 4 moves.
	//
	// ....HG => ....GH => ....G. => ....G. => ....G. => ...HG.
	// ......    ......    .....H    ....H.    ...H..    ......
	//
	// We need to move the goal to the left G.x times, and need to reset the hole
	// G.x - 1 times.

	var steps int

	// Move the hole to the left of the goal
	steps += hole.X + hole.Y + width - 1

	// Move the goal to the first column
	steps += width + 4*(width-1)

	fmt.Println(steps)
}

type Node struct {
	Point2D
	Size, Used, Avail int
}

func InputToNodes() []Node {
	// Drop the header lines
	in.Line()
	in.Line()

	return in.LinesToS(func(in in.Scanner[Node]) Node {
		return Node{
			Point2D: Point2D{X: in.Int(), Y: in.Int()},
			Size:    in.Int(),
			Used:    in.Int(),
			Avail:   in.Int(),
		}
	})
}
