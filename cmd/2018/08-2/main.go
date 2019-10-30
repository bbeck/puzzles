package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	root := InputToTree(2018, 8)
	fmt.Printf("value of root: %d\n", root.Value())
}

type Queue []int

func (q *Queue) Pop() int {
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

type Node struct {
	metadata []int
	children []*Node
}

func (n *Node) Value() int {
	numChildren := len(n.children)

	var sum int
	for _, idx := range n.metadata {
		switch {
		case numChildren == 0:
			sum += idx

		case idx <= 0 || idx > numChildren:
			sum += 0

		default:
			sum += n.children[idx-1].Value()
		}
	}

	return sum
}

func InputToTree(year, day int) *Node {
	q := Queue(InputToInts(year, day))

	var toNode func() *Node
	toNode = func() *Node {
		numChildren := q.Pop()
		numMetadata := q.Pop()

		var children []*Node
		for i := 0; i < numChildren; i++ {
			children = append(children, toNode())
		}

		var metadata []int
		for i := 0; i < numMetadata; i++ {
			metadata = append(metadata, q.Pop())
		}

		return &Node{
			metadata: metadata,
			children: children,
		}
	}

	return toNode()
}

func InputToInts(year, day int) []int {
	parts := strings.Split(aoc.InputToString(year, day), " ")

	var ints []int
	for _, p := range parts {
		ints = append(ints, aoc.ParseInt(p))
	}

	return ints
}
