package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	state, steps, program := InputToMachine()

	cursor := 0
	tape := make(map[int]int)

	for n := 0; n < steps; n++ {
		action := program[state][tape[cursor]]
		tape[cursor] = action.Write
		state = action.State

		if action.Move == "left" {
			cursor--
		} else {
			cursor++
		}
	}

	var count int
	for _, value := range tape {
		if value == 1 {
			count++
		}
	}
	fmt.Println(count)
}

type Action struct {
	Write int
	Move  string
	State string
}

func InputToMachine() (string, int, map[string]map[int]Action) {
	var lines []string
	for _, line := range lib.InputToLines() {
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, ":", "")
		line = strings.ReplaceAll(line, "-", "")
		lines = append(lines, strings.TrimSpace(line))
	}

	var start string
	fmt.Sscanf(lines[0], "Begin in state %s", &start)

	var steps int
	fmt.Sscanf(lines[1], "Perform a diagnostic checksum after %d steps", &steps)

	program := make(map[string]map[int]Action)
	for n := 3; n < len(lines); n += 10 {
		var state string
		fmt.Sscanf(lines[n], "In state %s", &state)
		program[state] = make(map[int]Action)

		for value := range []int{0, 1} {
			var write int
			fmt.Sscanf(lines[n+2+4*value], "Write the value %d", &write)

			var move string
			fmt.Sscanf(lines[n+3+4*value], "Move one slot to the %s", &move)

			var next string
			fmt.Sscanf(lines[n+4+4*value], "Continue with state %s", &next)

			program[state][value] = Action{Write: write, Move: move, State: next}
		}
	}

	return start, steps, program
}
