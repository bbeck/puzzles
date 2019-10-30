package main

import (
	"fmt"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	facility := InputToFacility(2018, 20)

	distances := map[aoc.Point2D]int{
		aoc.Point2D{X: 0, Y: 0}: 0,
	}

	aoc.BreadthFirstSearch(facility[aoc.Point2D{X: 0, Y: 0}], func(node aoc.Node) bool {
		room := node.(*Room)
		if _, ok := distances[room.location]; ok {
			return false
		}

		// This room's distance is the distance of its closest child + 1
		closest := math.MaxInt64
		for _, child := range room.Children() {
			cRoom := child.(*Room)
			if cDistance, ok := distances[cRoom.location]; ok && cDistance < closest {
				closest = cDistance
			}
		}

		distances[room.location] = closest + 1
		return false
	})

	var count int
	for _, distance := range distances {
		if distance >= 1000 {
			count++
		}
	}

	fmt.Printf("paths > 1000 doors: %d\n", count)
}

type Room struct {
	location   aoc.Point2D
	n, s, w, e bool

	facility *Facility
}

func (r *Room) ID() string {
	return r.location.String()
}

func (r *Room) Children() []aoc.Node {
	facility := r.facility

	var children []aoc.Node

	if r.n {
		children = append(children, (*facility)[r.location.Up()])
	}
	if r.s {
		children = append(children, (*facility)[r.location.Down()])
	}
	if r.w {
		children = append(children, (*facility)[r.location.Left()])
	}
	if r.e {
		children = append(children, (*facility)[r.location.Right()])
	}

	return children
}

type Facility map[aoc.Point2D]*Room

func InputToFacility(year, day int) Facility {
	regex := aoc.InputToString(year, day)

	// Start off assuming we're standing at (0, 0)
	location := aoc.Point2D{X: 0, Y: 0}

	facility := make(Facility)

	getRoom := func(p aoc.Point2D) *Room {
		room := facility[p]
		if room == nil {
			room = &Room{
				location: location,
				facility: &facility,
			}
			facility[p] = room
		}

		return room
	}

	stack := aoc.NewStack()
	for _, c := range regex {
		switch c {
		case 'N':
			getRoom(location).n = true
			location = location.Up()
			getRoom(location).s = true

		case 'S':
			getRoom(location).s = true
			location = location.Down()
			getRoom(location).n = true

		case 'W':
			getRoom(location).w = true
			location = location.Left()
			getRoom(location).e = true

		case 'E':
			getRoom(location).e = true
			location = location.Right()
			getRoom(location).w = true

		case '|':
			location = stack.Pop().(aoc.Point2D)
			stack.Push(location)

		case '(':
			stack.Push(location)

		case ')':
			location = stack.Pop().(aoc.Point2D)
		}
	}

	return facility
}
