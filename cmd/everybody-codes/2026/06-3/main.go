package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var tree = InputToTree()

	do := func(start string) Set[string] {
		var seen Set[string]
		BreadthFirstSearch(
			start,
			func(s string) []string { return tree[s] },
			func(s string) bool { seen.Add(s); return false },
		)
		return seen.DifferenceElems(start)
	}

	v1, v2, v3 := do("NULLv1"), do("NULLv2"), do("NULLv3")
	all := SetFrom[string]().
		Union(v1.Intersect(v2)).
		Union(v1.Intersect(v3)).
		Union(v2.Intersect(v3))
	fmt.Println(len(all))
}

func InputToTree() map[string][]string {
	var nodes = make(map[string][]string)
	for in.HasNext() {
		lhs, rhs := in.Cut(":")
		for r := range strings.SplitSeq(rhs, ",") {
			nodes[r] = append(nodes[r], lhs)
		}
	}
	return nodes
}
