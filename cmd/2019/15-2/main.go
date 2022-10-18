package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"github.com/bbeck/advent-of-code/aoc/cpus"
)

func main() {
	open, goal := Explore()

	// Perform an exhaustive breadth first traversal from the goal, keeping track
	// along the way of the minimum distance from the goal to the current point.
	// In addition we'll keep track of the longest distance recorded as well.
	distances := make(map[aoc.Point2D]int)
	longest := 0

	children := func(p aoc.Point2D) []aoc.Point2D {
		var children []aoc.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
				distances[child] = distances[p] + 1
				longest = aoc.Max(longest, distances[child])

				// No reason to revisit this child in the future.
				open.Remove(child)
			}
		}
		return children
	}

	isGoal := func(p aoc.Point2D) bool {
		return false
	}

	aoc.BreadthFirstSearch(goal, children, isGoal)
	fmt.Println(longest)
}

var Headings = []aoc.Heading{aoc.Up, aoc.Down, aoc.Left, aoc.Right}
var Reverse = map[aoc.Heading]aoc.Heading{
	aoc.Up:    aoc.Down,
	aoc.Down:  aoc.Up,
	aoc.Left:  aoc.Right,
	aoc.Right: aoc.Left,
}

func Explore() (aoc.Set[aoc.Point2D], aoc.Point2D) {
	var open aoc.Set[aoc.Point2D]
	var goal aoc.Point2D

	robot := NewRobot()
	current := aoc.Origin2D

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
			Memory: cpus.InputToIntcodeMemory(2019, 15),
			Input:  func() int { return <-commands },
			Output: func(value int) { status <- value },
		},
		Commands: commands,
		Status:   status,
	}
	go robot.CPU.Execute()
	return robot
}

func (r *Robot) Move(h aoc.Heading) int {
	mapping := map[aoc.Heading]int{
		aoc.Up:    1,
		aoc.Down:  2,
		aoc.Left:  3,
		aoc.Right: 4,
	}

	r.Commands <- mapping[h]
	return <-r.Status
}
