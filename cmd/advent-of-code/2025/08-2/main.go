package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ps []Point3D
	for in.HasNext() {
		ps = append(ps, Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()})
	}

	var distances PriorityQueue[[2]Point3D]
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			distances.Push([2]Point3D{ps[i], ps[j]}, Distance2(ps[i], ps[j]))
		}
	}

	var ds DisjointSet[Point3D]
	for {
		pair := distances.Pop()
		ds.UnionWithAdd(pair[0], pair[1])
		if IsFullyConnected(ps, ds) {
			fmt.Println(pair[0].X * pair[1].X)
			break
		}
	}
}

func IsFullyConnected(ps []Point3D, ds DisjointSet[Point3D]) bool {
	root, _ := ds.Find(ps[0])
	for _, p := range ps[1:] {
		id, _ := ds.Find(p)
		if root != id {
			return false
		}
	}
	return true
}

func Distance2(p, q Point3D) int {
	dx, dy, dz := p.X-q.X, p.Y-q.Y, p.Z-q.Z
	return dx*dx + dy*dy + dz*dz
}
