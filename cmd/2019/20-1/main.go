package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

type Cells map[aoc.Point2D]bool
type Portals map[aoc.Point2D]aoc.Point2D

var cells Cells
var portals Portals
var start, goal aoc.Point2D

func main() {
	cells, portals, start, goal = InputToMaze(2019, 20)

	isGoal := func(n aoc.Node) bool {
		return n.(Location).Point2D == goal
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(n aoc.Node) int {
		return 1 // TODO: might need to make this smarter
	}

	_, distance, found := aoc.AStarSearch(Location{start}, isGoal, cost, heuristic)
	if !found {
		log.Fatal("no path found")
	}

	fmt.Println("shortest path:", distance)
}

type Location struct {
	aoc.Point2D
}

func (l Location) ID() string {
	return l.String()
}

func (l Location) Children() []aoc.Node {
	var children []aoc.Node
	if cells[l.Up()] {
		children = append(children, Location{l.Up()})
	}
	if cells[l.Right()] {
		children = append(children, Location{l.Right()})
	}
	if cells[l.Down()] {
		children = append(children, Location{l.Down()})
	}
	if cells[l.Left()] {
		children = append(children, Location{l.Left()})
	}
	if other, ok := portals[l.Point2D]; ok {
		children = append(children, Location{other})
	}

	return children
}

func InputToMaze(year, day int) (Cells, Portals, aoc.Point2D, aoc.Point2D) {
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
	for label, ps := range labels {
		if len(ps) != 2 {
			log.Fatalf("incorrect number of targets for label: %s, %+v", label, ps)
		}

		portals[ps[0]] = ps[1]
		portals[ps[1]] = ps[0]
	}

	return cells, portals, start, goal
}
