package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	paths := InputToWirePaths(2019, 3)
	steps0 := paths[0].PointSteps()
	steps1 := paths[1].PointSteps()

	var distances []int
	for p, dist0 := range steps0 {
		if dist1, ok := steps1[p]; ok {
			distances = append(distances, dist0+dist1)
		}
	}

	sort.Ints(distances)
	fmt.Printf("shortest distance: %d\n", distances[0])
}

type WirePath struct {
	dirs    []string
	lengths []int
}

func (p WirePath) PointSteps() map[aoc.Point2D]int {
	steps := make(map[aoc.Point2D]int)

	var location aoc.Point2D
	var distance int
	for step := 0; step < len(p.dirs); step++ {
		dir := p.dirs[step]
		length := p.lengths[step]

		for n := 0; n < length; n++ {
			switch dir {
			case "U":
				location = location.Up()
			case "D":
				location = location.Down()
			case "L":
				location = location.Left()
			case "R":
				location = location.Right()
			}

			distance++
			steps[location] = distance
		}
	}

	return steps
}

func InputToWirePaths(year, day int) []WirePath {
	var paths []WirePath

	for _, line := range aoc.InputToLines(year, day) {
		var dirs []string
		var lengths []int

		for _, part := range strings.Split(line, ",") {
			dirs = append(dirs, string(part[0]))
			lengths = append(lengths, aoc.ParseInt(part[1:]))
		}

		paths = append(paths, WirePath{dirs: dirs, lengths: lengths})
	}

	return paths
}
