package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var location Cell
	var furthest int
	for _, step := range InputToSteps(2017, 11) {
		switch step {
		case "nw":
			location = location.NorthWest()
		case "n":
			location = location.North()
		case "ne":
			location = location.NorthEast()
		case "sw":
			location = location.SouthWest()
		case "s":
			location = location.South()
		case "se":
			location = location.SouthEast()
		}

		d := Distance(Cell{}, location)
		if d > furthest {
			furthest = d
		}
	}

	fmt.Printf("furthest: %d\n", furthest)
}

type Cell struct {
	x, y, z int
}

func (c Cell) NorthWest() Cell {
	return Cell{x: c.x - 1, y: c.y + 1, z: c.z + 0}
}

func (c Cell) North() Cell {
	return Cell{x: c.x + 0, y: c.y + 1, z: c.z - 1}
}

func (c Cell) NorthEast() Cell {
	return Cell{x: c.x + 1, y: c.y + 0, z: c.z - 1}
}

func (c Cell) SouthWest() Cell {
	return Cell{x: c.x - 1, y: c.y + 0, z: c.z + 1}
}

func (c Cell) South() Cell {
	return Cell{x: c.x + 0, y: c.y - 1, z: c.z + 1}
}

func (c Cell) SouthEast() Cell {
	return Cell{x: c.x + 1, y: c.y - 1, z: c.z + 0}
}

func (c Cell) ID() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}

func (c Cell) Children() map[aoc.Node]int {
	return map[aoc.Node]int{
		c.NorthWest(): 1,
		c.North():     1,
		c.NorthEast(): 1,
		c.SouthWest(): 1,
		c.South():     1,
		c.SouthEast(): 1,
	}
}

func Distance(start, end Cell) int {
	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	return (abs(end.x-start.x) + abs(end.y-start.y) + abs(end.z-start.z)) / 2
}

func InputToSteps(year, day int) []string {
	s := aoc.InputToString(year, day)
	return strings.Split(s, ",")
}
