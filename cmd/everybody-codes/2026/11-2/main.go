package main

import (
	"fmt"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	portal, edges := InputToPortalAndEdges()
	n := len(portal)

	dist1 := FloydWarshall(portal)
	dist2 := UpdateWithEdges(dist1, edges)

	var best int
	for start := range n {
		for end := start + 1; end < n; end++ {
			dist3 := UpdateWithEdges(dist2, [][2]int{{start, end}})

			var sum int
			for i := range dist1 {
				for j := range dist1[i] {
					sum += dist1[i][j] - dist3[i][j]
				}
			}

			best = Max(best, sum)
		}
	}
	fmt.Println(best)
}

func FloydWarshall(portal Portal) [][]int {
	n := len(portal)

	dist := Make2D[int](n, n)
	for i := range n {
		for j := range n {
			if i == j {
				continue
			}
			if portal[i][j] {
				dist[i][j] = 1
			} else {
				dist[i][j] = n
			}
		}
	}
	for k := range n {
		for i := range n {
			for j := range n {
				dist[i][j] = Min(dist[i][j], dist[i][k]+dist[k][j])
			}
		}
	}
	return dist
}

func UpdateWithEdges(dist [][]int, edges [][2]int) [][]int {
	dist = slices.Clone(dist)
	for i := range dist {
		dist[i] = slices.Clone(dist[i])
	}

	// Update the distance matrix for the additional edges.  This operates on a
	// simple principle. We can consider a new edge in the graph just by looking
	// at each pair of nodes and seeing if there is a shorter path through the
	// new edge.
	for _, edge := range edges {
		for i := range dist {
			for j := range dist[i] {
				dist[i][j] = Min(
					dist[i][j],
					dist[i][edge[0]]+1+dist[edge[1]][j],
					dist[i][edge[1]]+1+dist[edge[0]][j],
				)
			}
		}
	}

	return dist
}

type Portal [][]bool

func InputToPortalAndEdges() (Portal, [][2]int) {
	n := in.Int()

	var portal = Make2D[bool](n, n)
	for node := range n {
		portal[node][(node+1)%n] = true
		portal[node][(node+2)%n] = true
		portal[node][(node+3)%n] = true
		portal[(node+1)%n][node] = true
		portal[(node+2)%n][node] = true
		portal[(node+3)%n][node] = true
	}

	var edges [][2]int
	for in.HasNext() {
		from, to := in.Int()-1, in.Int()-1
		edges = append(edges, [2]int{from, to})
	}

	return portal, edges
}
