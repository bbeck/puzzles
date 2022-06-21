package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	g := InputToGraph()

	var s aoc.DisjointSet[int]
	for node, children := range g {
		for _, child := range children {
			s.UnionWithAdd(node, child)
		}
	}

	var roots aoc.Set[int]
	for node := range g {
		if root, found := s.Find(node); found {
			roots.Add(root)
		}
	}
	fmt.Println(len(roots))
}

func InputToGraph() map[int][]int {
	edges := make(map[int][]int)
	for _, line := range aoc.InputToLines(2017, 12) {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "<-> ", "")
		fields := strings.Fields(line)

		from := aoc.ParseInt(fields[0])
		for _, s := range fields[1:] {
			to := aoc.ParseInt(s)
			edges[from] = append(edges[from], to)
		}
	}

	return edges
}
