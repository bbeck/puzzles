package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"github.com/bbeck/advent-of-code/puz/cpus"
)

func main() {
	open, goal := Explore()

	// Perform an exhaustive breadth first traversal from the goal, keeping track
	// along the way of the minimum distance from the goal to the current point.
	// In addition we'll keep track of the longest distance recorded as well.
	distances := make(map[puz.Point2D]int)
	longest := 0

	children := func(p puz.Point2D) []puz.Point2D {
		var children []puz.Point2D
		for _, child := range p.OrthogonalNeighbors() {
			if open.Contains(child) {
				children = append(children, child)
				distances[child] = distances[p] + 1
				longest = puz.Max(longest, distances[child])

				// No reason to revisit this child in the future.
				open.Remove(child)
			}
		}
		return children
	}

	isGoal := func(p puz.Point2D) bool {
		return false
	}

	puz.BreadthFirstSearch(goal, children, isGoal)
	fmt.Println(longest)
}

var Headings = []puz.Heading{puz.Up, puz.Down, puz.Left, puz.Right}
var Reverse = map[puz.Heading]puz.Heading{
	puz.Up:    puz.Down,
	puz.Down:  puz.Up,
	puz.Left:  puz.Right,
	puz.Right: puz.Left,
}

func Explore() (puz.Set[puz.Point2D], puz.Point2D) {
	var open puz.Set[puz.Point2D]
	var goal puz.Point2D

	robot := NewRobot()
	current := puz.Origin2D

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

func (r *Robot) Move(h puz.Heading) int {
	mapping := map[puz.Heading]int{
		puz.Up:    1,
		puz.Down:  2,
		puz.Left:  3,
		puz.Right: 4,
	}

	r.Commands <- mapping[h]
	return <-r.Status
}
