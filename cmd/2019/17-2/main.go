package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Grid map[aoc.Point2D]bool

type Step struct {
	turn    string
	forward int
}

func main() {
	grid, robot := ReadGrid()
	path := ComputePath(robot, grid)
	main, A, B, C := ComputeSubprocesses(path)
	if main == "" {
		log.Fatal("unable to find a solution")
	}

	toInputs := func(s string) []int {
		var nums []int
		for _, c := range s {
			nums = append(nums, int(c))
		}

		return nums
	}

	// Convert the subprocesses into the stream of inputs
	var inputs []int
	inputs = append(inputs, toInputs(main)...)
	inputs = append(inputs, '\n')
	inputs = append(inputs, toInputs(A)...)
	inputs = append(inputs, '\n')
	inputs = append(inputs, toInputs(B)...)
	inputs = append(inputs, '\n')
	inputs = append(inputs, toInputs(C)...)
	inputs = append(inputs, '\n')
	inputs = append(inputs, 'n') // do we want a live feed?
	inputs = append(inputs, '\n')

	memory := InputToMemory(2019, 17)
	memory[0] = 2

	var output int
	cpu := &CPU{
		memory: memory,
		input: func(addr int) int {
			i := inputs[0]
			inputs = inputs[1:]
			return i
		},
		output: func(value int) {
			output = value
		},
	}
	cpu.Execute()

	fmt.Println("dust collected:", output)
}

func IsSolution(steps string) bool {
	steps = strings.ReplaceAll(steps, "A,", "")
	steps = strings.ReplaceAll(steps, "B,", "")
	steps = strings.ReplaceAll(steps, "C,", "")
	return len(steps) == 0
}

func ComputeSubprocesses(path []Step) (string, string, string, string) {
	// Convert the path to a string so that we can use the strings library on it.
	// We'll make sure every step is followed by a comma so we can do clean
	// substitutions.
	steps := ""
	for _, step := range path {
		steps += fmt.Sprintf("%s,%d,", step.turn, step.forward)
	}

	// Now that we have the path through the maze, we need to collapse all of it
	// into subroutines A, B, and C.  Each subroutine cannot be longer than 20
	// characters when represented as ASCII text.  Because we need to consume all
	// of the input we can focus on taking exclusively from the front of the path.
	//
	// Because the dataset is relatively small and we know we need 3 subroutines
	// we can brute force the sizes of subroutine A, B, and C.  The largest a
	// subroutine can be is 6 steps because when converted to ASCII each step
	// requires a minimum of 3 characters, but some will require more if there are
	// two digit distances.
	for _, a := range Subprocesses(steps) {
		stepsA := strings.ReplaceAll(steps, a+",", "A,")

		for _, b := range Subprocesses(stepsA) {
			stepsB := strings.ReplaceAll(stepsA, b+",", "B,")

			for _, c := range Subprocesses(stepsB) {
				stepsC := strings.ReplaceAll(stepsB, c+",", "C,")

				if IsSolution(stepsC) {
					return stepsC[0 : len(stepsC)-1], a, b, c
				}
			}
		}
	}

	return "", "", "", ""
}

func Subprocesses(steps string) []string {
	for len(steps) > 0 {
		if steps[0] == 'L' || steps[0] == 'R' {
			break
		}
		steps = steps[1:]
	}

	var subprocesses []string
	for i := 0; i < 20; i++ {
		// Subprocesses cannot call each other, so if we run into a call to another
		// subprocess then we're done and can't move any further to the right in the
		// string.
		if steps[i] == 'A' || steps[i] == 'B' || steps[i] == 'C' {
			break
		}

		// Look for the end of a step
		if steps[i] != ',' {
			continue
		}

		if steps[i-1] < '0' || steps[i-1] > '9' {
			continue
		}

		subprocesses = append(subprocesses, steps[0:i])
	}

	return subprocesses
}

func ComputePath(robot Turtle, grid Grid) []Step {
	// Pick a direction to turn (L or R) or return the empty string for a dead end
	choose := func(robot Turtle, grid Grid) string {
		left := Turtle{robot.location, robot.direction}
		left.Left()
		left.Forward()
		if grid[left.location] {
			return "L"
		}

		right := Turtle{robot.location, robot.direction}
		right.Right()
		right.Forward()
		if grid[right.location] {
			return "R"
		}

		return ""
	}

	// Determine if the robot can take a step forward without falling.
	canStep := func(robot Turtle, grid Grid) bool {
		t := Turtle{robot.location, robot.direction}
		t.Forward()
		return grid[t.location]
	}

	var path []Step
	for {
		var count int
		turn := choose(robot, grid)
		if turn == "L" {
			robot.Left()
		} else if turn == "R" {
			robot.Right()
		} else {
			// There was no turn, we're done.
			break
		}

		for canStep(robot, grid) {
			count++
			robot.Forward()
		}

		path = append(path, Step{turn, count})
	}

	return path
}

func ReadGrid() (Grid, Turtle) {
	grid := make(Grid)
	var robot Turtle

	current := aoc.Point2D{}

	output := func(value int) {
		switch value {
		case '.':
			current = current.Right()

		case '#':
			grid[current] = true
			current = current.Right()

		case '^':
			grid[current] = true
			robot.location = current
			robot.direction = "N"
			current = current.Right()

		case '>':
			grid[current] = true
			robot.location = current
			robot.direction = "E"
			current = current.Right()

		case 'v':
			grid[current] = true
			robot.location = current
			robot.direction = "S"
			current = current.Right()

		case '<':
			grid[current] = true
			robot.location = current
			robot.direction = "W"
			current = current.Right()

		case '\n':
			current = aoc.Point2D{0, current.Y + 1}
		}
	}

	cpu := &CPU{
		memory: InputToMemory(2019, 17),
		output: output,
	}
	cpu.Execute()

	return grid, robot
}

type Turtle struct {
	location  aoc.Point2D
	direction string
}

func (t *Turtle) Forward() {
	switch t.direction {
	case "N":
		t.location = t.location.Up()
	case "E":
		t.location = t.location.Right()
	case "S":
		t.location = t.location.Down()
	case "W":
		t.location = t.location.Left()
	}
}

func (t *Turtle) Left() {
	switch t.direction {
	case "N":
		t.direction = "W"
	case "E":
		t.direction = "N"
	case "S":
		t.direction = "E"
	case "W":
		t.direction = "S"
	}
}

func (t *Turtle) Right() {
	switch t.direction {
	case "N":
		t.direction = "E"
	case "E":
		t.direction = "S"
	case "S":
		t.direction = "W"
	case "W":
		t.direction = "N"
	}
}
