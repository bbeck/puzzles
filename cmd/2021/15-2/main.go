package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	cave, width, height := InputToCave()

	start := State{position: aoc.Point2D{}, cave: cave}
	goal := aoc.Point2D{X: width - 1, Y: height - 1}

	isGoal := func(node aoc.Node) bool {
		return node.(State).position == goal
	}

	cost := func(from, to aoc.Node) int {
		state := to.(State)
		return cave[state.position]
	}

	heuristic := func(node aoc.Node) int {
		return node.(State).position.ManhattanDistance(goal)
	}

	_, distance, found := aoc.AStarSearch(start, isGoal, cost, heuristic)
	if found {
		fmt.Println(distance)
	}
}

type State struct {
	position aoc.Point2D
	cave     map[aoc.Point2D]int
}

func (s State) ID() string {
	return s.position.String()
}

func (s State) Children() []aoc.Node {
	var children []aoc.Node
	for _, neighbor := range s.position.OrthogonalNeighbors() {
		if _, ok := s.cave[neighbor]; !ok {
			continue
		}

		children = append(children, State{
			position: neighbor,
			cave:     s.cave,
		})
	}

	return children
}

func InputToCave() (map[aoc.Point2D]int, int, int) {
	var cave = make(map[aoc.Point2D]int)
	var width, height int

	// First, read the cave definition from the input file.
	for y, line := range aoc.InputToLines(2021, 15) {
		for x, c := range line {
			cave[aoc.Point2D{X: x, Y: y}] = aoc.ParseInt(string(c))
			width = x + 1
		}
		height = y + 1
	}

	// Now, tile the existing cave definition 5 times in each
	// direction, incrementing the values by 1 each step.
	for y := 0; y < 5*height; y++ {
		for x := 0; x < 5*width; x++ {
			by := y % height
			bx := x % width
			offset := y/height + x/width

			value := cave[aoc.Point2D{X: bx, Y: by}] + offset
			for value > 9 {
				value -= 9
			}

			cave[aoc.Point2D{X: x, Y: y}] = value
		}
	}

	return cave, 5 * width, 5 * height
}
