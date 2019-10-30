package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	layers := InputToLayers(2017, 13)

	var severity int
	for tm := 0; tm <= layers[len(layers)-1].id; tm++ {
		for _, layer := range layers {
			if layer.id == tm {
				if layer.position == 0 {
					severity += tm * layer.depth
				}
			}

			layer.Step()
		}
	}

	fmt.Printf("severity: %d\n", severity)
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
