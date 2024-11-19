package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	var on []lib.Cube
	for _, c := range InputToCommands() {
		// Remove this command's cube from any cube that's already on so that we
		// don't double count the "on" volume.
		var next []lib.Cube
		for _, other := range on {
			for _, child := range other.Subtract(c.Cube) {
				next = append(next, child)
			}
		}
		on = next

		// Now that this cube doesn't overlap with any others, add it to the "on"
		// volume if the command is to turn these lights on.
		if c.State == "on" {
			on = append(on, c.Cube)
		}
	}

	var sum int
	for _, c := range on {
		sum += c.Volume()
	}
	fmt.Println(sum)
}

type Command struct {
	State string
	Cube  lib.Cube
}

func InputToCommands() []Command {
	return lib.InputLinesTo(func(line string) Command {
		var c Command
		fmt.Sscanf(line, "%s x=%d..%d,y=%d..%d,z=%d..%d", &c.State,
			&c.Cube.MinX, &c.Cube.MaxX, &c.Cube.MinY,
			&c.Cube.MaxY, &c.Cube.MinZ, &c.Cube.MaxZ,
		)

		// Because the endpoints of our ranges are inclusive we always grow the
		// cube by one unit along each dimension.
		c.Cube.MaxX += 1
		c.Cube.MaxY += 1
		c.Cube.MaxZ += 1

		return c
	})
}
