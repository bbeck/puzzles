package main

import (
	"fmt"
	"log"
	"math"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	points := InputToPoints(2018, 10)

	width, height := math.MaxInt64, math.MaxInt64
	var tm int
	for tm = 1; tm < 100000; {
		Forward(points)
		tm++

		w, h := Dimensions(points)
		if w > width && h > height {
			break
		}

		width = w
		height = h
	}

	// We now know the point in time when the points are closest to each other.
	// Let's rewind a few steps and start showing the space.
	N := 10
	for i := 0; i < N; i++ {
		Backward(points)
		tm--
	}

	for i := 0; i < 2*N; i++ {
		fmt.Printf("tm: %d\n", tm)
		tm++

		Forward(points)

		Show(points)
		fmt.Println()
	}
}

type Point struct {
	position aoc.Point2D
	velocity aoc.Point2D
}

func Forward(ps []*Point) {
	for _, p := range ps {
		p.position.X += p.velocity.X
		p.position.Y += p.velocity.Y
	}
}

func Backward(ps []*Point) {
	for _, p := range ps {
		p.position.X -= p.velocity.X
		p.position.Y -= p.velocity.Y
	}
}

func Dimensions(ps []*Point) (int, int) {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	grid := make(map[aoc.Point2D]bool)
	for _, p := range ps {
		grid[p.position] = true

		if p.position.X < minX {
			minX = p.position.X
		}
		if p.position.X > maxX {
			maxX = p.position.X
		}
		if p.position.Y < minY {
			minY = p.position.Y
		}
		if p.position.Y > maxY {
			maxY = p.position.Y
		}
	}

	return maxX - minX, maxY - minY
}

func Show(ps []*Point) {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	grid := make(map[aoc.Point2D]bool)
	for _, p := range ps {
		grid[p.position] = true

		if p.position.X < minX {
			minX = p.position.X
		}
		if p.position.X > maxX {
			maxX = p.position.X
		}
		if p.position.Y < minY {
			minY = p.position.Y
		}
		if p.position.Y > maxY {
			maxY = p.position.Y
		}
	}

	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			if grid[aoc.Point2D{X: x, Y: y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func InputToPoints(year, day int) []*Point {
	var points []*Point
	for _, line := range aoc.InputToLines(year, day) {
		var x, y, dx, dy int
		if _, err := fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy); err != nil {
			log.Fatalf("unable to parse point: %s", line)
		}

		points = append(points, &Point{
			position: aoc.Point2D{X: x, Y: y},
			velocity: aoc.Point2D{X: dx, Y: dy},
		})
	}

	return points
}
