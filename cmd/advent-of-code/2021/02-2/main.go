package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
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
	return lib.InputLinesTo(func(line string) Command {
		var command Command
		fmt.Sscanf(line, "%s %d", &command.Direction, &command.Distance)
		return command
	})
}
