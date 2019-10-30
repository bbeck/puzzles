package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid(2017, 19)

	loc := FindStart(grid)
	dir := "D"

	step := func(dir string) aoc.Point2D {
		switch dir {
		case "U":
			return loc.Up()
		case "D":
			return loc.Down()
		case "L":
			return loc.Left()
		case "R":
			return loc.Right()
		}

		log.Fatalf("unrecognized direction: %s", dir)
		return aoc.Point2D{}
	}

	// given 3 directions, pick the first one that works and return the new
	// location and direction.
	valid := func(loc aoc.Point2D) bool {
		cell := grid[loc]
		return cell != "" && cell != " "
	}

	choose := func(dir1, dir2, dir3 string) (aoc.Point2D, string) {
		if p := step(dir1); valid(p) {
			return p, dir1
		}

		if p := step(dir2); valid(p) {
			return p, dir2
		}

		if p := step(dir3); valid(p) {
			return p, dir3
		}

		return aoc.Point2D{}, ""
	}

	var steps int
	for done := false; !done; {
		switch {
		case !valid(loc):
			done = true

		case grid[loc] == "|" || grid[loc] == "-":
			loc = step(dir)
			steps++

		case grid[loc] == "+":
			if dir == "U" {
				loc, dir = choose("U", "L", "R")
			} else if dir == "D" {
				loc, dir = choose("D", "L", "R")
			} else if dir == "L" {
				loc, dir = choose("L", "U", "D")
			} else if dir == "R" {
				loc, dir = choose("R", "U", "D")
			}
			steps++

		default:
			loc = step(dir)
			steps++
		}
	}

	fmt.Printf("number of steps: %d\n", steps)
}

func FindStart(grid Grid) aoc.Point2D {
	for x := 0; ; x++ {
		p := aoc.Point2D{X: x, Y: 0}
		if grid[p] == "|" {
			return p
		}
	}
}

type Grid map[aoc.Point2D]string

func InputToGrid(year, day int) Grid {
	grid := make(Grid)
	for y, line := range aoc.InputToLines(year, day) {
		for x, c := range line {
			if c == ' ' {
				continue
			}

			p := aoc.Point2D{X: x, Y: y}
			grid[p] = string(c)
		}
	}

	return grid
}
