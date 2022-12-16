package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
	"math"
)

func main() {
	sensors := InputToSensors()

	minX := math.MaxInt
	maxX := 0
	for _, sensor := range sensors {
		distance := sensor.ManhattanDistance(sensor.Beacon)
		// fmt.Println(sensor, sensor.Beacon, distance)
		minX = aoc.Min(minX, sensor.X-distance)
		maxX = aoc.Max(maxX, sensor.X+distance)
	}
	fmt.Println("minX:", minX, "maxX:", maxX)

	var count int
	for x := minX; x <= maxX; x++ {
		if !InRange(sensors, x, 2000000) {
			count++
		}
	}
	fmt.Println(count)
}

func InRange(sensors []Sensor, x, y int) bool {
	p := aoc.Point2D{X: x, Y: y}
	for _, sensor := range sensors {
		if sensor.Point2D == p || sensor.Beacon == p {
			return true
		}
	}

	for _, sensor := range sensors {
		d1 := sensor.ManhattanDistance(sensor.Beacon)
		d2 := sensor.ManhattanDistance(p)
		if d2 <= d1 {
			return false
		}
	}
	return true
}

type Sensor struct {
	aoc.Point2D
	Beacon aoc.Point2D
}

func InputToSensors() []Sensor {
	return aoc.InputLinesTo(2022, 15, func(line string) (Sensor, error) {
		var sensor Sensor
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y, &sensor.Beacon.X, &sensor.Beacon.Y)
		return sensor, err
	})
}
