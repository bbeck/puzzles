package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for _, address := range InputToAddresses() {
		if SupportsTLS(address) {
			count++
		}
	}

	fmt.Println(count)
}

func SupportsTLS(a Address) bool {
	found := func(parts []string) bool {
		for _, s := range parts {
			for start := range len(s) - 3 {
				c0, c1, c2, c3 := s[start], s[start+1], s[start+2], s[start+3]
				if c0 != c2 && c0 == c3 && c1 == c2 {
					return true
				}
			}
		}
		return false
	}

	return found(a.Supernets) && !found(a.Hypernets)
}

type Address struct {
	Supernets []string
	Hypernets []string
}

func InputToAddresses() []Address {
	return in.LinesTo(func(in *in.Scanner[Address]) Address {
		var line = in.Line()
		line = strings.ReplaceAll(line, "[", " ")
		line = strings.ReplaceAll(line, "]", " ")
		fields := strings.Fields(line)

		var address Address
		for i, part := range fields {
			if i%2 == 0 {
				address.Supernets = append(address.Supernets, part)
			} else {
				address.Hypernets = append(address.Hypernets, part)
			}
		}

		return address
	})
}
