package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"sort"
)

const N = 4000000

func main() {
	sensors := InputToSensors()

	for y := 0; y <= N; y++ {
		var ranges []Range
		for _, s := range sensors {
			minX, maxX, ok := s.Range(y)
			if !ok {
				continue
			}

			ranges = append(ranges, Range{Min: minX, Max: maxX})
		}

		sort.Slice(ranges, func(i, j int) bool {
			if ranges[i].Min != ranges[j].Min {
				return ranges[i].Min < ranges[j].Min
			}
			return ranges[i].Max < ranges[j].Max
		})

		x := 0
		for _, r := range ranges {
			if x < r.Min {
				break
			}

			x = aoc.Max(x, r.Max)
		}

		if x < N {
			fmt.Println("y:", y, "x:", x, (x+1)*4000000+y)
			break
		}
	}
}

type Range struct {
	Min, Max int
}

type Sensor struct {
	aoc.Point2D
	Distance int
}

func (s Sensor) Range(y int) (int, int, bool) {
	// does this sensor see this y value?
	if s.Y-s.Distance > y || y > s.Y+s.Distance {
		return 0, 0, false
	}

	used := aoc.Abs(s.Y - y)
	minX := aoc.Max(0, s.X-(s.Distance-used))
	maxX := aoc.Min(N, s.X+(s.Distance-used))
	return minX, maxX, true
}

func InputToSensors() []Sensor {
	return aoc.InputLinesTo(2022, 15, func(line string) (Sensor, error) {
		var sensor Sensor
		var p aoc.Point2D
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y, &p.X, &p.Y)
		sensor.Distance = p.ManhattanDistance(sensor.Point2D)
		return sensor, err
	})
}
