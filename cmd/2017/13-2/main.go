package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	layers := InputToLayers()

	var delay int
outer:
	for {
		delay++

		for _, layer := range layers {
			if Position(layer.Range, delay+layer.Depth) == 0 {
				continue outer
			}
		}

		break
	}

	fmt.Println(delay)
}

func Position(r int, tm int) int {
	// The scanner moves in discrete steps between 0 and r-1 and then back again.
	// This means that it's period is 2*r-2.  Using this we can directly compute
	// where it's located at any point in time.
	period := 2*r - 2
	x := tm % period
	if x >= r {
		x = period - x
	}
	return x
}

type Layer struct {
	Depth int
	Range int
}

func InputToLayers() []Layer {
	return aoc.InputLinesTo(2017, 13, func(line string) (Layer, error) {
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		return Layer{
			Depth: aoc.ParseInt(fields[0]),
			Range: aoc.ParseInt(fields[1]),
		}, nil
	})
}
