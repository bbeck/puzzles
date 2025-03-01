package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	row, col := in.Int(), in.Int()

	// We first have to figure out which element of the sequence our coordinate
	// corresponds to.
	//
	//    | 1   2   3   4   5   6   7   8   9
	// ---+---+---+---+---+---+---+---+---+---+
	// 1 |  1   3   6  10  15  21  28  36  45
	// 2 |  2   5   9  14  20  27  35  44
	// 3 |  4   8  13  19  26  34  43
	// 4 |  7  12  18  25  33  42
	// 5 | 11  17  34  32  41
	// 6 | 16  23  31  40
	// 7 | 22  30  39
	// 8 | 29  38
	// 9 | 37
	//
	// To do this start by determining the value in column 1 of the desired row.
	//
	// f(row, 1) = 1 + 1+...+row-1 = 1 + (row-1)*row/2
	//
	// Next, we determine the value in the column our coordinate corresponds to.
	// To do this start at the value in the first column and add to it the value
	// of our row then row+1, and so on.
	//
	// f(row, col) = f(row, 1) + (row+1) + (row+2) + ... + (row+col-1)
	//             = [1 + 1+...+(row-1)] + [(row+1)+...(row+col-1)]
	//             = [1 + 1+...+(row-1)] + row + [(row+1)+...(row+col-1)] - row
	//             = (1-row) + (row+col-1)*(row+col)/2
	index := (1 - row) + (row+col-1)*(row+col)/2

	code := 20151125
	for i := 1; i < index; i++ {
		code = (code * 252533) % 33554393
	}
	fmt.Println(code)
}
