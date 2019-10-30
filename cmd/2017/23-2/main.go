package main

import (
	"fmt"
)

func main() {
	isComposite := func(n int) bool {
		if n == 2 {
			return false
		}

		if n%2 == 0 {
			return true
		}

		for i := 3; i <= n/2; i += 2 {
			if n%i == 0 {
				return true
			}
		}

		return false
	}

	var count int
	for b := 106700; b <= 123700; b += 17 {
		if isComposite(b) {
			count++
		}
	}

	fmt.Printf("h: %d\n", count)
}

func main2() {
	var h int

	for b := 106700; b <= 123700; b += 17 {
		var isComposite bool

	outer:
		for d := 2; d < b; d++ {
			for e := 2; e < b; e++ {
				if d*e == b {
					isComposite = true
					break outer
				}
			}
		}

		if isComposite {
			h++
		}
	}

	fmt.Printf("h: %d\n", h)
}
