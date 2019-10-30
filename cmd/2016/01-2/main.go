package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	start := aoc.Point2D{}

	current := start
	heading := "N"
	seen := map[aoc.Point2D]bool{
		current: true,
	}

outer:
	for _, direction := range InputToDirections(2016, 1) {
		switch heading {
		case "N":
			if direction.turn == "L" {
				heading = "W"
			} else {
				heading = "E"
			}
		case "S":
			if direction.turn == "L" {
				heading = "E"
			} else {
				heading = "W"
			}
		case "W":
			if direction.turn == "L" {
				heading = "S"
			} else {
				heading = "N"
			}
		case "E":
			if direction.turn == "L" {
				heading = "N"
			} else {
				heading = "S"
			}
		}

		for n := 0; n < direction.steps; n++ {
			switch heading {
			case "N":
				current = current.North()
			case "S":
				current = current.South()
			case "W":
				current = current.West()
			case "E":
				current = current.East()
			}

			if seen[current] {
				break outer
			}
			seen[current] = true
		}
	}

	dx := current.X
	if dx < 0 {
		dx *= -1
	}

	dy := current.Y
	if dy < 0 {
		dy *= -1
	}

	fmt.Printf("distance: %d\n", dx+dy)
}

type Direction struct {
	turn  string
	steps int
}

func InputToDirections(year, day int) []Direction {
	directions := make([]Direction, 0)

	s := aoc.InputToString(year, day)
	for _, part := range strings.Split(strings.ReplaceAll(s, ",", ""), " ") {
		directions = append(directions, Direction{part[0:1], aoc.ParseInt(part[1:])})
	}

	return directions
}
