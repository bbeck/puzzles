package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := InputToGrid(2017, 14)

	sets := make(map[aoc.Point2D]*aoc.DisjointSet)
	for p := range grid {
		sets[p] = aoc.NewDisjointSet(p)
	}

	// Union the adjacent on points together
	for p := range grid {
		if !grid[p] {
			continue
		}

		if grid[p.Up()] {
			sets[p].Union(sets[p.Up()])
		}

		if grid[p.Down()] {
			sets[p].Union(sets[p.Down()])
		}

		if grid[p.Left()] {
			sets[p].Union(sets[p.Left()])
		}

		if grid[p.Right()] {
			sets[p].Union(sets[p.Right()])
		}
	}

	seen := make(map[*aoc.DisjointSet]bool)
	for y := 0; y < 128; y++ {
		for x := 0; x < 128; x++ {
			p := aoc.Point2D{X: x, Y: y}
			s := sets[aoc.Point2D{X: x, Y: y}].Find()

			if grid[p] && !seen[s] {
				seen[s] = true
			}
		}
	}

	fmt.Printf("num regions: %d\n", len(seen))
}

func InputToGrid(year, day int) map[aoc.Point2D]bool {
	grid := make(map[aoc.Point2D]bool)

	input := aoc.InputToString(year, day)
	for y := 0; y < 128; y++ {
		for chunk, c := range KnotHash(fmt.Sprintf("%s-%d", input, y)) {
			switch c {
			case '0':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case '1':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case '2':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case '3':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case '4':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case '5':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case '6':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case '7':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case '8':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case '9':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case 'a':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case 'b':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case 'c':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case 'd':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = false
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true

			case 'e':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = false

			case 'f':
				grid[aoc.Point2D{X: chunk*4 + 0, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 1, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 2, Y: y}] = true
				grid[aoc.Point2D{X: chunk*4 + 3, Y: y}] = true
			}
		}
	}

	return grid
}

func KnotHash(s string) string {
	bs := []byte(s)
	bs = append(bs, []byte{17, 31, 73, 47, 23}...)

	ring := &Ring{data: make([]int, 0)}
	for i := 0; i <= 255; i++ {
		ring.Append(i)
	}

	for round := 0; round < 64; round++ {
		for _, length := range bs {
			ring.Twist(int(length))
		}
	}

	return ToHex(DenseHash(ring.data))
}

func DenseHash(data []int) []int {
	var hash []int
	for segment := 0; segment < 16; segment++ {
		var b int
		for i := 0; i < 16; i++ {
			b = b ^ data[segment*16+i]
		}

		hash = append(hash, b)
	}

	return hash
}

func ToHex(data []int) string {
	var s string
	for _, i := range data {
		s = s + fmt.Sprintf("%02x", i)
	}

	return s
}

type Ring struct {
	data    []int
	current int
	skip    int
}

func (r *Ring) Append(n int) {
	r.data = append(r.data, n)
}

func (r *Ring) Twist(length int) {
	N := len(r.data)
	swap := func(a, b int) {
		a = (a + N) % N
		b = (b + N) % N
		r.data[a], r.data[b] = r.data[b], r.data[a]
	}

	for i := 0; i < length/2; i++ {
		swap(r.current+i, r.current+length-i-1)
	}

	r.current = (r.current + length + N + r.skip) % N
	r.skip++
}

func (r *Ring) String() string {
	var builder strings.Builder
	for i := 0; i < len(r.data); i++ {
		if i == r.current {
			builder.WriteString(fmt.Sprintf("[%d]", r.data[i]))
		} else {
			builder.WriteString(fmt.Sprintf("%d", r.data[i]))
		}

		if i < len(r.data)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}
