package main

import (
	"bytes"
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
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
	in.Remove(",")

	var instructions []Instruction
	for in.HasNext() {
		switch in.Byte() {
		case 's':
			var n = in.Int()
			instructions = append(instructions, func(bs []byte) {
				cs := make([]byte, L)
				copy(cs, bs[L-n:])
				copy(cs[n:], bs[:L-n])
				copy(bs, cs)
			})

		case 'x':
			var a, b = in.Int(), in.Int()
			instructions = append(instructions, func(bs []byte) {
				bs[a], bs[b] = bs[b], bs[a]
			})

		case 'p':
			var a = in.Byte()
			in.Expect("/")
			var b = in.Byte()
			instructions = append(instructions, func(bs []byte) {
				ia, ib := bytes.IndexByte(bs, a), bytes.IndexByte(bs, b)
				bs[ia], bs[ib] = bs[ib], bs[ia]
			})
		}
	}

	return instructions
}
