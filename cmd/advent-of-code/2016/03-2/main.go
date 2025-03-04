package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for _, triangle := range InputToTriangles() {
		if triangle.IsValid() {
			count++
		}
	}

	fmt.Println(count)
}

type Triangle struct {
	Side1, Side2, Side3 int
}

func (t Triangle) IsValid() bool {
	return t.Side1+t.Side2 > t.Side3 &&
		t.Side1+t.Side3 > t.Side2 &&
		t.Side2+t.Side3 > t.Side1
}

func InputToTriangles() []Triangle {
	var triangles []Triangle
	var t1, t2, t3 Triangle
	for in.HasNext() {
		t1.Side1, t2.Side1, t3.Side1 = in.Int(), in.Int(), in.Int()
		t1.Side2, t2.Side2, t3.Side2 = in.Int(), in.Int(), in.Int()
		t1.Side3, t2.Side3, t3.Side3 = in.Int(), in.Int(), in.Int()
		triangles = append(triangles, t1, t2, t3)
	}

	return triangles
}
