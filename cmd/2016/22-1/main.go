package main

import (
	"fmt"
	"regexp"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	nodes := InputToNodes(2016, 22)

	var count int
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			a := nodes[i]
			b := nodes[j]

			if (a.used > 0) && (a.x != b.x || a.y != b.y) && (a.used < b.available) {
				count++
			}
		}
	}

	fmt.Printf("number of compatible nodes: %d\n", count)
}

type Node struct {
	x, y                  int
	size, used, available int
}

func InputToNodes(year, day int) []Node {
	var regex = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)

	var nodes []Node
	for _, line := range aoc.InputToLines(year, day)[2:] {
		matches := regex.FindStringSubmatch(line)

		x := aoc.ParseInt(matches[1])
		y := aoc.ParseInt(matches[2])
		size := aoc.ParseInt(matches[3])
		used := aoc.ParseInt(matches[4])
		avail := size - used

		nodes = append(nodes, Node{x, y, size, used, avail})
	}

	return nodes
}
