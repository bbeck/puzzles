package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
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
	nums := lib.InputLinesTo(func(line string) []int {
		parts := strings.Fields(line)
		return []int{
			lib.ParseInt(parts[0]),
			lib.ParseInt(parts[1]),
			lib.ParseInt(parts[2]),
		}
	})

	var triangles []Triangle
	for row := 0; row < len(nums); row += 3 {
		for col := 0; col < 3; col++ {
			triangles = append(triangles, Triangle{
				Side1: nums[row+0][col],
				Side2: nums[row+1][col],
				Side3: nums[row+2][col],
			})
		}
	}
	return triangles
}
