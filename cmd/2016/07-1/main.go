package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	var count int
	for _, address := range InputToAddresses(2016, 7) {
		if address.SupportsTLS() {
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

func (a Address) SupportsTLS() bool {
	found := func(parts []string) bool {
		for _, part := range parts {
			for i := 0; i < len(part)-3; i++ {
				if part[i] == part[i+3] && part[i+1] == part[i+2] && part[i] != part[i+1] {
					return true
				}
			}
		}

		return false
	}

	return found(a.parts) && !found(a.hypernet)
}
