package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	numPriests, numAcolytes, blocks := in.Int(), 10, 202400000
	pyramid := []int{1}

	for thickness, width := 1, 3; ; width += 2 {
		thickness = (thickness*numPriests)%numAcolytes + numAcolytes

		pyramid = append(append([]int{0}, pyramid...), 0)
		for i := 0; i < len(pyramid); i++ {
			pyramid[i] += thickness
		}

		used := Sum(pyramid...)
		for i := 1; i < len(pyramid)-1; i++ {
			used -= (width * numPriests * pyramid[i]) % numAcolytes
		}

		if used > blocks {
			blocks -= used
			break
		}
	}

	fmt.Println(Abs(blocks))
}
