package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	open, goal := Explore()

	children := func(p Point2D) []Point2D {
		var children []Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
			}
		}
		return children
	}

	isGoal := func(p Point2D) bool {
		return p == goal
	}

	if path, ok := BreadthFirstSearch(Origin2D, children, isGoal); ok {
		fmt.Println(len(path) - 1) // The path includes the starting point.
	}
}

var Headings = []Heading{Up, Down, Left, Right}
var Opposite = map[Heading]Heading{
	Up:    Down,
	Down:  Up,
	Left:  Right,
	Right: Left,
}

func Explore() (Set[Point2D], Point2D) {
	var open Set[Point2D]
	var goal Point2D

	robot := NewRobot()
	current := Origin2D

	var helper func()
	helper = func() {
		for _, heading := range Headings {
			status := robot.Move(heading)
			if status == 0 {
				continue
			}

			current = current.Move(heading)
			if status == 2 {
				goal = current
			}

			if open.Add(current) {
				helper()
			}

			robot.Move(Opposite[heading])
			current = current.Move(Opposite[heading])
		}
	}
	helper()

	return open, goal
}

type Robot struct {
	CPU      cpus.IntcodeCPU
	Commands chan int
	Status   chan int
}

func NewRobot() *Robot {
	commands := make(chan int)
	status := make(chan int)

	robot := &Robot{
		CPU: cpus.IntcodeCPU{
			Memory: cpus.InputToIntcodeMemory(),
			Input:  func() int { return <-commands },
			Output: func(value int) { status <- value },
		},
		Commands: commands,
		Status:   status,
	}
	go robot.CPU.Execute()
	return robot
}

func (r *Robot) Move(h Heading) int {
	mapping := map[Heading]int{
		Up:    1,
		Down:  2,
		Left:  3,
		Right: 4,
	}

	r.Commands <- mapping[h]
	return <-r.Status
}
