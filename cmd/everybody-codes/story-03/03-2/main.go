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
	var insert func(root, n *Node) bool
	insert = func(root, n *Node) bool {
		if root.Left == nil && (root.LeftColor == n.PlugColor || root.LeftShape == n.PlugShape) {
			root.Left = n
			return true
		}
		if root.Left != nil && insert(root.Left, n) {
			return true
		}
		if root.Right == nil && (root.RightColor == n.PlugColor || root.RightShape == n.PlugShape) {
			root.Right = n
			return true
		}
		if root.Right != nil && insert(root.Right, n) {
			return true
		}
		return false
	}

	root, nodes := &nodes[0], nodes[1:]
	for _, node := range nodes {
		insert(root, &node)
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
