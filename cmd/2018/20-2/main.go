package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	regex := aoc.InputToString(2018, 20)
	world, origin := ParseRegex(regex)

	// Determine the point that's furthest from the current location.  Because
	// we're looking for the furthest point we'll assume there are no cycles.
	// This means a basic breadth first search will work.
	distances := map[aoc.Point2D]int{
		origin: 0,
	}

	children := func(p aoc.Point2D) []aoc.Point2D {
		var children []aoc.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if !world.InBoundsPoint(child) || !world.GetPoint(child) {
				continue
			}

			if _, found := distances[child]; !found {
				distances[child] = distances[p] + 1
				children = append(children, child)
			}
		}

		return children
	}

	aoc.BreadthFirstSearch(origin, children, func(p aoc.Point2D) bool {
		return false
	})

	var count int
	for p, d := range distances {
		// Anything that's not at the same X/Y parity level as the origin is a door.
		if p.X%2 != origin.X%2 || p.Y%2 != origin.Y%2 {
			continue
		}

		// Count paths longer than 1000 units (double to take into account doors).
		if d >= 2000 {
			count++
		}
	}
	fmt.Println(count)
}

func ParseRegex(input string) (aoc.Grid2D[bool], aoc.Point2D) {
	input = strings.ReplaceAll(input, "^", "")
	input = strings.ReplaceAll(input, "$", "")

	// Which coordinates are open in the world.
	var open aoc.Set[aoc.Point2D]
	open.Add(aoc.Origin2D)

	// The current positions we're exploring.  This is the fringe of the search.
	// Every time we encounter a new group we'll push this set onto a stack so
	// that we can return to them whenever we need to branch within the group.
	// When we exit a group we'll pop from the stack.
	var currents aoc.Set[aoc.Point2D]
	currents.Add(aoc.Origin2D)

	// The stack of previous positions on the fringe that we've encountered.
	// These are kept so that when we're in an OR group we know which position
	// to return to when we encounter an OR operator.
	var previous aoc.Stack[aoc.Set[aoc.Point2D]]

	// A stack of positions that have been reached while in a group.  Whenever
	// we enter a group we'll push an empty set onto the stack.  As we process
	// the different OR expressions within the group we'll union the set of
	// positions reached with the top entry of this stack.  When we exit the
	// group the top of this stack will become the current positions set.
	var group aoc.Stack[aoc.Set[aoc.Point2D]]

	for _, ch := range input {
		var next aoc.Set[aoc.Point2D]

		switch ch {
		case 'N':
			for _, p := range currents.Entries() {
				open.Add(p.Up(), p.Up().Up())
				next.Add(p.Up().Up())
			}
		case 'E':
			for _, p := range currents.Entries() {
				open.Add(p.Right(), p.Right().Right())
				next.Add(p.Right().Right())
			}
		case 'S':
			for _, p := range currents.Entries() {
				open.Add(p.Down(), p.Down().Down())
				next.Add(p.Down().Down())
			}

		case 'W':
			for _, p := range currents.Entries() {
				open.Add(p.Left(), p.Left().Left())
				next.Add(p.Left().Left())
			}

		case '(':
			previous.Push(currents)
			group.Push(aoc.Set[aoc.Point2D]{})

		case '|':
			// We're going to backtrack.  Before we do remember where we ended up in
			// the group stack.
			g := group.Pop()
			g.Add(currents.Entries()...)
			group.Push(g)

			currents = previous.Peek()

		case ')':
			// We've finished this group.  Union together all the ending positions
			// for each term of the group.
			currents.Add(group.Pop().Entries()...)

			previous.Pop()
		}

		if len(next) > 0 {
			currents = next
		}
	}

	// Convert the set of open points into a grid.
	tl, br := aoc.GetBounds(open.Entries())
	grid := aoc.NewGrid2D[bool](br.X-tl.X+1, br.Y-tl.Y+1)
	for _, p := range open.Entries() {
		grid.Add(p.X-tl.X, p.Y-tl.Y, true)
	}

	return grid, aoc.Point2D{X: -tl.X, Y: -tl.Y}
}
