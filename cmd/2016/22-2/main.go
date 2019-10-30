package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	board := InputToBoard(2016, 22)

	// First we need to move the hole to be adjacent to the data.  It's important
	// that we arrive at the side of the hole that's closest to the goal.
	dx := board.goal.X - board.data.X
	adx := dx
	if adx < 0 {
		adx = -adx
	}

	dy := board.goal.Y - board.data.Y
	ady := dy
	if ady < 0 {
		ady = -ady
	}

	var goal1 aoc.Point2D
	if adx > ady {
		// The points are further away in the X direction, so we want to arrive at
		// the East or West of the data
		if dx < 0 {
			goal1 = board.data.West()
		} else {
			goal1 = board.data.East()
		}
	} else {
		// The points are further away in the Y direction, so we want to arrive at
		// the North or South of the data
		if dy < 0 {
			goal1 = board.data.South()
		} else {
			goal1 = board.data.North()
		}
	}

	isGoal1 := func(node aoc.Node) bool {
		board := node.(*Board)
		return board.hole == goal1
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic1 := func(node aoc.Node) int {
		// An estimate of the number of moves would be 1 less move than the
		// manhattan distance between the hole and the data
		board := node.(*Board)
		distance := board.hole.ManhattanDistance(board.data) - 1
		if distance < 0 {
			return 0
		}
		return distance
	}

	path1, distance1, found := aoc.AStarSearch(board, isGoal1, cost, heuristic1)
	if !found {
		log.Fatal("no path to having hole adjacent to data found")
	}

	// Next we need to move the data from it's location to the goal.  We start
	// where the hole is adjacent to the data
	board = path1[len(path1)-1].(*Board)

	isGoal2 := func(node aoc.Node) bool {
		// Check if the data is at the goal.
		board := node.(*Board)
		return board.data == board.goal
	}

	heuristic2 := func(node aoc.Node) int {
		// An estimate of the number of moves to get the data to the goal is the
		// manhattan distance between the two.  In reality the number of moves will
		// be more than that because the hole needs to keep moving around the data
		// to get it closer to the goal.
		board := node.(*Board)

		// We also want to keep the hole adjacent to the data at all times.  We're
		// unable to move the data at all if the hole isn't next to it.  We know
		// that the hole has strayed away from the data if its manhattan distance
		// to the data is greater than 2.
		if board.hole.ManhattanDistance(board.data) > 2 {
			return 1000000
		}

		return board.hole.ManhattanDistance(board.goal) - 1
	}

	path2, distance2, found := aoc.AStarSearch(board, isGoal2, cost, heuristic2)
	if !found {
		log.Fatal("no path to having data at goal found")
	}

	fmt.Printf("len(path1): %d, distance1: %d\n", len(path1), distance1)
	fmt.Printf("len(path2): %d, distance2: %d\n", len(path2), distance2)
	fmt.Printf("total distance: %d\n", distance1+distance2)
}

type Cell string
type Board struct {
	hole aoc.Point2D
	goal aoc.Point2D
	data aoc.Point2D

	maxX, maxY int
	cells      map[aoc.Point2D]Cell
}

const (
	EMPTY Cell = "."
	WALL  Cell = "#"
	GOAL  Cell = "G"
	DATA  Cell = "D"
	HOLE  Cell = "H"
)

func InputToBoard(year, day int) *Board {
	var maxX, maxY int
	for _, node := range InputToNodes(year, day) {
		if node.x > maxX {
			maxX = node.x
		}

		if node.y > maxY {
			maxY = node.y
		}
	}

	cells := make(map[aoc.Point2D]Cell)

	var hole, goal, data aoc.Point2D
	for _, node := range InputToNodes(year, day) {
		var cell = EMPTY
		if node.used == 0 {
			cell = HOLE
			hole = aoc.Point2D{X: node.x, Y: node.y}
		}
		if node.used > 100 {
			cell = WALL
		}
		if node.x == maxX && node.y == 0 {
			cell = DATA
			data = aoc.Point2D{X: node.x, Y: node.y}
		}
		if node.x == 0 && node.y == 0 {
			cell = GOAL
			goal = aoc.Point2D{X: node.x, Y: node.y}
		}

		cells[aoc.Point2D{X: node.x, Y: node.y}] = cell
	}

	return &Board{
		hole: hole,
		goal: goal,
		data: data,

		maxX:  maxX,
		maxY:  maxY,
		cells: cells,
	}
}

func (b *Board) ID() string {
	return fmt.Sprintf("H:%s G:%s D:%s", b.hole, b.goal, b.data)
}

func (b *Board) Children() []aoc.Node {
	cellAt := func(p aoc.Point2D) Cell {
		if cell, ok := b.cells[p]; ok {
			return cell
		}

		return WALL
	}

	holes := []aoc.Point2D{
		b.hole.North(),
		b.hole.South(),
		b.hole.West(),
		b.hole.East(),
	}

	var children []aoc.Node
	for _, hole := range holes {
		existing := cellAt(hole)

		if existing != WALL {
			cells := copy(b.cells)
			cells[b.hole] = cells[hole]
			cells[hole] = HOLE

			// If we've moved the hole into the cell where the data was then we need
			// to update the data pointer in the board.
			data := b.data
			if existing == DATA {
				data = b.hole
			}

			child := &Board{
				hole: hole,
				goal: b.goal,
				data: data,

				maxX:  b.maxX,
				maxY:  b.maxY,
				cells: cells,
			}
			children = append(children, child)
		}
	}

	return children
}

func (b *Board) String() string {
	var builder strings.Builder

	for y := 0; y <= b.maxY; y++ {
		for x := 0; x <= b.maxX; x++ {
			builder.WriteString(string(b.cells[aoc.Point2D{X: x, Y: y}]))
			builder.WriteString(" ")
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

type Node struct {
	x, y                  int
	size, used, available int
}

func InputToNodes(year, day int) []Node {
	var regex = regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+(\d+)%`)

	var nodes []Node
	for _, line := range aoc.InputToLines(year, day)[2:] {
		matches := regex.FindStringSubmatch(line)

		x := aoc.ParseInt(matches[1])
		y := aoc.ParseInt(matches[2])
		size := aoc.ParseInt(matches[3])
		used := aoc.ParseInt(matches[4])
		avail := size - used

		nodes = append(nodes, Node{x, y, size, used, avail})
	}

	return nodes
}

func copy(cells map[aoc.Point2D]Cell) map[aoc.Point2D]Cell {
	copy := make(map[aoc.Point2D]Cell)
	for k, v := range cells {
		copy[k] = v
	}

	return copy
}
