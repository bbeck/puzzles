package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/advent-of-code/puz"
)

const (
	Width  = 25
	Height = 6
)

func main() {
	layers := InputToLayers()

	counters := make([]puz.FrequencyCounter[int], len(layers))
	for i, layer := range layers {
		for _, b := range layer {
			counters[i].Add(b)
		}
	}

	sort.Slice(counters, func(i, j int) bool {
		return counters[i].GetCount(0) < counters[j].GetCount(0)
	})

	fmt.Println(counters[0].GetCount(1) * counters[0].GetCount(2))
}

type Layer []int

func InputToLayers() []Layer {
	var digits []int
	for _, b := range puz.InputToString() {
		digits = append(digits, puz.ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, digits[start:start+Width*Height])
	}
	return layers
}
