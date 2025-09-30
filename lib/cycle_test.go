package lib

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestFindCycle(t *testing.T) {
	type test struct {
		name   string
		next   func(in int) int
		prefix []int
		cycle  []int
	}

	tests := []test{
		{
			name: "no prefix",
			next: func(in int) int {
				if in < 9 {
					return in + 1
				}
				return 0
			},
			prefix: []int{},
			cycle:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "length 1 prefix",
			next: func(in int) int {
				if in < 9 {
					return in + 1
				}
				return 1
			},
			prefix: []int{0},
			cycle:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "length 3 prefix",
			next: func(in int) int {
				if in < 9 {
					return in + 1
				}
				return 3
			},
			prefix: []int{0, 1, 2},
			cycle:  []int{3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "length 1 cycle",
			next: func(in int) int {
				return 0
			},
			prefix: []int{},
			cycle:  []int{0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prefix, cycle := FindCycle(0, test.next)
			assert.Equal(t, test.prefix, prefix)
			assert.Equal(t, test.cycle, cycle)
		})
	}
}
