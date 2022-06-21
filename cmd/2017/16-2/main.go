package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	instructions := InputToInstructions()

	dance := func(s string) string {
		bs := []byte(s)
		for _, instruction := range instructions {
			instruction(bs)
		}
		return string(bs)
	}

	start := "abcdefghijklmnop"

	// Assume there's a cycle, keep dancing until we reach the starting state.
	seen := make(map[string]string)
	current := start
	for i := 0; i < 1000; i++ {
		if _, ok := seen[current]; ok {
			break
		}

		next := dance(current)
		seen[current] = next

		current = next
		if current == start {
			break
		}
	}

	remainder := 1_000_000_000 % len(seen)
	for i := 0; i < remainder; i++ {
		current = dance(current)
	}

	fmt.Println(current)
}

type Instruction func([]byte)

const L = 16

func InputToInstructions() []Instruction {
	fields := strings.Split(aoc.InputToString(2017, 16), ",")

	var instructions []Instruction
	for _, field := range fields {
		var instruction Instruction

		switch field[0] {
		case 's':
			n := aoc.ParseInt(field[1:])
			instruction = func(bs []byte) {
				cs := make([]byte, L)
				copy(cs, bs[L-n:])
				copy(cs[n:], bs[:L-n])
				copy(bs, cs)
			}

		case 'x':
			sa, sb, _ := strings.Cut(field[1:], "/")
			a, b := aoc.ParseInt(sa), aoc.ParseInt(sb)
			instruction = func(bs []byte) {
				bs[a], bs[b] = bs[b], bs[a]
			}

		case 'p':
			a, b, _ := strings.Cut(field[1:], "/")
			instruction = func(bs []byte) {
				ia, ib := bytes.IndexByte(bs, a[0]), bytes.IndexByte(bs, b[0])
				bs[ia], bs[ib] = bs[ib], bs[ia]
			}
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
