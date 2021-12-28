package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"strings"
)

func main() {
	// The input is a sequence of assembly language instructions performing some complex
	// calculation on an 14-digit input number and returning a number.  The goal is to
	// determine which inputs will cause the value returned to be zero and to find the
	// largest such input.
	//
	// Analyzing the assembly by hand shows that it has a regular structure to it.  There
	// are 14 copies of the same block of code, parameterized differently.  Each block
	// processes a single digit of the inputted number starting with the most significant
	// digit, and the output of the block flows into the next block.  So they form a linked
	// list of evaluations.
	//
	// Here is an example block in higher-level pseudocode with hardcoded values replaced
	// with the parameters A, B, and C.
	//
	// 1:  w = input()
	// 2:  x = (z % 26) + B
	// 3:  y = 0 if (x == w) else 1
	// 4:  z = z/A
	// 5:  z = z * (25*y + 1) + (w + C)*y
	//
	// Analyzing the 14 different copies of this block shows that parameter B only ever
	// takes on values 1 and 26 and each value occurs in exactly half of the blocks.
	// Using this information as well as analyzing lines 4 and 5 help to provide insight
	// into what the algorithm is doing.
	//
	// When A=26, line 4 will remove some of the least significant bits of the previous
	// block's z.
	//
	// When x=input(), line 5 will shift the bits in the previous block's z to the left
	// and insert some new ones.
	//
	// Thinking of the variable z as a stream of data or a stack, we can now see that
	// when A=1 we are pushing data (w+C) into the stream, and when x=input() we are
	// removing the most recently pushed data from the stream.  This LIFO behavior makes
	// z behave very much like a stack.
	//
	// Rephrasing the above pseudocode with the stack paradigm we now have.
	//
	// 1:  x = peek(z) + B
	// 2:  if A == 26: pop(z)
	// 3:  if x != input(): push(z, input()+C)
	//
	// From this we can see that we don't have any control over when a pop operation
	// happens, but we can influence when a push happens, and the algorithm permits
	// a block to perform both a push and a pop.  Since at the conclusion of the
	// algorithm we want z to have the value of 0, we need to make sure that
	// every value pushed into z is also popped.  Also, because there are the same
	// number of blocks with B=1 as there are with B=26, we have to ensure that
	// blocks that perform a pop don't also perform a push.
	//
	// Analyzing the blocks from my specific input we can see:
	//
	// 1:   A=1   B=13   C=13   =>   pushes in1 + 13
	// 2:   A=1   B=11   C=10   =>   pushes in2 + 10
	// 3:   A=1   B=15   C=5    =>   pushes in3 + 5
	// 4:   A=26  B=-11  C=14   =>   pops   in3 + 5    need in4 == in3 + 5 - 11
	// 5:   A=1   B=14   C=5    =>   pushes in5 + 5
	// 6:   A=26  B=0    C=15   =>   pops   in5 + 5    need in6 == in5 + 5 + 0
	// 7:   A=1   B=12   C=4    =>   pushes in7 + 4
	// 8:   A=1   B=12   C=11   =>   pushes in8 + 11
	// 9:   A=1   B=14   C=1    =>   pushes in9 + 1
	// 10:  A=26  B=-6   C=15   =>   pops   in9 + 1    need in10 == in9 + 1 - 6
	// 11:  A=26  B=-10  C=12   =>   pops   in8 + 11   need in11 == in8 + 11 - 10
	// 12:  A=26  B=-12  C=8    =>   pops   in7 + 4    need in12 == in7 + 4 - 12
	// 13:  A=26  B=-3   C=14   =>   pops   in2 + 10   need in13 == in2 + 10 - 3
	// 14:  A=26  B=-5   C=9    =>   pops   in1 + 13   need in14 == in1 + 13 - 5
	//
	// From this we end up with an underspecified system of equations which will
	// result in many solutions.  Knowing that our inputs are single digits that
	// are never zero can be used to significantly narrow down the set of possible
	// solutions.
	//
	// in4 = in3 - 6   => in3 = [7, 8, 9]
	// in6 = in5 + 5   => in5 = [1, 2, 3, 4]
	// in10 = in9 - 5  => in9 = [6, 7, 8, 9]
	// in11 = in8 + 1  => in8 = [1, 2, 3, 4, 5, 6, 7, 8]
	// in12 = in7 - 8  => in7 = [9]
	// in13 = in2 + 7  => in2 = [1, 2]
	// in14 = in1 + 8  => in1 = [1]

	var solutions [][]int
	for _, in1 := range []int{1} {
		for _, in2 := range []int{1, 2} {
			for _, in3 := range []int{7, 8, 9} {
				in4 := in3 - 6
				for _, in5 := range []int{1, 2, 3, 4} {
					in6 := in5 + 5
					for _, in7 := range []int{9} {
						for _, in8 := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
							for _, in9 := range []int{6, 7, 8, 9} {
								in10 := in9 - 5
								in11 := in8 + 1
								in12 := in7 - 8
								in13 := in2 + 7
								in14 := in1 + 8

								// This series of digits should always be a solution, but
								// double check just in case.
								inputs := []int{in1, in2, in3, in4, in5, in6, in7, in8, in9, in10, in11, in12, in13, in14}
								if Run(inputs) == 0 {
									solutions = append(solutions, inputs)
								}
							}
						}
					}
				}
			}
		}
	}

	// The solutions are already sorted, just show the last entry
	solution := solutions[len(solutions)-1]
	for _, d := range solution {
		fmt.Print(d)
	}
	fmt.Println()
}

func Run(inputs []int) int {
	inp := func() int {
		i := inputs[0]
		inputs = inputs[1:]
		return i
	}

	eql := func(a, b int) int {
		if a == b {
			return 1
		}
		return 0
	}

	registers := map[string]int{
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	get := func(s string) int {
		if v, ok := registers[s]; ok {
			return v
		}
		return aoc.ParseInt(s)
	}

	for _, i := range InputToInstructions() {
		switch i.op {
		case "inp":
			registers[i.args[0]] = inp()
		case "add":
			registers[i.args[0]] = get(i.args[0]) + get(i.args[1])
		case "mul":
			registers[i.args[0]] = get(i.args[0]) * get(i.args[1])
		case "div":
			registers[i.args[0]] = get(i.args[0]) / get(i.args[1])
		case "mod":
			registers[i.args[0]] = get(i.args[0]) % get(i.args[1])
		case "eql":
			registers[i.args[0]] = eql(get(i.args[0]), get(i.args[1]))
		}
	}

	return registers["z"]
}

type Instruction struct {
	op   string
	args []string
}

func (i Instruction) String() string {
	return i.op + " " + strings.Join(i.args, " ")
}

func InputToInstructions() []Instruction {
	var instructions []Instruction
	for _, line := range aoc.InputToLines(2021, 24) {
		parts := strings.Split(line, " ")
		instructions = append(instructions, Instruction{
			op:   parts[0],
			args: parts[1:],
		})
	}

	return instructions
}
