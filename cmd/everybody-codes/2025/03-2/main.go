package main

import (
	"fmt"
	"sort"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	numbers := SetFrom(in.Ints()...).Entries()
	sort.Ints(numbers)

	fmt.Println(Sum(numbers[:20]...))
}
