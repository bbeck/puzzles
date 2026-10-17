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
		if root.Left == nil && root.LeftSocket == n.Plug {
			root.Left = n
			return true
		}
		if root.Left != nil && insert(root.Left, n) {
			return true
		}
		if root.Right == nil && root.RightSocket == n.Plug {
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
	ID                      int
	Plug                    string
	LeftSocket, RightSocket string
	Data                    string
	Left, Right             *Node
}

func InputToNodes() []Node {
	in.Remove("id=", "plug=", "leftSocket=", "rightSocket=", "data=")

	return in.LinesToS(func(in in.Scanner[Node]) Node {
		fields := in.Split(",")
		return Node{
			ID:          ParseInt(fields[0]),
			Plug:        strings.TrimSpace(fields[1]),
			LeftSocket:  strings.TrimSpace(fields[2]),
			RightSocket: strings.TrimSpace(fields[3]),
			Data:        strings.TrimSpace(fields[4]),
		}
	})
}
