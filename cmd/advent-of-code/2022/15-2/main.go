package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

const N = 4000000

func main() {
	sensors := InputToSensors()

	// Each sensor forms a diamond with a specific radius.  If we assume the
	// point we're looking for is on the interior of the region of interest then
	// the point must be one unit outside a diamond.  If it were more than one
	// unit outside a diamonds then it wouldn't be a unique solution.
	//
	// Each diamond has two sides that are a line segment with slope +1 and two
	// sides that are a line segment with slope -1.  We can "nudge" these 1 unit
	// outside the diamond to track the perimeter.  The point we're interested in
	// will lie at the intersection of 4 of these line segments -- two of the +1
	// sloped ones and two of the -1 sloped ones.

	// Compute the y-intercept of the positively and negatively sloped lines.
	// We know that y = mx+b, so for +1 sloped lines b = y-x and for -1 sloped
	// lines b = y+x.
	var pos, neg []int
	for _, sensor := range sensors {
		// We're looking just outside the diamond, so we bump the radius by 1.
		x1, x2 := sensor.X-sensor.Radius-1, sensor.X+sensor.Radius+1
		y := sensor.Y

		pos = append(pos, y-x1)
		pos = append(pos, y-x2)
		neg = append(neg, y+x1)
		neg = append(neg, y+x2)
	}

	// Take a lines from each set and see if they intersect.  The two lines have
	// equations y = x+bp and y = -x+bn, thus we know x = (bn-bp)/2 and
	// y = (bn+bp)/2.
	var p lib.Point2D
outer:
	for i := 0; i < len(pos); i++ {
		for j := 0; j < len(neg); j++ {
			p = lib.Point2D{
				X: (neg[j] - pos[i]) / 2,
				Y: (neg[j] + pos[i]) / 2,
			}

			// Make sure our point of intersection is within the bounding box.
			if p.X < 0 || N < p.X || p.Y < 0 || N < p.Y {
				continue
			}

			// Also make sure this point doesn't lie within any of the sensor areas.
			ok := true
			for _, sensor := range sensors {
				if sensor.ManhattanDistance(p) < sensor.Radius {
					ok = false
					break
				}
			}

			if ok {
				break outer
			}
		}
	}

	fmt.Println(4000000*p.X + p.Y)
}

type Sensor struct {
	lib.Point2D
	Radius int
}

func InputToSensors() []Sensor {
	return lib.InputLinesTo(func(line string) Sensor {
		var sensor, beacon lib.Point2D
		fmt.Sscanf(
			line,
			"Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y,
			&beacon.X, &beacon.Y,
		)

		return Sensor{
			Point2D: sensor,
			Radius:  sensor.ManhattanDistance(beacon),
		}
	})
}
