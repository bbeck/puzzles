package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"reflect"
	"regexp"
	"strings"
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
		var possibilities aoc.Set[string]
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
	input := aoc.InputToString(2018, 16)
	regex := regexp.MustCompile(`\d+`)

	var nums []int
	for _, s := range regex.FindAllString(input, -1) {
		nums = append(nums, aoc.ParseInt(s))
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
