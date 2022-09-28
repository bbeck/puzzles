package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var value func(Node) int
	value = func(n Node) int {
		if len(n.Children) == 0 {
			return aoc.Sum(n.Metadata...)
		}

		var sum int
		for _, index := range n.Metadata {
			if index <= 0 || index > len(n.Children) {
				continue
			}

			sum += value(n.Children[index-1])
		}
		return sum
	}

	root := InputToTree()
	fmt.Println(value(root))
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
