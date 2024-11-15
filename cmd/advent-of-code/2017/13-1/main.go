package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	var severity int
	for _, layer := range InputToLayers() {
		if Position(layer.Range, layer.Depth) == 0 {
			severity += layer.Depth * layer.Range
		}
	}
	fmt.Println(severity)
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
	return puz.InputLinesTo(func(line string) Layer {
		line = strings.ReplaceAll(line, ":", "")
		fields := strings.Fields(line)

		return Layer{
			Depth: puz.ParseInt(fields[0]),
			Range: puz.ParseInt(fields[1]),
		}
	})
}
