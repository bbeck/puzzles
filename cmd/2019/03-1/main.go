package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := make(map[aoc.Point2D]int)
	intersections := make([]aoc.Point2D, 0)

	for id, path := range InputToWirePaths(2019, 3) {
		location := aoc.Point2D{}
		for step := 0; step < len(path.dirs); step++ {
			dir := path.dirs[step]
			length := path.lengths[step]

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

				if grid[location] != 0 && grid[location] != id+1 {
					intersections = append(intersections, location)
				}
				grid[location] = id + 1
			}
		}
	}

	origin := aoc.Point2D{}
	sort.Slice(intersections, func(i, j int) bool {
		a := intersections[i]
		b := intersections[j]
		return a.ManhattanDistance(origin) < b.ManhattanDistance(origin)
	})

	fmt.Printf("closest intersection: %+v, distance: %d\n",
		intersections[0], intersections[0].ManhattanDistance(origin))
}

type WirePath struct {
	dirs    []string
	lengths []int
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
