package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"strings"
)

func main() {
	memory := make(map[uint64]uint64)
	for _, i := range InputToInstructions() {
		EnumerateAddresses(i.Address|i.Or, i.Floating, func(address uint64) {
			memory[address] = i.Value
		})
	}

	var sum uint64
	for _, value := range memory {
		sum += value
	}
	fmt.Println(sum)
}

func EnumerateAddresses(address, floating uint64, fn func(uint64)) {
	var ones []int
	for b := 0; b < 36; b++ {
		if floating&(1<<b) > 0 {
			ones = append(ones, b)
		}
	}

	n := uint(len(ones))
	for bits := 0; bits < lib.Pow(2, n); bits++ {
		for i, bit := range ones {
			if bits&(1<<i) == 0 {
				address &= ^(1 << bit)
			} else {
				address |= 1 << bit
			}
		}

		fn(address)
	}
}

type Instruction struct {
	Address, Value, Or, Floating uint64
}

func InputToInstructions() []Instruction {
	lines := lib.InputToLines()

	var or, floating uint64
	var instructions []Instruction
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			for _, c := range line[7:] { // len("mask = ") == 7
				or <<= 1
				floating <<= 1

				if c == 'X' {
					floating |= 1
				} else if c == '1' {
					or |= 1
				}
			}
			continue
		}

		var address, value uint64
		fmt.Sscanf(line, "mem[%d] = %d", &address, &value)

		instructions = append(instructions, Instruction{
			Address:  address,
			Value:    value,
			Or:       or & 0xFFFFFFFFF,
			Floating: floating & 0xFFFFFFFFF,
		})
	}

	return instructions
}
