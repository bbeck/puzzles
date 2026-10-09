package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	shapes, regions := InputToShapesAndRegions()

	var count int
	for _, region := range regions {
		var available = region.Area
		for p, c := range region.Presents {
			available -= c * shapes[p].Area
		}

		if available >= 0 {
			count++
		}
	}
	fmt.Println(count)
}

type Shape struct {
	Area int
}

type Region struct {
	Area     int
	Presents []int
}

func InputToShapesAndRegions() ([]Shape, []Region) {
	isShape := func() bool {
		for i := range 10 {
			if in.HasPrefix(fmt.Sprintf("%d:", i)) {
				return true
			}
		}
		return false
	}

	var shapes []Shape
	for isShape() {
		in.Line() // id

		chunk := in.ChunkS()
		grid := chunk.StringGrid2D()
		shapes = append(shapes, Shape{
			Area: grid.Width * grid.Height,
		})
	}

	var regions []Region
	for in.HasNext() {
		width, height := in.Int(), in.Int()

		var presents []int
		for range len(shapes) {
			presents = append(presents, in.Int())
		}

		regions = append(regions, Region{
			Area:     width * height,
			Presents: presents,
		})
	}

	return shapes, regions
}
