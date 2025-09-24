package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/cpus"
)

func main() {
	memory := cpus.InputToIntcodeMemory()

	grid, turtle := Load(memory)
	path := BuildPath(grid, turtle)

	main, A, B, C := BuildRules(path)

	// Set up our program for the robot.  We want to tell it our routines and
	// then send it a n to indicate we don't want realtime updates as the robot
	// moves.
	var program strings.Builder
	for _, s := range []string{main, A, B, C, "n"} {
		program.WriteString(s)
		program.WriteByte('\n')
	}

	// Force the robot to wake.
	memory[0] = 2

	// Wire the program as input.
	input := make(chan int, program.Len())
	for _, c := range program.String() {
		input <- int(c)
	}

	// Run the robot keeping track of the last value outputted.
	var dust int
	robot := cpus.IntcodeCPU{
		Memory: memory,
		Input:  func() int { return <-input },
		Output: func(value int) { dust = value },
	}
	robot.Execute()

	fmt.Println(dust)
}

type Step struct {
	Turn    Heading
	Forward int
}

func BuildPath(grid Set[Point2D], robot Turtle) []Step {
	tryForward := func(r Turtle) bool {
		r.Forward(1)
		return grid.Contains(r.Location)
	}

	tryLeft := func(r Turtle) bool {
		r.TurnLeft()
		return tryForward(r)
	}

	tryRight := func(r Turtle) bool {
		r.TurnRight()
		return tryForward(r)
	}

	// Visualizing the map it's just a line that overlaps itself that needs to be
	// followed.  Because the robot starts off not facing the direction of the
	// line, we're going to make the assumption that each portion of the path is
	// structured as a turn followed by some number of steps to move forward.
	var path []Step
	for {
		var turn Heading
		if tryLeft(robot) {
			turn = Left
			robot.TurnLeft()
		} else if tryRight(robot) {
			turn = Right
			robot.TurnRight()
		} else {
			break
		}

		var count int
		for count = 0; tryForward(robot); count++ {
			robot.Forward(1)
		}

		path = append(path, Step{Turn: turn, Forward: count})
	}

	return path
}

func BuildRules(steps []Step) (string, string, string, string) {
	// We need to determine A, B, and C as sequences of steps that can be
	// combined to generate the path exactly.  Each of the A, B, and C sequences
	// cannot be longer than 20 characters.
	//
	// To accomplish this we'll work with the path as a string.  We'll choose a
	// prefix of the string to be one of our sequences, then remove any
	// occurrences of the sequence in the string.  If after the 3rd sequence is
	// chosen the path string is empty then we know we've found a solution.
	var path string
	for _, step := range steps {
		path = path + fmt.Sprintf("%s,%d,", step.Turn, step.Forward)
	}

	choose := func(s string) []string {
		// Only consider substrings that end on a comma.  This ensures we always
		// work with full steps.
		var choices []string
		for i := 0; i < len(s) && i < 20; i++ {
			if s[i] == ',' {
				choices = append(choices, s[:i+1])
			}
		}

		return choices
	}

	var A, B, C string
loop:
	for _, A = range choose(path) {
		withoutA := strings.ReplaceAll(path, A, "")
		for _, B = range choose(withoutA) {
			withoutB := strings.ReplaceAll(withoutA, B, "")
			for _, C = range choose(withoutB) {
				withoutC := strings.ReplaceAll(withoutB, C, "")
				if withoutC == "" {
					break loop
				}
			}
		}
	}

	main := path
	main = strings.ReplaceAll(main, A, "A,")
	main = strings.ReplaceAll(main, B, "B,")
	main = strings.ReplaceAll(main, C, "C,")
	main = strings.TrimRight(main, ",")
	A = strings.TrimRight(A, ",")
	B = strings.TrimRight(B, ",")
	C = strings.TrimRight(C, ",")
	return main, A, B, C
}

func Load(memory cpus.Memory) (Set[Point2D], Turtle) {
	var grid Set[Point2D]
	var robot Turtle

	// Build the grid.
	var current Point2D
	cpu := cpus.IntcodeCPU{
		Memory: memory.Copy(),
		Output: func(value int) {
			switch value {
			case '.':
				current = current.Right()
			case '#':
				grid.Add(current)
				current = current.Right()
			case '^':
				grid.Add(current)
				robot = Turtle{Location: current, Heading: Up}
				current = current.Right()
			case 'v':
				grid.Add(current)
				robot = Turtle{Location: current, Heading: Down}
				current = current.Right()
			case '<':
				grid.Add(current)
				robot = Turtle{Location: current, Heading: Left}
				current = current.Right()
			case '>':
				grid.Add(current)
				robot = Turtle{Location: current, Heading: Right}
				current = current.Right()
			case '\n':
				current = Point2D{X: 0, Y: current.Y + 1}
			}
		},
	}
	cpu.Execute()

	return grid, robot
}
