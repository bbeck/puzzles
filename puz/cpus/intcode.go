package cpus

import (
	"log"
	"strings"
	"sync/atomic"

	"github.com/bbeck/advent-of-code/puz"
)

type Memory map[int]int

type IntcodeCPU struct {
	Memory  Memory
	ip      int
	bp      int
	stopped atomic.Bool
	Input   func() int
	Output  func(int)
	Halt    func()
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

func (cpu *IntcodeCPU) Execute() {
	for !cpu.stopped.Load() {
		cpu.Step()
	}
}

func (cpu *IntcodeCPU) Stop() {
	cpu.stopped.Store(true)
}

func (cpu *IntcodeCPU) Step() {
	instruction := cpu.Memory[cpu.ip]
	op := instruction % 100
	aMode := (instruction / 100) % 10
	bMode := (instruction / 1000) % 10
	cMode := (instruction / 10000) % 10

	// Read a value obeying the parameter mode
	get := func(addr int, mode int) int {
		switch mode {
		case 0: // position mode
			return cpu.Memory[addr]
		case 1: // immediate mode
			return addr
		case 2: // relative mode
			return cpu.Memory[cpu.bp+addr]
		default:
			log.Fatalf("don't know how to get addr: %d, in mode: %d", addr, mode)
			return -1
		}
	}

	// Write a value obeying the parameter mode
	set := func(addr int, mode int, value int) {
		switch mode {
		case 0: // position mode
			cpu.Memory[addr] = value
		case 2: // relative mode
			cpu.Memory[cpu.bp+addr] = value
		default:
			log.Fatalf("don't know how to set addr: %d, in mode: %d", addr, mode)
		}
	}

	switch op {
	case OpcodeAdd:
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		c := cpu.Memory[cpu.ip+3]
		set(c, cMode, get(a, aMode)+get(b, bMode))
		cpu.ip += 4

	case OpcodeMul:
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		c := cpu.Memory[cpu.ip+3]
		set(c, cMode, get(a, aMode)*get(b, bMode))
		cpu.ip += 4

	case OpcodeIn:
		a := cpu.Memory[cpu.ip+1]
		set(a, aMode, cpu.Input())
		cpu.ip += 2

	case OpcodeOut:
		a := cpu.Memory[cpu.ip+1]
		cpu.Output(get(a, aMode))
		cpu.ip += 2

	case OpcodeJIT: // jump-if-true
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		if get(a, aMode) != 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OpcodeJIF: // jump-if-false
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		if get(a, aMode) == 0 {
			cpu.ip = get(b, bMode)
		} else {
			cpu.ip += 3
		}

	case OpcodeLT: // less than
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		c := cpu.Memory[cpu.ip+3]
		if get(a, aMode) < get(b, bMode) {
			set(c, cMode, 1)
		} else {
			set(c, cMode, 0)
		}
		cpu.ip += 4

	case OpcodeEQ: // equals
		a := cpu.Memory[cpu.ip+1]
		b := cpu.Memory[cpu.ip+2]
		c := cpu.Memory[cpu.ip+3]
		if get(a, aMode) == get(b, bMode) {
			set(c, cMode, 1)
		} else {
			set(c, cMode, 0)
		}
		cpu.ip += 4

	case OpcodeARB: // adjust relative base
		a := cpu.Memory[cpu.ip+1]
		cpu.bp += get(a, aMode)
		cpu.ip += 2

	case OpcodeHLT: // Halt
		if cpu.Halt != nil {
			cpu.Halt()
		}
		cpu.stopped.Store(true)

	default:
		log.Fatalf("unrecognized opcode: %d, instruction:%d, ip:%d", op, instruction, cpu.ip)
	}
}

func InputToIntcodeMemory(year, day int) Memory {
	opcodes := make(Memory)
	for addr, s := range strings.Split(puz.InputToString(year, day), ",") {
		opcodes[addr] = puz.ParseInt(s)
	}

	return opcodes
}
