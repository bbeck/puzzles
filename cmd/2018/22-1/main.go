package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	cave := InputToCave(2018, 22)

	var risk int
	for _, r := range cave {
		risk += r
	}
	fmt.Printf("risk level: %d\n", risk)
}

type Cave map[aoc.Point2D]int

func (c Cave) String() string {
	minX, maxX := math.MaxInt64, math.MinInt64
	minY, maxY := math.MaxInt64, math.MinInt64

	for p := range c {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	var builder strings.Builder
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			switch c[aoc.Point2D{X: x, Y: y}] {
			case 0:
				builder.WriteString(".")
			case 1:
				builder.WriteString("=")
			case 2:
				builder.WriteString("|")
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func InputToCave(year, day int) Cave {
	depth, target := InputToDepthAndTarget(year, day)

	var gi, el func(p aoc.Point2D) int

	gis := make(map[aoc.Point2D]int)
	gis[aoc.Point2D{}] = 0
	gis[target] = 0
	gi = func(p aoc.Point2D) int {
		if index, ok := gis[p]; ok {
			return index
		}

		if p.Y == 0 {
			return p.X * 16807
		}

		if p.X == 0 {
			return p.Y * 48271
		}

		index := el(p.Left()) * el(p.Up())
		gis[p] = index
		return index
	}

	el = func(p aoc.Point2D) int {
		return (depth + gi(p)) % 20183
	}

	cave := make(Cave)
	for y := 0; y <= target.Y; y++ {
		for x := 0; x <= target.X; x++ {
			p := aoc.Point2D{X: x, Y: y}
			cave[p] = el(p) % 3
		}
	}
	return cave
}

func InputToDepthAndTarget(year, day int) (int, aoc.Point2D) {
	var depth int
	var target aoc.Point2D
	for _, line := range aoc.InputToLines(year, day) {
		if strings.HasPrefix(line, "depth:") {
			depth = aoc.ParseInt(line[7:])
		} else if strings.HasPrefix(line, "target:") {
			parts := strings.Split(line[8:], ",")
			target.X = aoc.ParseInt(parts[0])
			target.Y = aoc.ParseInt(parts[1])
		} else {
			log.Fatalf("unrecognized line: %s", line)
		}
	}

	return depth, target
}
