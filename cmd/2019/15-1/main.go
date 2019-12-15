package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/bbeck/advent-of-code/aoc"
)

const (
	_ int = iota
	Empty
	Wall
	Goal
)

var grid map[aoc.Point2D]int

func main() {
	grid = Explore(2_000_000)

	var goal aoc.Point2D
	for k, v := range grid {
		if v == Goal {
			goal = k
		}
	}

	start := Location{}

	isGoal := func(node aoc.Node) bool {
		p := node.(Location).Point2D
		return grid[p] == Goal
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(node aoc.Node) int {
		p := node.(Location).Point2D
		return p.ManhattanDistance(goal)
	}

	_, distance, found := aoc.AStarSearch(start, isGoal, cost, heuristic)
	if !found {
		fmt.Println("no path found")
		return
	}

	fmt.Printf("shortest path: %d\n", distance)
}

func Explore(steps int) map[aoc.Point2D]int {
	grid := make(map[aoc.Point2D]int)

	// The currently location we're at as well as the next location we'll be at
	// if our move succeeds.
	var location, next aoc.Point2D

	cpu := &CPU{
		memory: InputToMemory(2019, 15),
	}

	cpu.input = func(int) int {
		dir := rand.Intn(4) + 1
		switch dir {
		case 1:
			next = location.Up()
		case 2:
			next = location.Down()
		case 3:
			next = location.Left()
		case 4:
			next = location.Right()
		default:
			log.Fatalf("invalid direction chosen: %d", dir)
		}

		return dir
	}

	cpu.output = func(value int) {
		switch value {
		case 0:
			grid[next] = Wall
		case 1:
			grid[next] = Empty
			location = next
		case 2:
			grid[next] = Goal
			location = next
		}

		// If we've captured the amount of data that we've wanted to we can
		// terminate the program and return our data.
		steps--
		if steps <= 0 {
			cpu.Stop()
		}
	}

	cpu.Execute()
	return grid
}

type Location struct {
	aoc.Point2D
}

func (l Location) ID() string {
	return l.String()
}

func (l Location) Children() []aoc.Node {
	var children []aoc.Node
	if grid[l.Up()] != Wall {
		children = append(children, Location{l.Up()})
	}
	if grid[l.Down()] != Wall {
		children = append(children, Location{l.Down()})
	}
	if grid[l.Left()] != Wall {
		children = append(children, Location{l.Left()})
	}
	if grid[l.Right()] != Wall {
		children = append(children, Location{l.Right()})
	}
	return children
}
