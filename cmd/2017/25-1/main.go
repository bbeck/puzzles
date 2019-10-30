package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	state, N, instructions := InputToMachine(2017, 25)
	tape := make(map[int]int)
	pos := 0

	for n := 0; n < N; n++ {
		for _, instr := range instructions {
			if instr.state != state {
				continue
			}

			// fmt.Printf("      %+v\n", instr)

			value := tape[pos]
			if value == 0 {
				tape[pos] = instr.zvalue
				pos += instr.zdelta
				state = instr.zstate
			} else {
				tape[pos] = instr.ovalue
				pos += instr.odelta
				state = instr.ostate
			}
			break
		}
	}

	var count int
	for _, v := range tape {
		if v == 1 {
			count++
		}
	}

	fmt.Printf("checksum: %d\n", count)
}

type Instruction struct {
	state          string
	zvalue, ovalue int    // value to write to the tape
	zdelta, odelta int    // change to position
	zstate, ostate string // state to transition to
}

func InputToMachine(year, day int) (string, int, []Instruction) {
	lines := aoc.InputToLines(year, day)
	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], ".", "")
		lines[i] = strings.ReplaceAll(lines[i], ":", "")
		lines[i] = strings.ReplaceAll(lines[i], "-", "")
		lines[i] = strings.Trim(lines[i], " ")
	}

	// Parse the initial state
	var initial string
	if _, err := fmt.Sscanf(lines[0], "Begin in state %s", &initial); err != nil {
		log.Fatalf("unable to parse initial state: %s", lines[0])
	}

	// Parse the number of iterations
	var N int
	if _, err := fmt.Sscanf(lines[1], "Perform a diagnostic checksum after %d steps", &N); err != nil {
		log.Fatalf("unable to parse num iterations: %s", lines[1])
	}

	var instructions []Instruction
	for n := 3; n < len(lines); n += 10 {
		var state string
		if _, err := fmt.Sscanf(lines[n], "In state %s", &state); err != nil {
			log.Fatalf("unable to parse state: %s", lines[n])
		}

		var zvalue int
		if _, err := fmt.Sscanf(lines[n+2], "Write the value %d", &zvalue); err != nil {
			log.Fatalf("unable to parse value line: %s", lines[n+2])
		}

		var zdirection string
		if _, err := fmt.Sscanf(lines[n+3], "Move one slot to the %s", &zdirection); err != nil {
			log.Fatalf("unable to parse value line: %s", lines[n+3])
		}
		var zdelta int
		if zdirection == "left" {
			zdelta = -1
		} else {
			zdelta = 1
		}

		var zstate string
		if _, err := fmt.Sscanf(lines[n+4], "Continue with state %s", &zstate); err != nil {
			log.Fatalf("unable to parse next state line: %s", lines[n+4])
		}

		var ovalue int
		if _, err := fmt.Sscanf(lines[n+6], "Write the value %d", &ovalue); err != nil {
			log.Fatalf("unable to parse value line: %s", lines[n+6])
		}

		var odirection string
		if _, err := fmt.Sscanf(lines[n+7], "Move one slot to the %s", &odirection); err != nil {
			log.Fatalf("unable to parse value line: %s", lines[n+7])
		}
		var odelta int
		if odirection == "left" {
			odelta = -1
		} else {
			odelta = 1
		}

		var ostate string
		if _, err := fmt.Sscanf(lines[n+8], "Continue with state %s", &ostate); err != nil {
			log.Fatalf("unable to parse next state line: %s", lines[n+8])
		}

		instructions = append(instructions, Instruction{
			state:  state,
			zvalue: zvalue,
			zdelta: zdelta,
			zstate: zstate,
			ovalue: ovalue,
			odelta: odelta,
			ostate: ostate,
		})
	}

	return initial, N, instructions
}
