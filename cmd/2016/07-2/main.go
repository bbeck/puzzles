package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, address := range InputToAddresses(2016, 7) {
		if address.SupportsSSL() {
			count++
		}
	}

	fmt.Printf("count: %d\n", count)
}

type Address struct {
	parts    []string
	hypernet []string
}

func InputToAddresses(year, day int) []Address {
	var addresses []Address
	for _, line := range aoc.InputToLines(year, day) {
		var address Address
		for _, p := range strings.Split(line, "]") {
			for i, part := range strings.Split(p, "[") {
				if i%2 == 0 {
					address.parts = append(address.parts, part)
				} else {
					address.hypernet = append(address.hypernet, part)
				}
			}
		}

		addresses = append(addresses, address)
	}

	return addresses
}

type Pair struct {
	c1, c2 byte
}

func (a Address) SupportsSSL() bool {
	pairs := func(parts []string) []Pair {
		var pairs []Pair
		for _, part := range parts {
			for i := 0; i < len(part)-2; i++ {
				if part[i] == part[i+2] && part[i] != part[i+1] {
					pairs = append(pairs, Pair{part[i], part[i+1]})
				}
			}
		}

		return pairs
	}

	for _, pair := range pairs(a.parts) {
		bab := fmt.Sprintf("%c%c%c", pair.c2, pair.c1, pair.c2)
		for _, s := range a.hypernet {
			if strings.Contains(s, bab) {
				return true
			}
		}
	}

	return false
}
