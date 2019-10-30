package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	programs := []byte("abcdefghijklmnop")
	for _, instruction := range InputToInstructions(2017, 16) {
		programs = instruction(programs)
	}

	fmt.Printf("programs: %s\n", string(programs))
}

type Instruction func([]byte) []byte

func InputToInstructions(year, day int) []Instruction {
	var instructions []Instruction
	for _, s := range strings.Split(aoc.InputToString(year, day), ",") {
		var instruction Instruction

		switch s[0] {
		case 's':
			size := aoc.ParseInt(s[1:])
			instruction = func(bs []byte) []byte {
				L := len(bs)
				return append(bs[L-size:], bs[:L-size]...)
			}

		case 'x':
			parts := strings.Split(s[1:], "/")
			pos1 := aoc.ParseInt(parts[0])
			pos2 := aoc.ParseInt(parts[1])
			instruction = func(bs []byte) []byte {
				bs[pos1], bs[pos2] = bs[pos2], bs[pos1]
				return bs
			}

		case 'p':
			parts := strings.Split(s[1:], "/")
			id1 := parts[0][0]
			id2 := parts[1][0]
			instruction = func(bs []byte) []byte {
				for i := 0; i < len(bs); i++ {
					if bs[i] == id1 {
						bs[i] = id2
					} else if bs[i] == id2 {
						bs[i] = id1
					}
				}

				return bs
			}

		default:
			log.Fatalf("unable to parse instruction: %s", s)
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
