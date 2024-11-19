package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib"
)

func main() {
	ns := lib.InputToInts()
	for i, a := range ns {
		for j, b := range ns[i+1:] {
			for _, c := range ns[j+1:] {
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
