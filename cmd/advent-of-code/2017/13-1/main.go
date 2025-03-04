package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib/in"
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
	return in.LinesToS(func(in in.Scanner[Layer]) Layer {
		return Layer{Depth: in.Int(), Range: in.Int()}
	})
}
