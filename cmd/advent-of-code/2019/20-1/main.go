package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"unicode"
)

func main() {
	grid, portals, start, goal := InputToMaze()

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D

		if other, found := portals[p]; found {
			children = append(children, other)
		}

		for _, child := range p.OrthogonalNeighbors() {
			if grid.GetPoint(child) {
				children = append(children, child)
			}
		}

		return children
	}

	isGoal := func(p lib.Point2D) bool {
		return p == goal
	}

	path, _ := lib.BreadthFirstSearch(start, children, isGoal)
	fmt.Println(len(path) - 1) // the path includes the starting point
}

func InputToMaze() (lib.Grid2D[bool], map[lib.Point2D]lib.Point2D, lib.Point2D, lib.Point2D) {
	lines := lib.InputToLines()
	width := len(lines[2]) + 2
	height := len(lines)

	get := func(x, y int) rune {
		if 0 <= y && y < height && 0 <= x && x < len(lines[y]) {
			return rune(lines[y][x])
		}
		return ' '
	}

	isLetter := unicode.IsLetter

	grid := lib.NewGrid2D[bool](width, height)
	labels := make(map[string][]lib.Point2D)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := get(x, y)
			grid.Set(x, y, c == '.')

			if c1, c2, c3 := c, get(x, y+1), get(x, y+2); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], lib.Point2D{X: x, Y: y + 2})
			}
			if c1, c2, c3 := get(x, y-1), c, get(x, y-2); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], lib.Point2D{X: x, Y: y - 2})
			}
			if c1, c2, c3 := c, get(x+1, y), get(x+2, y); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], lib.Point2D{X: x + 2, Y: y})
			}
			if c1, c2, c3 := get(x-1, y), c, get(x-2, y); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], lib.Point2D{X: x - 2, Y: y})
			}
		}
	}

	var start, goal lib.Point2D
	portals := make(map[lib.Point2D]lib.Point2D)

	for label, ps := range labels {
		switch label {
		case "AA":
			start = ps[0]
		case "ZZ":
			goal = ps[0]
		default:
			portals[ps[0]] = ps[1]
			portals[ps[1]] = ps[0]
		}
	}

	return grid, portals, start, goal
}
