package main

import (
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
	"reflect"
)

func main() {
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

	var count int
	for _, sample := range InputToSamples() {
		var possibilities Set[string]
		for op, instruction := range operations {
			regs := make([]int, len(sample.Before))
			copy(regs, sample.Before[:])

			instruction(sample.A, sample.B, sample.C, regs)

			if reflect.DeepEqual(regs, sample.After[:]) {
				possibilities.Add(op)
			}
		}

		if len(possibilities) >= 3 {
			count++
		}
	}

	fmt.Println(count)
}

type Operation func(a, b, c int, reg []int)

type Sample struct {
	Before  [4]int
	After   [4]int
	OpCode  int
	A, B, C int
}

func InputToSamples() []Sample {
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

	return samples
}
