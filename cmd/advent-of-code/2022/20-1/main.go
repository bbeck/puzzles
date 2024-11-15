package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	nodes, zero := InputToNodes()

	N := len(nodes)
	for _, n := range nodes {
		// First remove it from the list.
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev

		// Advance until we're at the insertion location.
		steps := n.Value % (N - 1)
		if steps < 0 {
			steps += N - 1
		}
		for steps > 0 {
			n.Next = n.Next.Next
			steps--
		}

		// Insert it back into its new location.
		n.Prev = n.Next.Prev
		n.Prev.Next = n
		n.Next.Prev = n
	}

	var sum int
	for i := 0; i < 3; i++ {
		steps := 1000 % N
		for steps > 0 {
			zero = zero.Next
			steps--
		}

		sum += zero.Value
	}
	fmt.Println(sum)
}

type Node struct {
	Value      int
	Prev, Next *Node
}

func InputToNodes() ([]*Node, *Node) {
	var nodes []*Node
	var zero *Node
	for _, n := range puz.InputToInts() {
		node := &Node{Value: n}
		nodes = append(nodes, node)

		if n == 0 {
			zero = node
		}
	}

	N := len(nodes)
	for i, n := range nodes {
		n.Prev = nodes[(i-1+N)%N]
		n.Next = nodes[(i+1+N)%N]
	}

	return nodes, zero
}
