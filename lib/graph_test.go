package lib

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestMinimumSpanningTree(t *testing.T) {
	type test struct {
		name     string
		vertices []string
		children func(string) []string
		cost     func(string, string) int
		total    int
		edges    []WeightedEdge[string]
	}

	tests := []test{
		{
			name:     "wikipdia example",
			vertices: []string{"A", "B", "C", "D", "E", "F", "G"},
			children: func(s string) []string {
				return map[string][]string{
					"A": {"B", "D"},
					"B": {"A", "C", "D", "E"},
					"C": {"B", "E"},
					"D": {"A", "B", "E", "F"},
					"E": {"B", "C", "D", "F", "G"},
					"F": {"D", "E", "G"},
					"G": {"E", "F"},
				}[s]
			},
			cost: func(a, b string) int {
				return map[string]map[string]int{
					"A": {"B": 7, "D": 5},
					"B": {"A": 7, "C": 8, "D": 9, "E": 7},
					"C": {"B": 8, "E": 5},
					"D": {"A": 5, "B": 9, "E": 15, "F": 6},
					"E": {"B": 7, "C": 5, "D": 15, "F": 8, "G": 9},
					"F": {"D": 6, "E": 8, "G": 11},
					"G": {"E": 9, "F": 11},
				}[a][b]
			},
			total: 39,
			edges: []WeightedEdge[string]{
				{From: "A", To: "B", Weight: 7},
				{From: "A", To: "D", Weight: 5},
				{From: "B", To: "E", Weight: 7},
				{From: "C", To: "E", Weight: 5},
				{From: "D", To: "F", Weight: 6},
				{From: "E", To: "G", Weight: 9},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cost, edges := MinimumSpanningTree(test.vertices, test.children, test.cost)
			assert.Equal(t, test.total, cost)

			// Edges may be in either direction.
			for i := range edges {
				if edges[i].From > edges[i].To {
					edges[i].From, edges[i].To = edges[i].To, edges[i].From
				}
			}
			assert.Equal(t, SetFrom(test.edges...), SetFrom(edges...))
		})
	}
}
