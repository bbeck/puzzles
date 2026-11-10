package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	portal, edges := InputToPortalAndEdges()
	dist1 := FloydWarshall(portal)

	for _, edge := range edges {
		portal[edge[0]][edge[1]] = true
		portal[edge[1]][edge[0]] = true
	}
	dist2 := FloydWarshall(portal)

	var sum int
	for i := range dist1 {
		for j := range dist1[i] {
			sum += dist1[i][j] - dist2[i][j]
		}
	}
	fmt.Println(sum)
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

type Portal [][]bool

func InputToPortalAndEdges() (Portal, [][2]int) {
	n := in.Int()

	var portal = Make2D[bool](n, n)
	for node := range n {
		portal[node][(node+1)%n] = true
		portal[(node+1)%n][node] = true
	}

	var edges [][2]int
	for in.HasNext() {
		from, to := in.Int()-1, in.Int()-1
		edges = append(edges, [2]int{from, to})
	}

	return portal, edges
}
