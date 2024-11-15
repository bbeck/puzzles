package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"math"
	"sort"
)

const Y = 2000000

func main() {
	sensors := InputToSensors()

	var ranges []Range
	for _, s := range sensors {
		if r := s.GetRange(Y); r != nil {
			ranges = append(ranges, *r)
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	var count int
	x := math.MinInt
	for _, r := range ranges {
		if x > r.Max {
			continue
		}
		if x < r.Min {
			x = r.Min
		}

		// In theory this will count spaces that contain a beacon making our count
		// too high.  But for some reason in practice it doesn't?
		count += r.Max - x
		x = r.Max
	}
	fmt.Println(count)
}

type Range struct {
	Min, Max int
}

type Sensor struct {
	puz.Point2D
	Beacon puz.Point2D
}

func (s Sensor) GetRange(y int) *Range {
	distance := s.ManhattanDistance(s.Beacon) - puz.Abs(s.Y-y)
	if distance < 0 {
		return nil
	}
	return &Range{
		Min: s.X - distance,
		Max: s.X + distance,
	}
}

func InputToSensors() []Sensor {
	return puz.InputLinesTo(2022, 15, func(line string) Sensor {
		var sensor, beacon puz.Point2D
		fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y,
			&beacon.X, &beacon.Y,
		)

		return Sensor{
			Point2D: sensor,
			Beacon:  beacon,
		}
	})
}
