package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var position, depth int
	for _, command := range InputToCommands() {
		switch command.direction {
		case "forward":
			position += command.distance
		case "down":
			depth += command.distance
		case "up":
			depth -= command.distance
		}
	}
	fmt.Println(position * depth)
}

type Command struct {
	direction string
	distance  int
}

func InputToCommands() []Command {
	var commands []Command
	for _, line := range aoc.InputToLines(2021, 2) {
		var command Command
		if _, err := fmt.Sscanf(line, "%s %d", &command.direction, &command.distance); err != nil {
			log.Fatal(err)
		}

		commands = append(commands, command)
	}

	return commands
}
