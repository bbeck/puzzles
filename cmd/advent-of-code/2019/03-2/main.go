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
	s0, s1 := paths[0].Steps(), paths[1].Steps()

	best := math.MaxInt
	for p, stepsP := range s0 {
		if stepsQ, found := s1[p]; found {
			best = Min(best, stepsP+stepsQ)
		}
	}
	fmt.Println(best)
}

type Path struct {
	Dirs    []string
	Lengths []int
}

func (p Path) Steps() map[Point2D]int {
	steps := make(map[Point2D]int)

	var current Point2D
	var count int
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
			count++

			if _, found := steps[current]; !found {
				steps[current] = count
			}
		}
	}

	return steps
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
