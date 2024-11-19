package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	ns := lib.InputToInts()

	for i := 25; i < len(ns); i++ {
		if !SumExists(ns[i-25:i], ns[i]) {
			fmt.Println(ns[i])
			break
		}
	}
}

func SumExists(ns []int, target int) bool {
	var seen lib.Set[int]
	for _, n := range ns {
		if seen.Contains(target - n) {
			return true
		}
		seen.Add(n)
	}
	return false
}
