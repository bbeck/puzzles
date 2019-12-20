package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

type Cells map[aoc.Point2D]bool
type Portals map[aoc.Point2D]aoc.Point2D
type Depths map[aoc.Point2D]int

var cells Cells
var portals Portals
var depths Depths
var start, goal aoc.Point2D

func main() {
	cells, portals, depths, start, goal = InputToMaze(2019, 20)

	isGoal := func(n aoc.Node) bool {
		l := n.(Location)
		return l.depth == 0 && l.Point2D == goal
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(n aoc.Node) int {
		return 1 // TODO: might need to make this smarter
	}

	_, distance, found := aoc.AStarSearch(Location{start, 0}, isGoal, cost, heuristic)
	if !found {
		log.Fatal("no path found")
	}

	fmt.Println("shortest path:", distance)
}

type Location struct {
	aoc.Point2D
	depth int
}

func (l Location) ID() string {
	return fmt.Sprintf("%s@%d", l.String(), l.depth)
}

func (l Location) Children() []aoc.Node {
	var children []aoc.Node
	if cells[l.Up()] {
		children = append(children, Location{l.Up(), l.depth})
	}
	if cells[l.Right()] {
		children = append(children, Location{l.Right(), l.depth})
	}
	if cells[l.Down()] {
		children = append(children, Location{l.Down(), l.depth})
	}
	if cells[l.Left()] {
		children = append(children, Location{l.Left(), l.depth})
	}
	if other, ok := portals[l.Point2D]; ok {
		depth := l.depth + depths[l.Point2D]
		if depth >= 0 {
			children = append(children, Location{other, depth})
		}
	}

	return children
}

func InputToMaze(year, day int) (Cells, Portals, Depths, aoc.Point2D, aoc.Point2D) {
	isLetter := func(r rune) bool {
		return 'A' <= r && r <= 'Z'
	}

	chars := make(map[aoc.Point2D]rune)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			p := aoc.Point2D{X: x, Y: y}
			chars[p] = c
		}
	}

	// First pass through to map all of the cells
	cells := make(Cells)
	for p, c := range chars {
		cells[p] = c == '.'
	}

	// Now make a second pass through looking for all of the letters to map
	// the start and goal as well as portals.
	var start, goal aoc.Point2D
	var labels = make(map[string][]aoc.Point2D)
	for p, c := range chars {
		var label string
		var target aoc.Point2D

		if isLetter(c) && isLetter(chars[p.Right()]) {
			label = fmt.Sprintf("%c%c", c, chars[p.Right()])
			if cells[p.Right().Right()] {
				target = p.Right().Right()
			} else {
				target = p.Left()
			}
		}

		if isLetter(c) && isLetter(chars[p.Down()]) {
			label = fmt.Sprintf("%c%c", c, chars[p.Down()])
			if cells[p.Down().Down()] {
				target = p.Down().Down()
			} else {
				target = p.Up()
			}
		}

		switch label {
		case "":
		case "AA":
			start = target
		case "ZZ":
			goal = target
		default:
			labels[label] = append(labels[label], target)
		}
	}

	portals := make(Portals)
	depths := make(Depths)
	for label, ps := range labels {
		if len(ps) != 2 {
			log.Fatalf("incorrect number of targets for label: %s, %+v", label, ps)
		}

		p1 := ps[0]
		p2 := ps[1]

		portals[p1] = p2
		portals[p2] = p1

		// Determine which point is on the inner vs outer ring.  Points on the
		// outside will become out of bounds 3 steps away.
		if chars[p1.Up().Up().Up()] == 0 ||
			chars[p1.Right().Right().Right()] == 0 ||
			chars[p1.Down().Down().Down()] == 0 ||
			chars[p1.Left().Left().Left()] == 0 {
			depths[p1] = -1
			depths[p2] = 1
		} else {
			depths[p1] = 1
			depths[p2] = -1
		}
	}

	return cells, portals, depths, start, goal
}
