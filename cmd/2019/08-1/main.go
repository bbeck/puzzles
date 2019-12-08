package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

const (
	Width  int = 25
	Height     = 6
)

func main() {
	layers := InputToLayers(2019, 8)

	count := func(l Layer, digit int) int {
		var count int
		for _, p := range l.pixels {
			if p == digit {
				count++
			}
		}
		return count
	}

	var bestCount = Width * Height
	var bestLayer Layer
	for _, layer := range layers {
		numZeroes := count(layer, 0)
		if numZeroes < bestCount {
			bestCount = numZeroes
			bestLayer = layer
		}
	}

	fmt.Printf("product: %d\n", count(bestLayer, 1)*count(bestLayer, 2))
}

type Layer struct {
	pixels []int
}

func InputToLayers(year, day int) []Layer {
	var digits []int
	for _, b := range aoc.InputToString(year, day) {
		digits = append(digits, aoc.ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, Layer{pixels: digits[start : start+Width*Height]})
	}
	return layers
}
