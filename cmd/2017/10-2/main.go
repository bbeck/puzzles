package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	bs := aoc.InputToBytes(2017, 10)
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

	fmt.Printf("hash: %s\n", ToHex(DenseHash(ring.data)))
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
		s = s + fmt.Sprintf("%2x", i)
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
