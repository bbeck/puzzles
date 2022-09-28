package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	root := InputToTree()

	var sum int
	Walk(root, func(n Node) {
		sum += aoc.Sum(n.Metadata...)
	})
	fmt.Println(sum)
}

func Walk(n Node, visit func(Node)) {
	visit(n)
	for _, child := range n.Children {
		Walk(child, visit)
	}
}

type Node struct {
	Children []Node
	Metadata []int
}

func InputToTree() Node {
	var ns aoc.Deque[int]
	for _, s := range strings.Fields(aoc.InputToString(2018, 8)) {
		ns.PushBack(aoc.ParseInt(s))
	}

	var next func() Node
	next = func() Node {
		numChildren := ns.PopFront()
		numMetadata := ns.PopFront()

		var n Node
		for i := 0; i < numChildren; i++ {
			n.Children = append(n.Children, next())
		}
		for i := 0; i < numMetadata; i++ {
			n.Metadata = append(n.Metadata, ns.PopFront())
		}

		return n
	}

	return next()
}
