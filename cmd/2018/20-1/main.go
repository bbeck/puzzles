package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	facility := InputToFacility(2018, 20)

	origin := aoc.Point2D{X: 0, Y: 0}
	player := Path{end: origin, facility: facility}

	var furthest aoc.Point2D
	var furthestDistance int
	aoc.BreadthFirstSearch(player, func(node aoc.Node) bool {
		path := node.(Path)
		if path.length > furthestDistance {
			furthestDistance = path.length
			furthest = path.end
		}

		// We want to explore every path
		return false
	})

	// We want to know how many doors we need to go through.  Each step we take
	// moves 2 steps in a direction -- the first step is to get the doorway, the
	// second step is to enter the room.  So half of the steps on the longest path
	// are through doors.
	fmt.Printf("furthest point: %s, num doors: %d\n", furthest, furthestDistance/2)
}

type Facility map[aoc.Point2D]bool

func (f Facility) String() string {
	var minX, maxX int
	var minY, maxY int

	for p := range f {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	var builder strings.Builder
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			p := aoc.Point2D{X: x, Y: y}
			if p.X == 0 && p.Y == 0 {
				builder.WriteString("X")
			} else if f[p] {
				builder.WriteString(".")
			} else {
				builder.WriteString("#")
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func InputToFacility(year, day int) Facility {
	regex := aoc.InputToString(year, day)

	// Start off assuming we're standing at (0, 0)
	location := aoc.Point2D{X: 0, Y: 0}

	facility := make(Facility)
	facility[location] = true

	stack := aoc.NewStack()
	for _, c := range regex {
		switch c {
		case '^':
			// nothing to do

		case '$':
			// nothing to do

		case 'N':
			location = location.Up()
			facility[location] = true // doorway
			location = location.Up()
			facility[location] = true // room

		case 'S':
			location = location.Down()
			facility[location] = true // doorway
			location = location.Down()
			facility[location] = true // room

		case 'W':
			location = location.Left()
			facility[location] = true // doorway
			location = location.Left()
			facility[location] = true // room

		case 'E':
			location = location.Right()
			facility[location] = true // doorway
			location = location.Right()
			facility[location] = true // room

		case '|':
			location = stack.Pop().(aoc.Point2D)
			stack.Push(location)

		case '(':
			stack.Push(location)

		case ')':
			location = stack.Pop().(aoc.Point2D)

		default:
			log.Fatalf("unrecognized character in regex: %v", c)
		}
	}

	return facility
}

type Path struct {
	end      aoc.Point2D
	facility Facility

	length int
}

func (p Path) ID() string {
	return p.end.String()
}

func (p Path) Children() []aoc.Node {
	locations := []aoc.Point2D{
		p.end.Up(),
		p.end.Down(),
		p.end.Left(),
		p.end.Right(),
	}

	var children []aoc.Node
	for _, location := range locations {
		if p.facility[location] {
			children = append(children, Path{
				end:      location,
				facility: p.facility,
				length:   p.length + 1,
			})
		}
	}

	return children
}
