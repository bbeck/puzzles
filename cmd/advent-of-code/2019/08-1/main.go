package main

import (
	"fmt"
	"sort"

	"github.com/bbeck/puzzles/lib"
)

const (
	Width  = 25
	Height = 6
)

func main() {
	layers := InputToLayers()

	counters := make([]lib.FrequencyCounter[int], len(layers))
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
	for _, b := range lib.InputToString() {
		digits = append(digits, lib.ParseInt(string(b)))
	}

	var layers []Layer
	for start := 0; start < len(digits); start += Width * Height {
		layers = append(layers, digits[start:start+Width*Height])
	}
	return layers
}
