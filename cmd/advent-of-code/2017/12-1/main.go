package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var s DisjointSet[int]
	for node, children := range InputToGraph() {
		for _, child := range children {
			s.UnionWithAdd(node, child)
		}
	}

	fmt.Println(s.Size(0))
}

func InputToGraph() map[int][]int {
	edges := make(map[int][]int)

	in.LinesToS(func(in in.Scanner[any]) any {
		edges[in.Int()] = in.Ints()
		return nil
	})

	return edges
}
