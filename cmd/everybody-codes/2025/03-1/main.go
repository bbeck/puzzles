package main

import (
	"fmt"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	numbers := SetFrom(in.Ints()...).Entries()
	fmt.Println(Sum(numbers...))
}
