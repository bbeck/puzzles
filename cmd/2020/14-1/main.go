package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	instructions := InputToProgram(2020, 14)

	memory := make(map[int]int)
	for _, instruction := range instructions {
		value := (instruction.value & instruction.andMask) | instruction.orMask
		memory[instruction.address] = value
	}

	var sum int
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

type Instruction struct {
	andMask int
	orMask  int
	address int
	value   int
}

func InputToProgram(year, day int) []Instruction {
	var instructions []Instruction

	var andMask, orMask int
	for _, line := range aoc.InputToLines(year, day) {
		var mask string
		if _, err := fmt.Sscanf(line, "mask = %s", &mask); err == nil {
			andMask = 0
			orMask = 0

			for _, c := range mask {
				andMask <<= 1
				orMask <<= 1

				switch c {
				case '0':
					andMask |= 0
				case '1':
					orMask |= 1
				case 'X':
					andMask |= 1
					orMask |= 0
				}
			}
			continue
		}

		var address, value int
		if _, err := fmt.Sscanf(line, "mem[%d] = %d", &address, &value); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		instructions = append(instructions, Instruction{
			andMask: andMask,
			orMask:  orMask,
			address: address,
			value:   value,
		})
	}

	return instructions
}
