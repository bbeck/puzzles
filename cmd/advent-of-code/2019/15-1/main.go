package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	open, goal := Explore()

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
			}
		}
		return children
	}

	isGoal := func(p lib.Point2D) bool {
		return p == goal
	}

	if path, ok := lib.BreadthFirstSearch(lib.Origin2D, children, isGoal); ok {
		fmt.Println(len(path) - 1) // The path includes the starting point.
	}
}

var Headings = []lib.Heading{lib.Up, lib.Down, lib.Left, lib.Right}
var Reverse = map[lib.Heading]lib.Heading{
	lib.Up:    lib.Down,
	lib.Down:  lib.Up,
	lib.Left:  lib.Right,
	lib.Right: lib.Left,
}

func Explore() (lib.Set[lib.Point2D], lib.Point2D) {
	var open lib.Set[lib.Point2D]
	var goal lib.Point2D

	robot := NewRobot()
	current := lib.Origin2D

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

			robot.Move(Reverse[heading])
			current = current.Move(Reverse[heading])
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

func (r *Robot) Move(h lib.Heading) int {
	mapping := map[lib.Heading]int{
		lib.Up:    1,
		lib.Down:  2,
		lib.Left:  3,
		lib.Right: 4,
	}

	r.Commands <- mapping[h]
	return <-r.Status
}