package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	grid := make([][]bool, 0)

	input := aoc.InputToString(2017, 14)
	for i := 0; i < 128; i++ {
		row := make([]bool, 0)
		for _, c := range KnotHash(fmt.Sprintf("%s-%d", input, i)) {
			switch c {
			case '0':
				row = append(row, []bool{false, false, false, false}...)
			case '1':
				row = append(row, []bool{false, false, false, true}...)
			case '2':
				row = append(row, []bool{false, false, true, false}...)
			case '3':
				row = append(row, []bool{false, false, true, true}...)
			case '4':
				row = append(row, []bool{false, true, false, false}...)
			case '5':
				row = append(row, []bool{false, true, false, true}...)
			case '6':
				row = append(row, []bool{false, true, true, false}...)
			case '7':
				row = append(row, []bool{false, true, true, true}...)
			case '8':
				row = append(row, []bool{true, false, false, false}...)
			case '9':
				row = append(row, []bool{true, false, false, true}...)
			case 'a':
				row = append(row, []bool{true, false, true, false}...)
			case 'b':
				row = append(row, []bool{true, false, true, true}...)
			case 'c':
				row = append(row, []bool{true, true, false, false}...)
			case 'd':
				row = append(row, []bool{true, true, false, true}...)
			case 'e':
				row = append(row, []bool{true, true, true, false}...)
			case 'f':
				row = append(row, []bool{true, true, true, true}...)
			}
		}

		grid = append(grid, row)
	}

	var count int
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] {
				count++
			}
		}
	}

	fmt.Printf("used: %d\n", count)
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
