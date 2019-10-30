package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	layers := InputToLayers(2017, 13)

	var delay int
	for delay = 0; ; delay++ {
		if Simulate(layers) {
			break
		}

		for _, layer := range layers {
			layer.Step()
		}
	}

	fmt.Printf("delay: %d\n", delay)
}

func Simulate(initial []*Layer) bool {
	layers := make([]*Layer, len(initial))
	for i, layer := range initial {
		layers[i] = &Layer{
			id:        layer.id,
			depth:     layer.depth,
			position:  layer.position,
			direction: layer.direction,
		}
	}

	for tm := 0; tm <= layers[len(layers)-1].id; tm++ {
		for _, layer := range layers {
			if layer.id == tm && layer.position == 0 {
				return false
			}

			layer.Step()
		}
	}

	return true
}

type Layer struct {
	id        int
	depth     int
	position  int
	direction string
}

func (l *Layer) Step() {
	switch l.direction {
	case "U":
		if l.position == 0 {
			l.position = 1
			l.direction = "D"
			return
		}

		l.position--

	case "D":
		if l.position == l.depth-1 {
			l.position = l.depth - 2
			l.direction = "U"
			return
		}

		l.position++
	}
}

func InputToLayers(year, day int) []*Layer {
	var layers []*Layer
	for _, line := range aoc.InputToLines(year, day) {
		parts := strings.Split(strings.ReplaceAll(line, ":", ""), " ")
		id := aoc.ParseInt(parts[0])
		depth := aoc.ParseInt(parts[1])

		layers = append(layers, &Layer{id: id, depth: depth, direction: "D"})
	}

	return layers
}
