package main

import (
	"fmt"
	"math"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	paths := InputToPaths()
	ps0, ps1 := paths[0].Points(), paths[1].Points()

	best := math.MaxInt
	for p := range ps0 {
		if ps1.Contains(p) {
			best = Min(best, Origin2D.ManhattanDistance(p))
		}
	}
	fmt.Println(best)
}

type Path struct {
	Dirs    []string
	Lengths []int
}

func (p Path) Points() Set[Point2D] {
	var ps Set[Point2D]

	var current Point2D
	for i := range p.Dirs {
		for range p.Lengths[i] {
			switch p.Dirs[i] {
			case "U":
				current = current.Up()
			case "D":
				current = current.Down()
			case "L":
				current = current.Left()
			case "R":
				current = current.Right()
			}

			ps.Add(current)
		}
	}

	return ps
}

func InputToPaths() []Path {
	return in.LinesTo(func(line string) Path {
		var path Path
		for part := range strings.SplitSeq(line, ",") {
			path.Dirs = append(path.Dirs, string(part[0]))
			path.Lengths = append(path.Lengths, ParseInt(part[1:]))
		}
		return path
	})
}
