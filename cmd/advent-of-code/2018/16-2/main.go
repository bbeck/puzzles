package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"reflect"
	"regexp"
	"strings"
)

func main() {
	operations := IdentifyOperations()

	regs := make([]int, 4)
	for _, instruction := range InputToProgram() {
		operations[instruction.OpCode](instruction.A, instruction.B, instruction.C, regs)
	}
	fmt.Println(regs[0])
}

func IdentifyOperations() []Operation {
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
	var possibilities [16]lib.Set[string]
	for i := 0; i < len(possibilities); i++ {
		for op := range operations {
			possibilities[i].Add(op)
		}
	}

	// Filter opcode possibilities by what each sample says
	for _, sample := range InputToSamples() {
		var possible lib.Set[string]
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
	var mapped lib.Set[string]
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

func InputToSamples() []Sample {
	input := lib.InputToString()
	regex := regexp.MustCompile(`\d+`)

	var nums []int
	for _, s := range regex.FindAllString(input, -1) {
		nums = append(nums, lib.ParseInt(s))
	}

	var samples []Sample
	for i := 0; i < strings.Count(input, "Before"); i++ {
		var sample Sample
		sample.Before = [4]int{nums[12*i+0], nums[12*i+1], nums[12*i+2], nums[12*i+3]}
		sample.OpCode = nums[12*i+4]
		sample.A = nums[12*i+5]
		sample.B = nums[12*i+6]
		sample.C = nums[12*i+7]
		sample.After = [4]int{nums[12*i+8], nums[12*i+9], nums[12*i+10], nums[12*i+11]}
		samples = append(samples, sample)
	}
	return samples
}

type Instruction struct {
	OpCode  int
	A, B, C int
}

func InputToProgram() []Instruction {
	input := lib.InputToString()
	regex := regexp.MustCompile(`\d+`)

	var nums []int
	for _, s := range regex.FindAllString(input, -1) {
		nums = append(nums, lib.ParseInt(s))
	}

	var instructions []Instruction
	for i := 12 * strings.Count(input, "Before"); i < len(nums); i += 4 {
		instructions = append(instructions, Instruction{
			OpCode: nums[i],
			A:      nums[i+1],
			B:      nums[i+2],
			C:      nums[i+3],
		})
	}
	return instructions
}
