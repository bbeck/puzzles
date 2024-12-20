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

func Run(instructions []string) int {
	N := len(instructions)

	var essence int
	for power, step := 10, 0; step < 10; step++ {
		action := instructions[step%N]

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

func InputToInstructions() map[string][]string {
	instructions := make(map[string][]string)
	for _, line := range InputToLines() {
		line = strings.ReplaceAll(line, ":", " ")
		line = strings.ReplaceAll(line, ",", " ")
		fields := strings.Fields(line)

		instructions[fields[0]] = fields[1:]
	}

	return instructions
}
