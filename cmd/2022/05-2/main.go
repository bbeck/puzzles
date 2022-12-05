package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	crates, instructions := InputToCrates(), InputToInstructions()
	for _, instr := range instructions {
		var elems aoc.Deque[rune]
		for n := 0; n < instr.Count; n++ {
			elems.PushBack(crates[instr.Start-1].PopFront())
		}
		for n := 0; n < instr.Count; n++ {
			crates[instr.End-1].PushFront(elems.PopBack())
		}
	}

	var sb strings.Builder
	for _, crate := range crates {
		sb.WriteRune(crate.PeekFront())
	}
	fmt.Println(sb.String())
}

type Crate struct {
	aoc.Deque[rune]
}

func InputToCrates() []Crate {
	lines := aoc.InputToLines(2022, 5)
	N := (len(lines[0]) + 1) / 4 // Lines are padded with trailing spaces

	crates := make([]Crate, N)
	for _, line := range lines {
		if strings.Index(line, "[") == -1 {
			break
		}

		for n := 0; n < N; n++ {
			if c := line[4*n+1]; c != ' ' {
				crates[n].PushBack(rune(c))
			}
		}
	}
	return crates
}

type Instruction struct {
	Count, Start, End int
}

func InputToInstructions() []Instruction {
	lines := aoc.InputToLines(2022, 5)

	var i int
	for i = 0; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}
	}

	var instructions []Instruction
	for i = i + 1; i < len(lines); i++ {
		var instr Instruction
		_, _ = fmt.Sscanf(lines[i], "move %d from %d to %d", &instr.Count, &instr.Start, &instr.End)
		instructions = append(instructions, instr)
	}

	return instructions
}
