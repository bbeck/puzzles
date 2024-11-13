package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

const (
	Width  = 25
	Height = 6

	Black       = 0
	White       = 1
	Transparent = 2
)

func main() {
	layers := InputToLayers()

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			switch Color(layers, x, y) {
			case Black:
				fmt.Print(" ")
			case White:
				fmt.Print("█")
			case Transparent:
				fmt.Print("!")
			}
		}
		fmt.Println()
	}
}

func Color(layers []Layer, x, y int) int {
	index := y*Width + x
	for _, layer := range layers {
		if layer[index] != Transparent {
			return layer[index]
		}
	}
	return Transparent
}

type Layer []int

func InputToLayers() []Layer {
	var digits []int
	for _, b := range puz.InputToString(2019, 8) {
		digits = append(digits, puz.ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, digits[start:start+Width*Height])
	}
	return layers
}
