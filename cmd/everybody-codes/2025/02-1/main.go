package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	Ax, Ay := in.Int(), in.Int()
	Rx, Ry := 0, 0

	for range 3 {
		Rx, Ry = Mul(Rx, Ry, Rx, Ry)
		Rx, Ry = Div(Rx, Ry, 10, 10)
		Rx, Ry = Add(Rx, Ry, Ax, Ay)
	}
	fmt.Printf("[%d,%d]\n", Rx, Ry)
}

func Add(X1, Y1, X2, Y2 int) (int, int) { return X1 + X2, Y1 + Y2 }
func Mul(X1, Y1, X2, Y2 int) (int, int) { return X1*X2 - Y1*Y2, X1*Y2 + Y1*X2 }
func Div(X1, Y1, X2, Y2 int) (int, int) { return X1 / X2, Y1 / Y2 }
