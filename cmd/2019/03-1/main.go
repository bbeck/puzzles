package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	paths := InputToPaths()
	ps0, ps1 := paths[0].Points(), paths[1].Points()

	best := math.MaxInt
	for p := range ps0 {
		if ps1.Contains(p) {
			best = aoc.Min(best, aoc.Origin2D.ManhattanDistance(p))
		}
	}
	fmt.Println(best)
}

type Path struct {
	Dirs    []string
	Lengths []int
}

func (p Path) Points() aoc.Set[aoc.Point2D] {
	var ps aoc.Set[aoc.Point2D]

	var current aoc.Point2D
	for i := 0; i < len(p.Dirs); i++ {
		for n := 0; n < p.Lengths[i]; n++ {
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
	return aoc.InputLinesTo(2019, 3, func(line string) Path {
		var path Path
		for _, part := range strings.Split(line, ",") {
			path.Dirs = append(path.Dirs, string(part[0]))
			path.Lengths = append(path.Lengths, aoc.ParseInt(part[1:]))
		}

		return path
	})
}
