package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

const (
	Width  int = 25
	Height     = 6
)

func main() {
	layers := InputToLayers(2019, 8)
	output := Render(layers)
	fmt.Println("output:")
	fmt.Println(output)
}

func Render(layers []Layer) Layer {
	var output Layer = make([]int, Width*Height)

pixel:
	for i := 0; i < Width*Height; i++ {
		for _, layer := range layers {
			output[i] = layer[i]
			if output[i] != 2 {
				continue pixel
			}
		}
	}

	return output
}

type Layer []int

func (l Layer) String() string {
	var builder strings.Builder
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			switch l[y*Width+x] {
			case 0:
				builder.WriteString(" ")
			case 1:
				builder.WriteString("█")
			case 2:
				builder.WriteString("!")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func InputToLayers(year, day int) []Layer {
	var digits []int
	for _, b := range aoc.InputToString(year, day) {
		digits = append(digits, aoc.ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, digits[start:start+Width*Height])
	}
	return layers
}
