package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	fields, ticket, tickets := InputToFieldsAndTickets()

	// Create a set of allowed values for each field.
	allowed := make(map[string]lib.Set[int])
	for name, ranges := range fields {
		for _, r := range ranges {
			for v := r.Min; v <= r.Max; v++ {
				set := allowed[name]
				set.Add(v)
				allowed[name] = set
			}
		}
	}

	// Put every seen example value for a field into that field's set.
	examples := make([]lib.Set[int], len(ticket))
	for _, t := range tickets {
		if !IsValid(t, allowed) {
			continue
		}

		for i, v := range t {
			examples[i].Add(v)
		}
	}

	product := 1
	for index, name := range DeriveFieldMapping(allowed, examples) {
		if strings.HasPrefix(name, "departure") {
			product *= ticket[index]
		}
	}
	fmt.Println(product)
}

func IsValid(ticket []int, fields map[string]lib.Set[int]) bool {
	var vs lib.Set[int]
	vs.Add(ticket...)

	for _, s := range fields {
		vs = vs.Difference(s)
	}

	return len(vs) == 0
}

func DeriveFieldMapping(fields map[string]lib.Set[int], values []lib.Set[int]) []string {
	// Determine possible fields for each set of values.
	choices := make([]lib.Set[string], len(values))
	for i := 0; i < len(values); i++ {
		for name, possible := range fields {
			if len(values[i].Difference(possible)) == 0 {
				choices[i].Add(name)
			}
		}
	}

	var assigned lib.Set[string]
	mapping := make([]string, len(values))

	for len(assigned) < len(fields) {
		for i := 0; i < len(choices); i++ {
			choices[i] = choices[i].Difference(assigned)

			if len(choices[i]) == 1 {
				mapping[i] = choices[i].Entries()[0]
				assigned.Add(mapping[i])
			}
		}
	}

	return mapping
}

type Range struct {
	Min, Max int
}

func InputToFieldsAndTickets() (map[string][]Range, []int, [][]int) {
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

	index += 2
	var ticket []int
	for _, s := range strings.Split(lines[index], ",") {
		ticket = append(ticket, lib.ParseInt(s))
	}

	index += 3
	var tickets [][]int
	for ; index < len(lines); index++ {
		var values []int
		for _, s := range strings.Split(lines[index], ",") {
			values = append(values, lib.ParseInt(s))
		}

		tickets = append(tickets, values)
	}

	return fields, ticket, tickets
}
