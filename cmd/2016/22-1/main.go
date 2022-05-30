package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	nodes := InputToNodes()

	var count int
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < len(nodes); j++ {
			if i == j || nodes[i].Used == 0 || nodes[i].Used > nodes[j].Avail {
				continue
			}

			count++
		}
	}

	fmt.Println(count)
}

type Node struct {
	aoc.Point2D
	Size, Used, Avail int
}

func InputToNodes() []Node {
	var nodes []Node
	for _, line := range aoc.InputToLines(2016, 22) {
		if !strings.HasPrefix(line, "/dev/grid") {
			continue
		}

		line = strings.ReplaceAll(line, "/dev/grid/node-", "")
		line = strings.ReplaceAll(line, "x", "")
		line = strings.ReplaceAll(line, "y", "")
		line = strings.ReplaceAll(line, "T", "")
		line = strings.ReplaceAll(line, "-", " ")
		fields := strings.Fields(line)

		nodes = append(nodes, Node{
			Point2D: aoc.Point2D{X: aoc.ParseInt(fields[0]), Y: aoc.ParseInt(fields[1])},
			Size:    aoc.ParseInt(fields[2]),
			Used:    aoc.ParseInt(fields[3]),
			Avail:   aoc.ParseInt(fields[4]),
		})
	}

	return nodes
}
