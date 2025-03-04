package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	g := InputToGraph()

	var s DisjointSet[int]
	for node, children := range g {
		for _, child := range children {
			s.UnionWithAdd(node, child)
		}
	}

	var roots Set[int]
	for node := range g {
		if root, found := s.Find(node); found {
			roots.Add(root)
		}
	}
	fmt.Println(len(roots))
}

func InputToGraph() map[int][]int {
	edges := make(map[int][]int)

	in.LinesToS(func(in in.Scanner[any]) any {
		edges[in.Int()] = in.Ints()
		return nil
	})

	return edges
}
