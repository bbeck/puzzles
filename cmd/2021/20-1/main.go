package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	algorithm, image := InputToAlgorithmAndImage()

	for step := 0; step < 2; step++ {
		image = Enhance(image, algorithm, algorithm[0] == '#' && step%2 == 1)
	}

	var count int
	for _, v := range image {
		if v {
			count++
		}
	}
	fmt.Println(count)
}

func Show(image map[aoc.Point2D]bool, border bool) {
	var points []aoc.Point2D
	for p := range image {
		points = append(points, p)
	}

	minX, minY, maxX, maxY := aoc.GetBounds(points)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			q := aoc.Point2D{X: x, Y: y}

			var value bool
			if q.X < minX || q.Y < minY || q.X > maxX || q.Y > maxY {
				value = border
			} else {
				value = image[q]
			}

			if value {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func Enhance(image map[aoc.Point2D]bool, algorithm string, border bool) map[aoc.Point2D]bool {
	var points []aoc.Point2D
	for p := range image {
		points = append(points, p)
	}
	minX, minY, maxX, maxY := aoc.GetBounds(points)

	output := make(map[aoc.Point2D]bool)
	for x := minX - 2; x <= maxX+2; x++ {
		for y := minY - 2; y <= maxY+2; y++ {
			p := aoc.Point2D{X: x, Y: y}
			index := GetNumber(p, image, border)

			output[p] = algorithm[index] == '#'
		}
	}

	return output
}

func GetNumber(p aoc.Point2D, image map[aoc.Point2D]bool, border bool) int {
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
