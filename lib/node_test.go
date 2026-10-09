package lib

import (
	"math"
	"slices"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestDijkstra(t *testing.T) {
	type test struct {
		name          string
		start         int
		children      func(int) []int
		cost          func(int, int) int
		expectedCosts map[int]int
		expectedPrevs map[int][]int
	}

	tests := []test{
		{
			name:  "GeeksForGeeks example",
			start: 0,
			children: func(n int) []int {
				return [][]int{
					{1, 2},
					{0, 4},
					{0, 3},
					{2, 4},
					{1, 3},
				}[n]
			},
			cost: func(f, t int) int {
				return [][]int{
					{0, 4, 8, math.MaxInt, math.MaxInt},
					{4, 0, math.MaxInt, math.MaxInt, 6},
					{8, math.MaxInt, 0, 2, math.MaxInt},
					{math.MaxInt, math.MaxInt, 2, 0, 10},
					{math.MaxInt, 6, math.MaxInt, 10, 0},
				}[f][t]
			},
			expectedCosts: map[int]int{0: 0, 1: 4, 2: 8, 3: 10, 4: 10},
			expectedPrevs: map[int][]int{1: {0}, 2: {0}, 3: {2}, 4: {1}},
		},
		{
			name:  "Plainenglish.io example",
			start: 0,
			children: func(n int) []int {
				return [][]int{
					{1, 2},
					{0, 2, 3},
					{0, 1, 3, 4},
					{1, 2, 4, 5},
					{1, 2, 3, 5},
					{3, 4},
				}[n]
			},
			cost: func(f, t int) int {
				return [][]int{
					{0, 9, 4, math.MaxInt, math.MaxInt, math.MaxInt},
					{9, 0, 2, 7, 3, math.MaxInt},
					{4, 2, 0, 1, 6, math.MaxInt},
					{math.MaxInt, 7, 1, 0, 4, 8},
					{math.MaxInt, 3, 6, 4, 0, 2},
					{math.MaxInt, math.MaxInt, math.MaxInt, 8, 2, 0},
				}[f][t]
			},
			expectedCosts: map[int]int{0: 0, 1: 6, 2: 4, 3: 5, 4: 9, 5: 11},
			expectedPrevs: map[int][]int{1: {2}, 2: {0}, 3: {2}, 4: {3}, 5: {4}},
		},
		{
			name:  "Diamond (requires multiple previous nodes)",
			start: 0,
			children: func(n int) []int {
				return [][]int{
					{1, 2},
					{0, 3},
					{0, 3},
					{1, 2},
				}[n]
			},
			cost: func(f, t int) int {
				return [][]int{
					{math.MaxInt, 1, 1, math.MaxInt},
					{1, 0, math.MaxInt, 1},
					{1, math.MaxInt, math.MaxInt, 1},
					{math.MaxInt, 1, 1, math.MaxInt},
				}[f][t]
			},
			expectedCosts: map[int]int{0: 0, 1: 1, 2: 1, 3: 2},
			expectedPrevs: map[int][]int{1: {0}, 2: {0}, 3: {1, 2}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dist, prevs := Dijkstra(test.start, test.children, test.cost)
			assert.Equal(t, test.expectedCosts, dist)

			// The predecessors of a node come from a set, so their order is not
			// deterministic.  Sort both sides so only the contents are compared.
			for _, ps := range prevs {
				slices.Sort(ps)
			}
			for _, ps := range test.expectedPrevs {
				slices.Sort(ps)
			}
			assert.Equal(t, test.expectedPrevs, prevs)
		})
	}
}
