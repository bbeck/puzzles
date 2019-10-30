package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	lengths := InputToInts(2017, 10)

	ring := &Ring{data: make([]int, 0)}
	for i := 0; i <= 255; i++ {
		ring.Append(i)
	}

	for _, length := range lengths {
		ring.Twist(length)
	}

	fmt.Printf("checksum: %d\n", ring.data[0]*ring.data[1])
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

func InputToInts(year, day int) []int {
	var ints []int
	for _, part := range strings.Split(aoc.InputToString(year, day), ",") {
		ints = append(ints, aoc.ParseInt(part))
	}

	return ints
}
