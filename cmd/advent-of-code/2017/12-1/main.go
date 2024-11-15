package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var s puz.DisjointSet[int]
	for node, children := range InputToGraph() {
		for _, child := range children {
			s.UnionWithAdd(node, child)
		}
	}

	fmt.Println(s.Size(0))
}

func InputToGraph() map[int][]int {
	edges := make(map[int][]int)
	for _, line := range puz.InputToLines(2017, 12) {
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, "<-> ", "")
		fields := strings.Fields(line)

		from := puz.ParseInt(fields[0])
		for _, s := range fields[1:] {
			to := puz.ParseInt(s)
			edges[from] = append(edges[from], to)
		}
	}

	return edges
}
