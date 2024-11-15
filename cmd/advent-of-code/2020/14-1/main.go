package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	memory := make(map[uint64]uint64)
	for _, i := range InputToInstructions() {
		memory[i.Address] = (i.Value & i.And) | i.Or
	}

	var sum uint64
	for _, value := range memory {
		sum += value
	}
	fmt.Println(sum)
}

type Instruction struct {
	Address, Value, And, Or uint64
}

func InputToInstructions() []Instruction {
	lines := puz.InputToLines()

	var and, or uint64
	var instructions []Instruction
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			for _, c := range line[7:] { // len("mask = ") == 7
				and <<= 1
				or <<= 1

				if c == 'X' {
					and |= 1
				} else if c == '0' {
					and |= 0
				} else if c == '1' {
					or |= 1
				}
			}
			continue
		}

		var address, value uint64
		fmt.Sscanf(line, "mem[%d] = %d", &address, &value)

		instructions = append(instructions, Instruction{
			Address: address,
			Value:   value,
			And:     and & 0xFFFFFFFFF,
			Or:      or & 0xFFFFFFFFF,
		})
	}

	return instructions
}
