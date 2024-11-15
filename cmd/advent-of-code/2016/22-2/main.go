package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	var width, height int
	var hole puz.Point2D
	for _, node := range InputToNodes() {
		width = puz.Max(width, node.X)
		height = puz.Max(height, node.Y)

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
	puz.Point2D
	Size, Used, Avail int
}

func InputToNodes() []Node {
	var nodes []Node
	for _, line := range puz.InputToLines(2016, 22) {
		if !strings.HasPrefix(line, "/dev/grid") {
			continue
		}

		line = strings.ReplaceAll(line, "/dev/grid/node-", "")
		line = strings.ReplaceAll(line, "x", "")
		line = strings.ReplaceAll(line, "y", "")
		line = strings.ReplaceAll(line, "T", "")
		line = strings.ReplaceAll(line, "-", " ")
		fields := strings.Fields(line)

		nodes = append(nodes, Node{
			Point2D: puz.Point2D{X: puz.ParseInt(fields[0]), Y: puz.ParseInt(fields[1])},
			Size:    puz.ParseInt(fields[2]),
			Used:    puz.ParseInt(fields[3]),
			Avail:   puz.ParseInt(fields[4]),
		})
	}

	return nodes
}
