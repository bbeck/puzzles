package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	fields, tickets := InputToFieldsAndTickets()

	var valid lib.Set[int]
	for _, rs := range fields {
		for _, r := range rs {
			for i := r.Min; i <= r.Max; i++ {
				valid.Add(i)
			}
		}
	}

	var sum int
	for _, ticket := range tickets {
		for _, value := range ticket {
			if !valid.Contains(value) {
				sum += value
			}
		}
	}
	fmt.Println(sum)
}

type Range struct {
	Min, Max int
}

func InputToFieldsAndTickets() (map[string][]Range, [][]int) {
	lines := lib.InputToLines()
	fields := make(map[string][]Range)

	var index int
	for ; lines[index] != ""; index++ {
		key, rest, _ := strings.Cut(lines[index], ": ")
		r1, r2, _ := strings.Cut(rest, " or ")
		min1, max1, _ := strings.Cut(r1, "-")
		min2, max2, _ := strings.Cut(r2, "-")

		fields[key] = append(fields[key], Range{Min: lib.ParseInt(min1), Max: lib.ParseInt(max1)})
		fields[key] = append(fields[key], Range{Min: lib.ParseInt(min2), Max: lib.ParseInt(max2)})
	}

	// Skip over my ticket and blank lines
	index += 5

	var tickets [][]int
	for ; index < len(lines); index++ {
		var values []int
		for _, s := range strings.Split(lines[index], ",") {
			values = append(values, lib.ParseInt(s))
		}

		tickets = append(tickets, values)
	}

	return fields, tickets
}
