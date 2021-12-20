package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	algorithm, image := InputToAlgorithmAndImage()

	// The border is infinite, so depending on the value of the algorithm when you have
	// a solid block of '.' or '#' it might toggle for each step.
	toggles := algorithm[0] == '#' && algorithm[511] == '.'

	for step := 0; step < 50; step++ {
		image = Enhance(image, algorithm, toggles && step%2 == 1)
	}

	var count int
	for _, v := range image {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

func GetBounds(image map[aoc.Point2D]bool) (int, int, int, int) {
	var points []aoc.Point2D
	for p := range image {
		points = append(points, p)
	}
	return aoc.GetBounds(points)
}

func Enhance(image map[aoc.Point2D]bool, algorithm string, border bool) map[aoc.Point2D]bool {
	minX, minY, maxX, maxY := GetBounds(image)

	output := make(map[aoc.Point2D]bool)
	for x := minX - 2; x <= maxX+2; x++ {
		for y := minY - 2; y <= maxY+2; y++ {
			p := aoc.Point2D{X: x, Y: y}
			index := GetIndex(p, image, border)

			output[p] = algorithm[index] == '#'
		}
	}

	return output
}

func GetIndex(p aoc.Point2D, image map[aoc.Point2D]bool, border bool) int {
	var n int
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			q := aoc.Point2D{X: p.X + dx, Y: p.Y + dy}

			var value, ok bool
			if value, ok = image[q]; !ok {
				value = border
			}

			n = n << 1
			if value {
				n |= 1
			}
		}
	}
	return n
}

func InputToAlgorithmAndImage() (string, map[aoc.Point2D]bool) {
	lines := aoc.InputToLines(2021, 20)

	algorithm := lines[0]

	image := make(map[aoc.Point2D]bool)
	for y := 2; y < len(lines); y++ {
		for x, c := range lines[y] {
			if c == '#' {
				image[aoc.Point2D{X: x, Y: y - 2}] = true
			} else {
				image[aoc.Point2D{X: x, Y: y - 2}] = false
			}
		}
	}

	return algorithm, image
}
