package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	nodes := InputToNodes()

	var count int
	for i := range nodes {
		for j := range nodes {
			if i == j || nodes[i].Used == 0 || nodes[i].Used > nodes[j].Avail {
				continue
			}

			count++
		}
	}

	fmt.Println(count)
}

type Node struct {
	Point2D
	Size, Used, Avail int
}

func InputToNodes() []Node {
	// Drop the header lines
	in.Line()
	in.Line()

	return in.LinesToS(func(in in.Scanner[Node]) Node {
		return Node{
			Point2D: Point2D{X: in.Int(), Y: in.Int()},
			Size:    in.Int(),
			Used:    in.Int(),
			Avail:   in.Int(),
		}
	})
}
