package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
)

func main() {
	ns := lib.InputToInts()

	var count int
	for i := 1; i < len(ns); i++ {
		if ns[i] > ns[i-1] {
			count++
		}
	}
	fmt.Println(count)
}
