package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	program := InputToProgram(2017, 18)

	bus := NewBus(2)

	go Run(program, 0, bus)
	count := Run(program, 1, bus)
	fmt.Printf("send count: %d\n", count)
}

func Run(program []Instruction, id int, bus *Bus) int {
	registers := map[string]int{
		"p": id,
	}

	// choose between the value in a register and an immediate
	choose := func(r string, i int) int {
		if r != "" {
			return registers[r]
		}

		return i
	}

	// get the test value for a jgz
	test := func(i Instruction) int {
		v, err := strconv.Atoi(i.target)
		if err == nil {
			return v
		}

		return registers[i.target]
	}

	var sends int
	for pc := 0; pc >= 0 && pc < len(program); {
		instruction := program[pc]

		switch instruction.op {
		case "snd":
			value := choose(instruction.register, instruction.immediate)
			if !bus.Send(id, value) {
				return sends
			}
			sends++
			pc++

		case "set":
			registers[instruction.target] = choose(instruction.register, instruction.immediate)
			pc++

		case "add":
			registers[instruction.target] += choose(instruction.register, instruction.immediate)
			pc++

		case "mul":
			registers[instruction.target] *= choose(instruction.register, instruction.immediate)
			pc++

		case "mod":
			registers[instruction.target] %= choose(instruction.register, instruction.immediate)
			pc++

		case "rcv":
			value, ok := bus.Receive(id)
			if !ok {
				return sends
			}

			registers[instruction.register] = value
			pc++

		case "jgz":
			if test(instruction) > 0 {
				pc += choose(instruction.register, instruction.immediate)
			} else {
				pc++
			}
		}
	}

	return sends
}

// A bus is a construct that allows two goroutines to exchange information with
// each other, but prevents deadlocks from happening.
type Bus struct {
	mutex   sync.Mutex
	chans   []chan int
	blocked []bool
	other   func(id int) int
	done    bool
}

func NewBus(n int) *Bus {
	chans := make([]chan int, n)
	for i := 0; i < n; i++ {
		chans[i] = make(chan int, 1000)
	}

	return &Bus{
		mutex:   sync.Mutex{},
		chans:   chans,
		blocked: make([]bool, n),
		other:   func(id int) int { return n - 1 - id },
	}
}

func (b *Bus) Send(id int, data int) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	// We can always send data to a process as long as we're not done.
	// if b.done {
	// 	return false
	// }

	b.chans[b.other(id)] <- data
	return true
}

func (b *Bus) Receive(id int) (int, bool) {
	b.mutex.Lock()

	// If we're done then return failure.
	if b.done {
		b.mutex.Unlock()
		return 0, false
	}

	// Try a non-blocking read to see if there's data available.
	select {
	case value, ok := <-b.chans[id]:
		if ok {
			b.mutex.Unlock()
			return value, true
		}

	default:
		// There wasn't any data, continue.
	}

	// We fell through so there wasn't any data ready.  We're going to have to
	// block, but first we need to make sure the other process isn't also blocked
	// trying to read.  If it is then we have deadlock and we need to finish.
	if b.blocked[b.other(id)] {
		// Don't let any new receives happen
		b.done = true

		// Wake up anyone that's currently receiving
		for _, c := range b.chans {
			close(c)
		}

		b.mutex.Unlock()
		return 0, false
	}

	// Prepare to block while looking for a value.
	b.blocked[id] = true

	// Release the mutex to allow the other goroutine to enter.  We'll acquire it
	// again as soon as we wake up.
	b.mutex.Unlock()
	value, ok := <-b.chans[id]
	b.mutex.Lock()

	if !ok {
		// We didn't read a value, our channel must have been closed.  This means
		// we've reached a deadlock.  Return a failure.
		b.mutex.Unlock()
		return 0, false
	}

	b.blocked[id] = false
	b.mutex.Unlock()
	return value, true
}

type Instruction struct {
	op        string
	target    string
	immediate int
	register  string
}

func (i Instruction) String() string {
	if i.register != "" {
		return fmt.Sprintf("%s %s %s", i.op, i.target, i.register)
	} else {
		return fmt.Sprintf("%s %s %d", i.op, i.target, i.immediate)
	}
}

func InputToProgram(year, day int) []Instruction {
	// parse an argument as either an immediate or the register it refers to
	parse := func(s string) (int, string) {
		immediate, err := strconv.Atoi(s)
		if err == nil {
			return immediate, ""
		}

		return 0, s
	}

	var program []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		tokens := strings.Split(line, " ")

		var target string
		var immediate int
		var register string

		switch len(tokens) {
		case 2:
			immediate, register = parse(tokens[1])

		case 3:
			target = tokens[1]
			immediate, register = parse(tokens[2])
		}

		program = append(program, Instruction{
			op:        tokens[0],
			target:    target,
			immediate: immediate,
			register:  register,
		})
	}

	return program
}
