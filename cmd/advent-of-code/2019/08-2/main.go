package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
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

	for y := range Height {
		for x := range Width {
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
	for _, b := range in.String() {
		digits = append(digits, ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, digits[start:start+Width*Height])
	}
	return layers
}
