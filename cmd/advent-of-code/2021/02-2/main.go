package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var position, depth, aim int
	for _, command := range InputToCommands() {
		switch command.Direction {
		case "forward":
			position += command.Distance
			depth += command.Distance * aim
		case "down":
			aim += command.Distance
		case "up":
			aim -= command.Distance
		}
	}
	fmt.Println(position * depth)
}

type Command struct {
	Direction string
	Distance  int
}

func InputToCommands() []Command {
	return puz.InputLinesTo(2021, 2, func(line string) Command {
		var command Command
		fmt.Sscanf(line, "%s %d", &command.Direction, &command.Distance)
		return command
	})
}
