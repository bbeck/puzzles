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
		// Each instruction can target multiple addresses
		for _, address := range EnumerateAddresses(instruction) {
			memory[address] = instruction.value
		}
	}

	var sum int
	for _, v := range memory {
		sum += v
	}
	fmt.Println(sum)
}

func EnumerateAddresses(instruction Instruction) []int {
	// First build an and-mask and or-mask for the address.  Everywhere in the
	// mask that has a 1 will need to appear as part of the or-mask and everywhere
	// in the mask that has an 0 will need to appear as a 1 in the and-mask.
	var andMask, orMask int
	for _, c := range instruction.mask {
		andMask <<= 1
		orMask <<= 1

		switch c {
		case '0':
			andMask |= 1

		case '1':
			orMask |= 1
		}
	}
	address := (instruction.address & andMask) | orMask

	// Build a bit vector with a single bit set that corresponds to each position
	// in the mask that contains an X.
	var bitvectors []int
	for index, c := range instruction.mask {
		if c == 'X' {
			bitvectors = append(bitvectors, 1<<(len(instruction.mask)-index-1))
		}
	}

	// Next, for each bit vector generate 2 addresses based on the existing set
	// of addresses.  One that has a 0 in the set bit and one that has a 1 in the
	// set bit.
	set := map[int]struct{}{
		address: {},
	}
	for _, bv := range bitvectors {
		newSet := make(map[int]struct{})
		for address := range set {
			newSet[address|bv] = struct{}{}
			newSet[address & ^bv] = struct{}{}
		}

		set = newSet
	}

	// Finally convert the address set into a list of addresses
	var addresses []int
	for address := range set {
		addresses = append(addresses, address)
	}
	return addresses
}

type Instruction struct {
	mask    string
	address int
	value   int
}

func InputToProgram(year, day int) []Instruction {
	var instructions []Instruction

	var mask string
	for _, line := range aoc.InputToLines(year, day) {
		if _, err := fmt.Sscanf(line, "mask = %s", &mask); err == nil {
			continue
		}

		var address, value int
		if _, err := fmt.Sscanf(line, "mem[%d] = %d", &address, &value); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		instructions = append(instructions, Instruction{
			mask:    mask,
			address: address,
			value:   value,
		})
	}

	return instructions
}
