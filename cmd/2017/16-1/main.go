package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	programs := []byte("abcdefghijklmnop")
	for _, instruction := range InputToInstructions() {
		instruction(programs)
	}

	fmt.Println(string(programs))
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
