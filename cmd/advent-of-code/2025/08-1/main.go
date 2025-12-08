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
	for i := range 1000 {
		ds.UnionWithAdd(pairs[i][0], pairs[i][1])
	}

	var sizes = make(map[Point3D]int)
	for _, p := range ps {
		id, _ := ds.Find(p)
		sizes[id] = ds.Size(id)
	}

	var largest = slices.Sorted(maps.Values(sizes))
	fmt.Println(Product(largest[len(largest)-3:]...))
}

func Distance2(p, q Point3D) int {
	dx, dy, dz := p.X-q.X, p.Y-q.Y, p.Z-q.Z
	return dx*dx + dy*dy + dz*dz
}
