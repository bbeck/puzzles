package main

import (
	"fmt"
	"maps"
	"slices"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var ps []Point3D
	for in.HasNext() {
		ps = append(ps, Point3D{X: in.Int(), Y: in.Int(), Z: in.Int()})
	}

	var distances = make(map[int][][2]Point3D)
	for i := 0; i < len(ps); i++ {
		for j := i + 1; j < len(ps); j++ {
			d := Distance2(ps[i], ps[j])
			distances[d] = append(distances[d], [2]Point3D{ps[i], ps[j]})
		}
	}

	var pairs [][2]Point3D
	for _, dist := range slices.Sorted(maps.Keys(distances)) {
		pairs = append(pairs, distances[dist]...)
	}

	var ds DisjointSet[Point3D]
	for _, p := range pairs {
		ds.UnionWithAdd(p[0], p[1])
		if IsFullyConnected(ps, ds) {
			fmt.Println(p[0].X * p[1].X)
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
