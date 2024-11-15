package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	algorithm, image := InputToAlgorithmAndImage()
	border := false

	for n := 0; n < 50; n++ {
		image, border = Enhance(image, algorithm, border)
	}

	fmt.Println(len(image))
}

func Enhance(image puz.Set[puz.Point2D], algorithm []bool, border bool) (puz.Set[puz.Point2D], bool) {
	tl, br := puz.GetBounds(image.Entries())

	index := func(x, y int) int {
		var n int
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				var value bool
				if tl.X <= x+dx && x+dx <= br.X && tl.Y <= y+dy && y+dy <= br.Y {
					value = image.Contains(puz.Point2D{X: x + dx, Y: y + dy})
				} else {
					value = border
				}

				n <<= 1
				if value {
					n |= 1
				}
			}
		}
		return n
	}

	var next puz.Set[puz.Point2D]
	for y := tl.Y - 2; y <= br.Y+2; y++ {
		for x := tl.X - 2; x <= br.X+2; x++ {
			if algorithm[index(x, y)] {
				next.Add(puz.Point2D{X: x, Y: y})
			}
		}
	}

	// Toggle the border if the all 0's rule and all 1's rule say to
	if toggle := algorithm[0] && !algorithm[len(algorithm)-1]; toggle {
		border = !border
	}
	return next, border
}

func InputToAlgorithmAndImage() ([]bool, puz.Set[puz.Point2D]) {
	lines := puz.InputToLines()

	var algorithm []bool
	for _, c := range lines[0] {
		algorithm = append(algorithm, c == '#')
	}

	var image puz.Set[puz.Point2D]
	for y := 2; y < len(lines); y++ {
		for x, c := range lines[y] {
			if c == '#' {
				image.Add(puz.Point2D{X: x, Y: y - 2})
			}
		}
	}

	return algorithm, image
}
