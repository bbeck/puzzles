package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/bbeck/advent-of-code/aoc"
)

var n int64 = 119_315_717_514_047 // <- prime
var N = big.NewInt(n)
var t int64 = 101_741_582_076_661 // <- prime
var T = big.NewInt(t)

func main() {
	//  https://www.reddit.com/r/adventofcode/comments/ee0rqi/2019_day_22_solutions/fbnkaju

	var offset, increment = big.NewInt(0), big.NewInt(1)
	for _, instruction := range InputToInstructions(2019, 22) {
		switch instruction.kind {
		case DealNewStackKind:
			increment.Neg(increment)
			offset.Add(offset, increment)
			offset.Mod(offset, N)

		case CutKind:
			var temp big.Int
			temp.Mul(increment, big.NewInt(instruction.arg))

			offset.Add(offset, &temp)
			offset.Mod(offset, N)

		case DealWithIncrementKind:
			var inverse big.Int
			inverse.ModInverse(big.NewInt(instruction.arg), N)

			increment.Mul(increment, &inverse)
			increment.Mod(increment, N)
		}
	}

	// 1 pass of the shuffle is now encoded in offset and increment.  To figure
	// out the values of offset and increment after T shuffles we can use the
	// following:
	//
	//   incrementT = increment^T mod N
	//
	//   offsetT = offset * (1 + increment + increment^2 + ... + increment^T) mod N
	//           = offset * (1 - increment^T)/(1 - increment) mod N
	//           = offset * (1 - increment^T) * inverse(1 - increment) mod N
	var incrementT big.Int
	incrementT.Exp(increment, T, N)

	var temp1 big.Int
	temp1.Exp(increment, T, N)
	temp1.Sub(big.NewInt(1), &temp1)

	var temp2 big.Int
	temp2.Sub(big.NewInt(1), increment)
	temp2.ModInverse(&temp2, N)

	var offsetT big.Int
	offsetT.Mul(offset, &temp1)
	offsetT.Mul(&offsetT, &temp2)
	offsetT.Mod(&offsetT, N)

	var answer big.Int
	answer.Mul(&incrementT, big.NewInt(2020))
	answer.Add(&answer, &offsetT)
	answer.Mod(&answer, N)
	fmt.Println("answer:", &answer)
}

const (
	DealNewStackKind      string = "DealNewStack"
	CutKind               string = "Cut"
	DealWithIncrementKind string = "DealWithIncrement"
)

type Instruction struct {
	kind string
	arg  int64
}

func InputToInstructions(year, day int) []Instruction {
	var instructions []Instruction
	for _, line := range aoc.InputToLines(year, day) {
		var argument int64

		if _, err := fmt.Sscanf(line, "deal with increment %d", &argument); err == nil {
			instructions = append(instructions, Instruction{
				kind: DealWithIncrementKind,
				arg:  argument,
			})
			continue
		}

		if line == "deal into new stack" {
			instructions = append(instructions, Instruction{
				kind: DealNewStackKind,
			})
			continue
		}

		if _, err := fmt.Sscanf(line, "cut %d", &argument); err == nil {
			instructions = append(instructions, Instruction{
				kind: CutKind,
				arg:  argument,
			})
			continue
		}

		log.Fatalf("unable to parse line: %s", line)
	}

	return instructions
}
