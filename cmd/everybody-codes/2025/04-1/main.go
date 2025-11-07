package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	gears := in.Ints()
	fmt.Println(2025 * gears[0] / gears[len(gears)-1])
}
