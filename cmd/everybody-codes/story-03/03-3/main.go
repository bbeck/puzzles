package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	root := BuildTree(InputToNodes())
	fmt.Println(Checksum(root))
}

func BuildTree(nodes []Node) *Node {
	// Insert a node into the tree and if the insert displaces a node continue
	// the search to see if the displaced node fits anywhere.  If not, then start
	// a new search with the displaced node from the root of the tree.
	var insert func(root, n *Node) *Node
	insert = func(root, n *Node) *Node {
		if root.Left == nil && (root.LeftColor == n.PlugColor || root.LeftShape == n.PlugShape) {
			// There's nothing here, we can bond without displacing anything.
			root.Left = n
			return nil
		}

		// If we insert the node and displace something, we'll want to skip the
		// recursion into the thing we just inserted.
		var skip bool

		if root.Left != nil && (root.LeftColor != root.Left.PlugColor || root.LeftShape != root.Left.PlugShape) {
			// There's something here, but we can displace it because it's not a
			// strong bond.
			if root.LeftColor == n.PlugColor && root.LeftShape == n.PlugShape {
				n, root.Left = root.Left, n
				skip = true
			}
		}

		if !skip && root.Left != nil {
			n = insert(root.Left, n)
			if n == nil {
				return nil
			}
		}
		skip = false

		if root.Right == nil && (root.RightColor == n.PlugColor || root.RightShape == n.PlugShape) {
			// There's nothing here, we can bond without displacing anything.
			root.Right = n
			return nil
		}

		if root.Right != nil && (root.RightColor != root.Right.PlugColor || root.RightShape != root.Right.PlugShape) {
			// There's something here, but we can displace it because it's not a
			// strong bond.
			if root.RightColor == n.PlugColor && root.RightShape == n.PlugShape {
				n, root.Right = root.Right, n
				skip = true
			}
		}

		if !skip && root.Right != nil {
			n = insert(root.Right, n)
			if n == nil {
				return nil
			}
		}

		return n
	}

	root, nodes := &nodes[0], nodes[1:]
	for i := range nodes {
		node := &nodes[i]
		for node != nil {
			node = insert(root, node)
		}
	}
	return root
}

func Checksum(root *Node) int {
	var checksum, count int

	// Perform an inorder traversal, computing the checksum as we go.
	var visit func(*Node)
	visit = func(root *Node) {
		if root.Left != nil {
			visit(root.Left)
		}
		count++
		checksum += count * root.ID
		if root.Right != nil {
			visit(root.Right)
		}
	}
	visit(root)

	return checksum
}

type Node struct {
	ID                     int
	PlugColor, PlugShape   string
	LeftColor, LeftShape   string
	RightColor, RightShape string
	Data                   string
	Left, Right            *Node
}

func InputToNodes() []Node {
	in.Remove("id=", "plug=", "leftSocket=", "rightSocket=", "data=")

	return in.LinesToS(func(in in.Scanner[Node]) Node {
		fields := in.Split(",")
		plugColor, plugShape, _ := strings.Cut(strings.TrimSpace(fields[1]), " ")
		leftColor, leftShape, _ := strings.Cut(strings.TrimSpace(fields[2]), " ")
		rightColor, rightShape, _ := strings.Cut(strings.TrimSpace(fields[3]), " ")

		return Node{
			ID:         ParseInt(fields[0]),
			PlugColor:  plugColor,
			PlugShape:  plugShape,
			LeftColor:  leftColor,
			LeftShape:  leftShape,
			RightColor: rightColor,
			RightShape: rightShape,
			Data:       strings.TrimSpace(fields[4]),
		}
	})
}
