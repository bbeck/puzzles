package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	board, waypoints := InputToBoard(2016, 24)
	start := board.location
	N := len(waypoints)

	// We'll perform the Floyd-Warshall algorithm to determine the all-pairs
	// shortest path between any two waypoints on the board.
	dist := make([][]int, N)
	for i := 0; i < N; i++ {
		dist[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i != j {
				dist[i][j] = distance(board, waypoints[i], waypoints[j])
			}
		}
	}

	for k := 1; k < N; k++ {
		for i := 1; i < N; i++ {
			for j := 1; j < N; j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// We now know the minimum distance between any two waypoints on the board.
	// Now we just need to find the permutation through the waypoints that yields
	// the shortest path.
	best := math.MaxInt64
	bestPerm := make([]aoc.Point2D, 0)
	aoc.EnumeratePermutations(N, func(perm []int) {
		// Don't bother checking this permutation if it doesn't start at the correct
		// location.
		if waypoints[perm[0]] != start {
			return
		}

		var cost int
		for i := 1; i < N; i++ {
			cost += dist[perm[i-1]][perm[i]]
		}
		cost += dist[perm[N-1]][perm[0]]

		if cost < best {
			best = cost
			bestPerm = make([]aoc.Point2D, N)
			for i := 0; i < N; i++ {
				bestPerm[i] = waypoints[perm[i]]
			}
		}
	})

	fmt.Printf("best: %d\n", best)
	fmt.Printf("order:\n")
	for i := 0; i < N; i++ {
		fmt.Printf("  %s\n", bestPerm[i])
	}
}

func distance(board *Board, start, end aoc.Point2D) int {
	board.location = start

	isGoal := func(node aoc.Node) bool {
		board := node.(*Board)
		return board.location == end
	}

	cost := func(from, to aoc.Node) int {
		return 1
	}

	heuristic := func(node aoc.Node) int {
		board := node.(*Board)
		return board.location.ManhattanDistance(end)
	}

	_, distance, found := aoc.AStarSearch(board, isGoal, cost, heuristic)
	if !found {
		log.Fatalf("unable to find path between: %s and %s", start, end)
	}

	return distance
}

type Board struct {
	location aoc.Point2D
	width    int
	height   int
	cells    map[aoc.Point2D]string
}

func (b *Board) ID() string {
	return b.location.String()
}

func (b *Board) Children() []aoc.Node {
	locations := []aoc.Point2D{
		b.location.Up(),
		b.location.Down(),
		b.location.Left(),
		b.location.Right(),
	}

	var children []aoc.Node
	for _, nloc := range locations {
		if cell, ok := b.cells[nloc]; ok && cell != WALL {
			child := &Board{
				location: nloc,
				width:    b.width,
				height:   b.height,
				cells:    b.cells,
			}
			children = append(children, child)
		}
	}

	return children
}

func (b *Board) String() string {
	var builder strings.Builder
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			p := aoc.Point2D{X: x, Y: y}

			if p == b.location {
				builder.WriteString("P")
			} else {
				builder.WriteString(b.cells[p])
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

const (
	EMPTY string = "."
	WALL  string = "#"
)

func InputToBoard(year, day int) (*Board, []aoc.Point2D) {
	var width, height int       // the dimensions of the board
	var location aoc.Point2D    // my current location
	var waypoints []aoc.Point2D // points to visit in some order

	cells := make(map[aoc.Point2D]string)
	for y, line := range aoc.InputToLines(year, day) {
		height = y + 1

		for x, c := range line {
			width = x + 1
			p := aoc.Point2D{X: x, Y: y}

			if c == '0' {
				location = p
			}

			switch c {
			case '.':
				cells[p] = EMPTY
			case '#':
				cells[p] = WALL
			default:
				cells[p] = EMPTY
				waypoints = append(waypoints, p)
			}
		}
	}

	return &Board{
		width:    width,
		height:   height,
		location: location,
		cells:    cells,
	}, waypoints
}
