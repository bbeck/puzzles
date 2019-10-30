package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	samples, program := InputToSamplesAndProgram(2018, 16)
	mapping := DeriveInstructionMapping(samples)

	registers := Registers{0, 0, 0, 0}
	for _, instruction := range program {
		op, a, b, c := instruction.op, instruction.a, instruction.b, instruction.c
		registers = mapping[op](a, b, c, registers)
	}

	for i := 0; i < 4; i++ {
		fmt.Printf("register %d: %d\n", i, registers[i])
	}
}

// /////////////////////////////////////////////////////////////////////////////

type Instruction func(a, b, c int, registers Registers) Registers

var instructions = map[string]Instruction{
	"addr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] + registers[b]
		return registers
	},
	"addi": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] + b
		return registers
	},
	"mulr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] * registers[b]
		return registers
	},
	"muli": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] * b
		return registers
	},
	"banr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] & registers[b]
		return registers
	},
	"bani": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] & b
		return registers
	},
	"borr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] | registers[b]
		return registers
	},
	"bori": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a] | b
		return registers
	},
	"setr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = registers[a]
		return registers
	},
	"seti": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		registers[c] = a
		return registers
	},
	"gtir": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if a > registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
	"gtri": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if registers[a] > b {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
	"gtrr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if registers[a] > registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
	"eqir": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if a == registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
	"eqri": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if registers[a] == b {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
	"eqrr": func(a, b, c int, registers Registers) Registers {
		registers = append(Registers{}, registers...)
		if registers[a] == registers[b] {
			registers[c] = 1
		} else {
			registers[c] = 0
		}
		return registers
	},
}

func DeriveInstructionMapping(samples []Sample) map[int]Instruction {
	opcodes := make([]string, 0)
	for opcode := range instructions {
		opcodes = append(opcodes, opcode)
	}

	// initially every op id can be any op code
	possibilities := make(map[int][]string)
	for i := 0; i < 16; i++ {
		possibilities[i] = append([]string{}, opcodes...)
	}

	// intersect two sets of strings
	intersect := func(s1, s2 []string) []string {
		m := make(map[string]bool)
		for _, s := range s1 {
			m[s] = true
		}

		var result []string
		for _, s := range s2 {
			if m[s] {
				result = append(result, s)
			}
		}

		return result
	}

	// For each sample's op, determine the possible op codes
	for _, sample := range samples {
		var matches []string
		for opcode, instruction := range instructions {
			after := instruction(sample.a, sample.b, sample.c, sample.before)
			if after.Equals(sample.after) {
				matches = append(matches, opcode)
			}
		}

		possibilities[sample.op] = intersect(possibilities[sample.op], matches)
	}

	// subtract any entry from a set of strings that's in another set of strings
	subtract := func(s1, s2 []string) []string {
		m := make(map[string]bool)
		for _, s := range s2 {
			m[s] = true
		}

		var result []string
		for _, s := range s1 {
			if !m[s] {
				result = append(result, s)
			}
		}
		return result
	}

	// Keep looping until we have a mapping for every instruction
	mapping := make(map[int]Instruction)
	used := make([]string, 0) // the used op codes
	for len(mapping) < len(instructions) {
		for op, opcodes := range possibilities {
			// Remove any already used op codes
			opcodes = subtract(opcodes, used)

			// If only one choice remains then we can assign it
			if len(opcodes) == 1 {
				mapping[op] = instructions[opcodes[0]]
				used = append(used, opcodes[0])
			}
		}
	}

	return mapping
}

// /////////////////////////////////////////////////////////////////////////////

type Registers []int

func (r Registers) Equals(other Registers) bool {
	for i := 0; i < len(r); i++ {
		if r[i] != other[i] {
			return false
		}
	}

	return true
}

// /////////////////////////////////////////////////////////////////////////////

type Sample struct {
	before  Registers
	after   Registers
	op      int
	a, b, c int
}

type CompiledInstruction struct {
	op, a, b, c int
}

func InputToSamplesAndProgram(year, day int) ([]Sample, []CompiledInstruction) {
	lines := aoc.InputToLines(year, day)

	registers := func(s string) Registers {
		var a, b, c, d int
		if _, err := fmt.Sscanf(s, "[%d, %d, %d, %d]", &a, &b, &c, &d); err != nil {
			log.Fatalf("unable to parse registers: %s", s)
		}

		return Registers{a, b, c, d}
	}

	var samples []Sample
	var i int
	for i = 0; i < len(lines); i += 4 {
		if lines[i] == "" && lines[i+1] == "" {
			break
		}

		before := registers(lines[i+0][8:])
		after := registers(lines[i+2][8:])

		var op, a, b, c int
		if _, err := fmt.Sscanf(lines[i+1], "%d %d %d %d", &op, &a, &b, &c); err != nil {
			log.Fatalf("unable to parse operation: %s", lines[i+1])
		}

		samples = append(samples, Sample{
			before: before,
			after:  after,
			op:     op,
			a:      a,
			b:      b,
			c:      c,
		})
	}

	// There's 2 blank lines between the samples and the program
	i += 2

	var program []CompiledInstruction
	for ; i < len(lines); i++ {
		var op, a, b, c int
		if _, err := fmt.Sscanf(lines[i], "%d %d %d %d", &op, &a, &b, &c); err != nil {
			log.Fatalf("unable to parse instruction: %s", lines[i])
		}

		program = append(program, CompiledInstruction{op: op, a: a, b: b, c: c})
	}

	return samples, program
}
