package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"github.com/bbeck/advent-of-code/lib/cpus"
)

func main() {
	open, goal := Explore()

	// Perform an exhaustive breadth first traversal from the goal, keeping track
	// along the way of the minimum distance from the goal to the current point.
	// In addition we'll keep track of the longest distance recorded as well.
	distances := make(map[lib.Point2D]int)
	longest := 0

	children := func(p lib.Point2D) []lib.Point2D {
		var children []lib.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
				distances[child] = distances[p] + 1
				longest = lib.Max(longest, distances[child])

				// No reason to revisit this child in the future.
				open.Remove(child)
			}
		}
		return children
	}

	isGoal := func(p lib.Point2D) bool {
		return false
	}

	lib.BreadthFirstSearch(goal, children, isGoal)
	fmt.Println(longest)
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
