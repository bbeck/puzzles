package main

import (
	"fmt"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	instructions := InputToInstructions()

	essences := make(map[string]int)
	for squire, ins := range instructions {
		essences[squire] = Run(ins)
	}

	squires := Keys(essences)
	sort.Slice(squires, func(i, j int) bool {
		return essences[squires[i]] > essences[squires[j]]
	})

	fmt.Println(strings.Join(squires, ""))
}

func Run(instructions string) int {
	var essence int
	power := 10

	for step := 0; step < 10; step++ {
		action := string(instructions[step%len(instructions)])

		switch action {
		case "+":
			power++
		case "-":
			power--
		}
		essence += power
	}

	return essence
}

func InputToInstructions() map[string]string {
	instructions := make(map[string]string)
	for _, line := range InputToLines() {
		lhs, rhs, _ := strings.Cut(line, ":")
		rhs = strings.ReplaceAll(rhs, ",", "")
		instructions[lhs] = rhs
	}

	return instructions
}
