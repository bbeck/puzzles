package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
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
	in.Remove("  ")

	var start string
	in.Scanf("Begin in state %s.", &start)

	var steps int
	in.Scanf("Perform a diagnostic checksum after %d steps.", &steps)

	var program = make(map[string]map[int]Action)
	for in.HasNext() {
		var chunk = in.ChunkS()

		var state string
		chunk.Scanf("In state %s:", &state)

		program[state] = make(map[int]Action)
		for chunk.HasPrefix("If the current value is") {
			var value int
			chunk.Scanf("If the current value is %d:", &value)

			var write int
			chunk.Scanf("- Write the value %d.", &write)

			var move, next string
			chunk.Scanf("- Move one slot to the %s.", &move)
			chunk.Scanf("- Continue with state %s.", &next)

			program[state][value] = Action{Write: write, Move: move, State: next}
		}
	}
	return start, steps, program
}
