package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	instructions := map[string]Instruction{
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

	var count int
	for _, sample := range InputToSamples(2018, 16) {
		var n int
		for _, instruction := range instructions {
			after := instruction(sample.a, sample.b, sample.c, sample.before)
			if after.Equal(sample.after) {
				n++
			}
		}

		if n >= 3 {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

type Registers []int

func (r Registers) Equal(other Registers) bool {
	for i := 0; i < len(r); i++ {
		if r[i] != other[i] {
			return false
		}
	}

	return true
}

type Instruction func(a, b, c int, registers Registers) Registers

type Sample struct {
	before  Registers
	after   Registers
	op      int
	a, b, c int
}

func InputToSamples(year, day int) []Sample {
	lines := aoc.InputToLines(year, day)

	registers := func(s string) Registers {
		var a, b, c, d int
		if _, err := fmt.Sscanf(s, "[%d, %d, %d, %d]", &a, &b, &c, &d); err != nil {
			log.Fatalf("unable to parse registers: %s", s)
		}

		return Registers{a, b, c, d}
	}

	var samples []Sample
	for i := 0; i < len(lines); i += 4 {
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

	return samples
}
