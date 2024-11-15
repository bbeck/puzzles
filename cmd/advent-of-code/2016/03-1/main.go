package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
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
	return puz.InputLinesTo(func(line string) Triangle {
		parts := strings.Fields(line)
		return Triangle{
			Side1: puz.ParseInt(parts[0]),
			Side2: puz.ParseInt(parts[1]),
			Side3: puz.ParseInt(parts[2]),
		}
	})
}
