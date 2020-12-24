package main

import (
	"fmt"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := make(map[HexPoint]bool)
	for _, p := range InputToCoordinates(2020, 24) {
		grid[p] = !grid[p]
	}

	var sum int
	for _, value := range grid {
		if value {
			sum++
		}
	}
	fmt.Println(sum)
}

func InputToCoordinates(year, day int) []HexPoint {
	var coordinates []HexPoint
	for _, line := range aoc.InputToLines(year, day) {
		var p HexPoint

		for i := 0; i < len(line); {
			switch {
			case line[i:i+1] == "e":
				p = p.East()
				i += 1
			case i+2 <= len(line) && line[i:i+2] == "se":
				p = p.SouthEast()
				i += 2
			case i+2 <= len(line) && line[i:i+2] == "sw":
				p = p.SouthWest()
				i += 2
			case line[i:i+1] == "w":
				p = p.West()
				i += 1
			case i+2 <= len(line) && line[i:i+2] == "nw":
				p = p.NorthWest()
				i += 2
			case i+2 <= len(line) && line[i:i+2] == "ne":
				p = p.NorthEast()
				i += 2
			}
		}

		coordinates = append(coordinates, p)
	}

	return coordinates
}

type HexPoint struct {
	X, Y, Z int
}

func (p HexPoint) East() HexPoint {
	return HexPoint{X: p.X + 1, Y: p.Y - 1, Z: p.Z}
}

func (p HexPoint) SouthEast() HexPoint {
	return HexPoint{X: p.X, Y: p.Y - 1, Z: p.Z + 1}
}

func (p HexPoint) SouthWest() HexPoint {
	return HexPoint{X: p.X - 1, Y: p.Y, Z: p.Z + 1}
}

func (p HexPoint) West() HexPoint {
	return HexPoint{X: p.X - 1, Y: p.Y + 1, Z: p.Z}
}

func (p HexPoint) NorthWest() HexPoint {
	return HexPoint{X: p.X, Y: p.Y + 1, Z: p.Z - 1}
}

func (p HexPoint) NorthEast() HexPoint {
	return HexPoint{X: p.X + 1, Y: p.Y, Z: p.Z - 1}
}
