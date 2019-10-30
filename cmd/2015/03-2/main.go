package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	santa := aoc.Point2D{0, 0}
	robot := aoc.Point2D{0, 0}

	seen := map[aoc.Point2D]int{
		santa: 1,
	}

	bs := aoc.InputToBytes(2015, 3)
	for i := 0; i < len(bs); i += 2 {
		switch bs[i] {
		case '^':
			santa = santa.North()
		case '<':
			santa = santa.West()
		case '>':
			santa = santa.East()
		case 'v':
			santa = santa.South()
		default:
			log.Fatalf("unrecognized location: %s", string(bs[i]))
		}

		seen[santa]++
	}

	for i := 1; i < len(bs); i += 2 {
		switch bs[i] {
		case '^':
			robot = robot.North()
		case '<':
			robot = robot.West()
		case '>':
			robot = robot.East()
		case 'v':
			robot = robot.South()
		default:
			log.Fatalf("unrecognized location: %s", string(bs[i]))
		}

		seen[robot]++
	}

	fmt.Printf("number of locations: %d\n", len(seen))
}
