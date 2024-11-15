package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	paths := InputToPaths()
	s0, s1 := paths[0].Steps(), paths[1].Steps()

	best := math.MaxInt
	for p, stepsP := range s0 {
		if stepsQ, found := s1[p]; found {
			best = puz.Min(best, stepsP+stepsQ)
		}
	}
	fmt.Println(best)
}

type Path struct {
	Dirs    []string
	Lengths []int
}

func (p Path) Steps() map[puz.Point2D]int {
	steps := make(map[puz.Point2D]int)

	var current puz.Point2D
	var count int
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
			count++

			if _, found := steps[current]; !found {
				steps[current] = count
			}
		}
	}

	return steps
}

func InputToPaths() []Path {
	return puz.InputLinesTo(func(line string) Path {
		var path Path
		for _, part := range strings.Split(line, ",") {
			path.Dirs = append(path.Dirs, string(part[0]))
			path.Lengths = append(path.Lengths, puz.ParseInt(part[1:]))
		}

		return path
	})
}
