package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var count int
	for _, address := range InputToAddresses() {
		if SupportsSSL(address) {
			count++
		}
	}

	fmt.Println(count)
}

func SupportsSSL(a Address) bool {
	var abas []string
	for _, supernet := range a.Supernets {
		for start := range len(supernet) - 2 {
			c0, c1, c2 := supernet[start], supernet[start+1], supernet[start+2]
			if c0 != c1 && c0 == c2 {
				abas = append(abas, supernet[start:start+3])
			}
		}
	}

	for _, aba := range abas {
		bab := fmt.Sprintf("%c%c%c", aba[1], aba[0], aba[1])
		for _, hypernet := range a.Hypernets {
			if strings.Contains(hypernet, bab) {
				return true
			}
		}
	}

	return false
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
