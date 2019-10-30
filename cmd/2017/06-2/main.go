package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	banks := InputToBanks(2017, 6)

	seen := make(map[string]int)

	var cycle int
	for cycle = 1; ; cycle++ {
		id := banks.ID()
		if seen[id] > 0 {
			break
		}

		seen[id] = cycle
		banks = banks.Distribute(banks.Largest())
	}

	length := cycle - seen[banks.ID()]
	fmt.Printf("cycle length: %d\n", length)
}

type Banks []int

func (b Banks) ID() string {
	var builder strings.Builder
	for i, bank := range b {
		builder.WriteString(fmt.Sprintf("%d", bank))
		if i < len(b)-1 {
			builder.WriteString(" ")
		}
	}

	return builder.String()
}

func (b Banks) Largest() int {
	max := b[0]

	var index int
	for i := 1; i < len(b); i++ {
		if b[i] > max {
			max = b[i]
			index = i
		}
	}

	return index
}

func (b Banks) Distribute(id int) Banks {
	updated := append(Banks(nil), b...)

	N := len(b)
	div := b[id] / N
	mod := b[id] % N

	updated[id] = 0
	for i := 1; i <= N; i++ {
		index := (id + i + N) % N

		updated[index] += div
		if mod > 0 {
			updated[index]++
			mod--
		}
	}

	return updated
}

func InputToBanks(year, day int) Banks {
	var banks Banks
	for _, token := range strings.Split(aoc.InputToString(year, day), "\t") {
		banks = append(banks, aoc.ParseInt(token))
	}

	return banks
}
