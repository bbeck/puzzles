package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"unicode"
)

func main() {
	grid, portals, depths, start, goal := InputToMaze()
	start3d := puz.Point3D{X: start.X, Y: start.Y, Z: 0}
	goal3d := puz.Point3D{X: goal.X, Y: goal.Y, Z: 0}

	children := func(p3d puz.Point3D) []puz.Point3D {
		p := puz.Point2D{X: p3d.X, Y: p3d.Y}

		var children []puz.Point3D
		if other, found := portals[p]; found {
			if depth := p3d.Z + depths[p]; depth >= 0 {
				children = append(children, puz.Point3D{X: other.X, Y: other.Y, Z: p3d.Z + depths[p]})
			}
		}

		for _, child := range p.OrthogonalNeighbors() {
			if grid.GetPoint(child) {
				children = append(children, puz.Point3D{X: child.X, Y: child.Y, Z: p3d.Z})
			}
		}

		return children
	}

	isGoal := func(p puz.Point3D) bool {
		return p == goal3d
	}

	path, _ := puz.BreadthFirstSearch(start3d, children, isGoal)
	fmt.Println(len(path) - 1) // the path includes the starting point
}

func InputToMaze() (puz.Grid2D[bool], map[puz.Point2D]puz.Point2D, map[puz.Point2D]int, puz.Point2D, puz.Point2D) {
	lines := puz.InputToLines(2019, 20)
	width := len(lines[2]) + 2
	height := len(lines)

	get := func(x, y int) rune {
		if 0 <= y && y < height && 0 <= x && x < len(lines[y]) {
			return rune(lines[y][x])
		}
		return ' '
	}

	isLetter := unicode.IsLetter

	grid := puz.NewGrid2D[bool](width, height)
	labels := make(map[string][]puz.Point2D)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := get(x, y)
			grid.Set(x, y, c == '.')

			if c1, c2, c3 := c, get(x, y+1), get(x, y+2); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], puz.Point2D{X: x, Y: y + 2})
			}
			if c1, c2, c3 := get(x, y-1), c, get(x, y-2); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], puz.Point2D{X: x, Y: y - 2})
			}
			if c1, c2, c3 := c, get(x+1, y), get(x+2, y); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], puz.Point2D{X: x + 2, Y: y})
			}
			if c1, c2, c3 := get(x-1, y), c, get(x-2, y); isLetter(c1) && isLetter(c2) && c3 == '.' {
				label := string(c1) + string(c2)
				labels[label] = append(labels[label], puz.Point2D{X: x - 2, Y: y})
			}
		}
	}

	var start, goal puz.Point2D
	portals := make(map[puz.Point2D]puz.Point2D)
	depths := make(map[puz.Point2D]int)

	for label, ps := range labels {
		switch label {
		case "AA":
			start = ps[0]
		case "ZZ":
			goal = ps[0]
		default:
			portals[ps[0]] = ps[1]
			portals[ps[1]] = ps[0]

			for _, p := range ps {
				if p.X == 2 || p.X == width-3 || p.Y == 2 || p.Y == height-3 {
					depths[p] = -1
				} else {
					depths[p] = 1
				}
			}
		}
	}

	return grid, portals, depths, start, goal
}
