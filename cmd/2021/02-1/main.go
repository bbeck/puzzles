package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var position, depth int
	for _, command := range InputToCommands() {
		switch command.Direction {
		case "forward":
			position += command.Distance
		case "down":
			depth += command.Distance
		case "up":
			depth -= command.Distance
		}
	}
	fmt.Println(position * depth)
}

type Command struct {
	Direction string
	Distance  int
}

func InputToCommands() []Command {
	return aoc.InputLinesTo(2021, 2, func(line string) (Command, error) {
		var command Command
		_, err := fmt.Sscanf(line, "%s %d", &command.Direction, &command.Distance)
		return command, err
	})
}
