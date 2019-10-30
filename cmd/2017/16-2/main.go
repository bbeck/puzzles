package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	// Since we're going to have to do this a billion times we need to be more
	// efficient and only do the operations that we need to do.  For this we can
	// assume that there's a cycle involved somewhere in the outputs.  Once we
	// find the cycle we can just use the previously calculated solutions instead
	// of running the instructions over again.
	instructions := InputToInstructions(2017, 16)
	programs := []byte("abcdefghijklmnop")

	seen := make(map[string]string)
	for i := 0; i < 1_000_000_000; i++ {
		// Check and see if we've already seen this permutation.
		key := string(programs)
		if _, ok := seen[key]; ok {
			break
		}

		// This is a new permutation, compute the next one.
		for _, instruction := range instructions {
			programs = instruction(programs)
		}
		seen[key] = string(programs)
	}

	// We've now computed all of the solutions we need to.  This is a cycle so
	// we don't need to iterate a billion times, we just need to iterate part of
	// the last iteration of the loop.
	ps := "abcdefghijklmnop"
	for i := 0; i < 1_000_000_000%len(seen); i++ {
		ps = seen[ps]
	}

	fmt.Printf("programs: %s\n", ps)
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
