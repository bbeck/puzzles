package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"reflect"
)

func main() {
	samples, instructions := InputToSamplesAndInstructions()
	operations := IdentifyOperations(samples)

	regs := make([]int, 4)
	for _, instruction := range instructions {
		operations[instruction.OpCode](instruction.A, instruction.B, instruction.C, regs)
	}
	fmt.Println(regs[0])
}

func IdentifyOperations(samples []Sample) []Operation {
	gt := func(a, b int) int {
		if a > b {
			return 1
		}
		return 0
	}

	eq := func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	}

	operations := map[string]Operation{
		"addr": func(a, b, c int, reg []int) { reg[c] = reg[a] + reg[b] },
		"addi": func(a, b, c int, reg []int) { reg[c] = reg[a] + b },
		"mulr": func(a, b, c int, reg []int) { reg[c] = reg[a] * reg[b] },
		"muli": func(a, b, c int, reg []int) { reg[c] = reg[a] * b },
		"banr": func(a, b, c int, reg []int) { reg[c] = reg[a] & reg[b] },
		"bani": func(a, b, c int, reg []int) { reg[c] = reg[a] & b },
		"borr": func(a, b, c int, reg []int) { reg[c] = reg[a] | reg[b] },
		"bori": func(a, b, c int, reg []int) { reg[c] = reg[a] | b },
		"setr": func(a, b, c int, reg []int) { reg[c] = reg[a] },
		"seti": func(a, b, c int, reg []int) { reg[c] = a },
		"gtir": func(a, b, c int, reg []int) { reg[c] = gt(a, reg[b]) },
		"gtri": func(a, b, c int, reg []int) { reg[c] = gt(reg[a], b) },
		"gtrr": func(a, b, c int, reg []int) { reg[c] = gt(reg[a], reg[b]) },
		"eqir": func(a, b, c int, reg []int) { reg[c] = eq(a, reg[b]) },
		"eqri": func(a, b, c int, reg []int) { reg[c] = eq(reg[a], b) },
		"eqrr": func(a, b, c int, reg []int) { reg[c] = eq(reg[a], reg[b]) },
	}

	// Start with every opcode being able to be any operation
	var possibilities [16]Set[string]
	for i := 0; i < len(possibilities); i++ {
		for op := range operations {
			possibilities[i].Add(op)
		}
	}

	// Filter opcode possibilities by what each sample says
	for _, sample := range samples {
		var possible Set[string]
		for op, operation := range operations {
			regs := make([]int, len(sample.Before))
			copy(regs, sample.Before[:])

			operation(sample.A, sample.B, sample.C, regs)

			if reflect.DeepEqual(regs, sample.After[:]) {
				possible.Add(op)
			}
		}

		possibilities[sample.OpCode] = possibilities[sample.OpCode].Intersect(possible)
	}

	// Derive the mapping from opcode to operations by repeatedly going through
	// the list of possibilities and identifying opcodes that only have one
	// possibility.  Once a mapping is identified it can be removed from the set
	// of possibilities for other opcodes.
	var mapped Set[string]
	mapping := make(map[int]string)
	for len(mapping) < len(operations) {
		for opcode, possible := range possibilities {
			if len(possible) == 1 {
				name := possible.Entries()[0]
				mapping[opcode] = name
				mapped.Add(name)
			}
		}

		for opcode := range possibilities {
			possibilities[opcode] = possibilities[opcode].Difference(mapped)
		}
	}

	var index []Operation
	for i := 0; i < len(operations); i++ {
		index = append(index, operations[mapping[i]])
	}
	return index
}

type Operation func(a, b, c int, reg []int)

type Sample struct {
	Before  [4]int
	After   [4]int
	OpCode  int
	A, B, C int
}

type Instruction struct {
	OpCode  int
	A, B, C int
}

func InputToSamplesAndInstructions() ([]Sample, []Instruction) {
	var samples []Sample
	for in.HasNext() {
		if !in.HasPrefix("Before") {
			break
		}

		chunk := in.ChunkS()
		samples = append(samples, Sample{
			Before: [4]int{chunk.Int(), chunk.Int(), chunk.Int(), chunk.Int()},
			OpCode: chunk.Int(),
			A:      chunk.Int(),
			B:      chunk.Int(),
			C:      chunk.Int(),
			After:  [4]int{chunk.Int(), chunk.Int(), chunk.Int(), chunk.Int()},
		})
	}

	var instructions []Instruction
	for in.HasNext() {
		instructions = append(instructions, Instruction{
			OpCode: in.Int(),
			A:      in.Int(),
			B:      in.Int(),
			C:      in.Int(),
		})
	}

	return samples, instructions
}
