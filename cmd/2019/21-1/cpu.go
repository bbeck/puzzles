package main

import (
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

type Memory map[int]int

type CPU struct {
	memory  Memory
	ip      int
	bp      int
	input   func(addr int) int
	output  func(value int)
	halt    func()
	stopped bool
}

const (
	OpcodeAdd = 1
	OpcodeMul = 2
	OpcodeIn  = 3
	OpcodeOut = 4
	OpcodeJIT = 5
	OpcodeJIF = 6
	OpcodeLT  = 7
	OpcodeEQ  = 8
	OpcodeARB = 9
	OpcodeHLT = 99
)

func (cpu *CPU) Execute() {
	for {
		cpu.Step()
		if cpu.stopped {
			break
		}
	}
}

func (cpu *CPU) Stop() {
	cpu.stopped = true
}

func (cpu *CPU) Step() {
	instruction := cpu.memory[cpu.ip]
	op := instruction % 100
	aMode := (instruction / 100) % 10
	bMode := (instruction / 1000) % 10
	cMode := (instruction / 10000) % 10

	// Read a value obeying the parameter mode
	get := func(addr int, mode int) int {
		switch mode {
		case 0: // position mode
			return cpu.memory[addr]
		case 1: // immediate mode
			return addr
		case 2: // relative mode
			return cpu.memory[cpu.bp+addr]
		}

		log.Fatalf("don't know how to get addr: %d, in mode: %d\addr", addr, mode)
		return -1
	}

	// Write a value obeying the parameter mode
	set := func(addr int, value int, mode int) {
		switch mode {
		case 0: // position mode
			cpu.memory[addr] = value
		case 2: // relative mode
			cpu.memory[cpu.bp+addr] = value
		default:
			log.Fatalf("don't know how to set addr: %d, in mode: %d\addr", addr, mode)
		}
	}

	switch op {
	case OpcodeAdd:
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		set(c, get(a, aMode)+get(b, bMode), cMode)
		cpu.ip += 4

	case OpcodeMul:
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		set(c, get(a, aMode)*get(b, bMode), cMode)
		cpu.ip += 4

	case OpcodeIn:
		a := cpu.memory[cpu.ip+1]
		set(a, cpu.input(a), aMode)
		cpu.ip += 2

	case OpcodeOut:
		a := cpu.memory[cpu.ip+1]
		cpu.output(get(a, aMode))
		cpu.ip += 2

	case OpcodeJIT: // jump-if-true
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		if get(a, aMode) != 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OpcodeJIF: // jump-if-false
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		if get(a, aMode) == 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OpcodeLT: // less than
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		if get(a, aMode) < get(b, bMode) {
			set(c, 1, cMode)
		} else {
			set(c, 0, cMode)
		}
		cpu.ip += 4

	case OpcodeEQ: // equals
		a := cpu.memory[cpu.ip+1]
		b := cpu.memory[cpu.ip+2]
		c := cpu.memory[cpu.ip+3]
		if get(a, aMode) == get(b, bMode) {
			set(c, 1, cMode)
		} else {
			set(c, 0, cMode)
		}
		cpu.ip += 4

	case OpcodeARB: // adjust relative base
		a := cpu.memory[cpu.ip+1]
		cpu.bp += get(a, aMode)
		cpu.ip += 2

	case OpcodeHLT: // halt
		if cpu.halt != nil {
			cpu.halt()
		}
		cpu.stopped = true

	default:
		log.Fatalf("unrecognized opcode: %d, instruction:%d, ip:%d", op, instruction, cpu.ip)
	}
}

func InputToMemory(year, day int) Memory {
	opcodes := make(Memory)
	for addr, s := range strings.Split(aoc.InputToString(year, day), ",") {
		opcodes[addr] = aoc.ParseInt(s)
	}

	return opcodes
}
