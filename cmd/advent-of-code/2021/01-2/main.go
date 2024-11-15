package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	ns := puz.InputToInts(2021, 1)

	var count int
	for i := 1; i < len(ns)-2; i++ {
		if ns[i+2] > ns[i-1] { // ns[i] and ns[i+1] are common to both windows
			count++
		}
	}
	fmt.Println(count)
}
