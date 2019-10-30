package main

import (
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	bots := InputToNanobots(2018, 23)

	var largest Nanobot
	for _, bot := range bots {
		if bot.radius > largest.radius {
			largest = bot
		}
	}

	var count int
	for _, bot := range bots {
		if largest.InRange(bot.location) {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

type Nanobot struct {
	location Point3D
	radius   int
}

// Determine if a point is in range of this nanobot.
func (n Nanobot) InRange(p Point3D) bool {
	dx := n.location.X - p.X
	if dx < 0 {
		dx = -dx
	}
	dy := n.location.Y - p.Y
	if dy < 0 {
		dy = -dy
	}
	dz := n.location.Z - p.Z
	if dz < 0 {
		dz = -dz
	}

	return dx+dy+dz < n.radius
}

type Point3D struct {
	X, Y, Z int
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}

func InputToNanobots(year, day int) []Nanobot {
	var nanobots []Nanobot
	for _, line := range aoc.InputToLines(year, day) {
		var x, y, z, r int
		if _, err := fmt.Sscanf(line, "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r); err != nil {
			log.Fatalf("unable to parse line: %s", line)
		}

		nanobots = append(nanobots, Nanobot{Point3D{x, y, z}, r})
	}

	return nanobots
}
