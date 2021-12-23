package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"log"
)

func main() {
	// Keep track of which cubes have all of their lights on.
	var on []Cube

	for _, command := range InputToCommands() {
		var next []Cube

		// Remove this command's cube from any cube that's already on.  We
		// do this regardless of whether this is an "off" command or an "on"
		// command.  In the "off" case, the lights aren't supposed to be on,
		// so removing them is the correct thing to do.  In the "on" case, since
		// we're about to add this command's cube to the on list we'd be double
		// counting any lights that overlap.
		for _, other := range on {
			for _, child := range other.Subtract(command.Cube) {
				next = append(next, child)
			}
		}

		// If this cube's lights should be on, add it to the list of on cubes.
		// This is safe to do because we removed all overlapping lights that
		// were on so we wont' be double counting.
		if command.State == "on" {
			next = append(next, command.Cube)
		}

		on = next
	}

	var sum int
	for _, c := range on {
		sum += c.Count()
	}
	fmt.Println(sum)
}

type Cube struct {
	x1, x2, y1, y2, z1, z2 int
}

func (c Cube) Count() int {
	return (c.x2 - c.x1) * (c.y2 - c.y1) * (c.z2 - c.z1)
}

func (c Cube) Subtract(o Cube) []Cube {
	// Cubes only overlap if they overlap along all axes.  A failure on any axis means
	// there's no overlap.
	if c.x2 < o.x1 || o.x2 < c.x1 || c.y2 < o.y1 || o.y2 < c.y1 || c.z2 < o.z1 || o.z2 < c.z1 {
		// When there's no overlap, the subtraction does nothing to the original cube.
		return []Cube{c}
	}

	o = Cube{
		x1: aoc.MinInt(aoc.MaxInt(c.x1, o.x1), c.x2),
		x2: aoc.MinInt(aoc.MaxInt(c.x1, o.x2), c.x2),
		y1: aoc.MinInt(aoc.MaxInt(c.y1, o.y1), c.y2),
		y2: aoc.MinInt(aoc.MaxInt(c.y1, o.y2), c.y2),
		z1: aoc.MinInt(aoc.MaxInt(c.z1, o.z1), c.z2),
		z2: aoc.MinInt(aoc.MaxInt(c.z1, o.z2), c.z2),
	}

	var cubes []Cube
	cube := Cube{c.x1, o.x1, c.y1, c.y2, c.z1, c.z2}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}
	cube = Cube{o.x2, c.x2, c.y1, c.y2, c.z1, c.z2}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}
	cube = Cube{o.x1, o.x2, c.y1, o.y1, c.z1, c.z2}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}
	cube = Cube{o.x1, o.x2, o.y2, c.y2, c.z1, c.z2}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}
	cube = Cube{o.x1, o.x2, o.y1, o.y2, c.z1, o.z1}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}
	cube = Cube{o.x1, o.x2, o.y1, o.y2, o.z2, c.z2}
	if cube.Count() > 0 {
		cubes = append(cubes, cube)
	}

	return cubes
}

type Command struct {
	State string
	Cube  Cube
}

func InputToCommands() []Command {
	var commands []Command
	for _, line := range aoc.InputToLines(2021, 22) {
		var state string
		var minX, maxX, minY, maxY, minZ, maxZ int

		if _, err := fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &state, &minX, &maxX, &minY, &maxY, &minZ, &maxZ); err != nil {
			log.Fatal(err)
		}

		commands = append(commands, Command{
			State: state,
			Cube: Cube{
				x1: aoc.MinInt(minX, maxX),
				x2: aoc.MaxInt(minX, maxX) + 1,
				y1: aoc.MinInt(minY, maxY),
				y2: aoc.MaxInt(minY, maxY) + 1,
				z1: aoc.MinInt(minZ, maxZ),
				z2: aoc.MaxInt(minZ, maxZ) + 1,
			},
		})
	}

	return commands
}
