package main

import (
	"bytes"
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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

	current := WalkCycle(start, 1_000_000_000, dance)
	fmt.Println(current)
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
