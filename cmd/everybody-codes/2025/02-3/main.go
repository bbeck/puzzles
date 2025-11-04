package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	Tx, Ty := in.Int(), in.Int()
	Bx, By := Add(Tx, Ty, 1000, 1000)

	var count int
	for x := Tx; x <= Bx; x++ {
		for y := Ty; y <= By; y++ {
			if IsEngraved(x, y) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func IsEngraved(Nx, Ny int) bool {
	var Rx, Ry int
	for range 100 {
		Rx, Ry = Mul(Rx, Ry, Rx, Ry)
		Rx, Ry = Div(Rx, Ry, 100000, 100000)
		Rx, Ry = Add(Rx, Ry, Nx, Ny)

		if Rx < -1000000 || Rx > 1000000 || Ry < -1000000 || Ry > 1000000 {
			return false
		}
	}
	return true
}

func Add(X1, Y1, X2, Y2 int) (int, int) { return X1 + X2, Y1 + Y2 }
func Mul(X1, Y1, X2, Y2 int) (int, int) { return X1*X2 - Y1*Y2, X1*Y2 + Y1*X2 }
func Div(X1, Y1, X2, Y2 int) (int, int) { return X1 / X2, Y1 / Y2 }
