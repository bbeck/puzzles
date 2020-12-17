package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ranges, myTicket, nearbyTickets := Input(2020, 16)

	// Use the nearby tickets along with the valid ranges for fields to determine
	// which indices of a ticket correspond to which field.  We'll do this by
	// eliminating possibilities when a ticket index falls out of bounds of a
	// field's valid range.

	// Initialize every index to all possible field names.
	possibilities := make(map[int]*aoc.Set)
	for i := 0; i < len(myTicket); i++ {
		s := aoc.NewSet()
		for name := range ranges {
			s.Add(name)
		}

		possibilities[i] = &s
	}

	// Go through each valid nearby ticket and use it to limit the possible fields
	// for each index based on the valid ranges the field.
	for _, ticket := range nearbyTickets {
		if !IsValid(ticket, ranges) {
			continue
		}

		for index, value := range ticket {
			for name, r := range ranges {
				if !r.Contains(value) {
					possibilities[index].Remove(name)
				}
			}
		}
	}

	// At this point we can go through and make assignments when a possibility set
	// ends up with just a single field in it.  Once a field is assigned it needs
	// to be removed from any other possibility.  This process will require
	// multiple passes through the sets.
	assigned := aoc.NewSet()
	assignments := map[string]int{}
	for len(assignments) < len(possibilities) {
		for index, ps := range possibilities {
			entries := ps.Difference(assigned).Entries()
			if len(entries) == 1 {
				entry := entries[0].(string)
				assigned.Add(entry)
				assignments[entry] = index
			}
		}
	}

	// Now find the departure fields and the values from my ticket and multiply
	// them together.
	var product = 1
	for name, index := range assignments {
		if strings.HasPrefix(name, "departure") {
			product *= myTicket[index]
		}
	}
	fmt.Println(product)
}

func IsValid(ticket Ticket, ranges Ranges) bool {
	var isValid = true

	for _, n := range ticket {
		var ok bool
		for _, r := range ranges {
			if r.Contains(n) {
				ok = true
				break
			}
		}

		if !ok {
			isValid = false
		}
	}

	return isValid
}

type Ranges map[string]aoc.Set
type Ticket []int

func ParseRange(s string) (string, aoc.Set) {
	fields := strings.Split(s, ": ")
	name := fields[0]
	rhs := strings.ReplaceAll(fields[1], "-", " ")

	var min1, max1, min2, max2 int
	_, err := fmt.Sscanf(rhs, "%d %d or %d %d", &min1, &max1, &min2, &max2)
	if err != nil {
		log.Fatalf("unable to parse range: %s", s)
	}

	set := aoc.NewSet()
	for i := min1; i <= max1; i++ {
		set.Add(i)
	}
	for i := min2; i <= max2; i++ {
		set.Add(i)
	}

	return name, set
}

func ParseTicket(s string) []int {
	var ticket []int
	for _, n := range strings.Split(s, ",") {
		ticket = append(ticket, aoc.ParseInt(n))
	}

	return ticket
}

func Input(year, day int) (Ranges, Ticket, []Ticket) {
	lines := aoc.InputToLines(year, day)

	var current int

	var ranges = make(Ranges)
	for {
		if len(lines[current]) == 0 {
			break
		}

		name, set := ParseRange(lines[current])
		ranges[name] = set

		current++
	}
	current++ // blank line
	current++ // 'your ticket:' line

	// Parse my ticket (skipping the "your ticket:" line)
	var myTicket = ParseTicket(lines[current])

	current++ // my ticket
	current++ // blank line
	current++ // 'nearby tickets:' line

	var nearbyTickets []Ticket
	for {
		if current >= len(lines) {
			break
		}

		nearbyTickets = append(nearbyTickets, ParseTicket(lines[current]))
		current++
	}

	return ranges, myTicket, nearbyTickets
}
