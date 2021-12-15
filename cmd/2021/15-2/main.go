package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	m := InputToCave()

	var maxX, maxY int
	for p := range m {
		maxX = aoc.MaxInt(maxX, p.X)
		maxY = aoc.MaxInt(maxY, p.Y)
	}

	start := Position{Point2D: aoc.Point2D{X: 0, Y: 0}, m: m}
	goal := aoc.Point2D{X: maxX, Y: maxY}

	visit := func(node aoc.Node) bool {
		return node.(Position).Point2D == goal
	}
	cost := func(from, to aoc.Node) int {
		b := to.(Position)
		cost := b.m[b.Point2D]

		return cost
	}

	heuristic := func(node aoc.Node) int {
		return 1
	}

	_, distance, _ := aoc.AStarSearch(start, visit, cost, heuristic)
	fmt.Println(distance)
}

type Position struct {
	aoc.Point2D
	m map[aoc.Point2D]int
}

func (p Position) ID() string {
	return p.String()
}

func (p Position) Children() []aoc.Node {
	candidates := []aoc.Point2D{p.Up(), p.Right(), p.Down(), p.Left()}

	var children []aoc.Node
	for _, c := range candidates {
		if _, ok := p.m[c]; ok {
			children = append(children, Position{
				Point2D: c,
				m:       p.m,
			})
		}
	}
	return children
}

func InputToCave() map[aoc.Point2D]int {
	m := make(map[aoc.Point2D]int)

	var width, height int
	for y, line := range aoc.InputToLines(2021, 15) {
		width = y + 1
		for x, c := range line {
			m[aoc.Point2D{X: x, Y: y}] = aoc.ParseInt(string(c))
			height = x + 1
		}
	}

	var get func(x, y int) int
	get = func(x, y int) int {
		p := aoc.Point2D{X: x, Y: y}
		if x < width && y < height {
			return m[p]
		}

		var n int
		for x >= width {
			x -= width
			n++
		}
		for y >= height {
			y -= height
			n++
		}

		v := m[aoc.Point2D{X: x, Y: y}] + n
		if v > 9 {
			v -= 9
		}
		return v
	}

	for y := 0; y < 5*height; y++ {
		for x := 0; x < 5*width; x++ {
			m[aoc.Point2D{X: x, Y: y}] = get(x, y)
		}
	}

	return m
}
