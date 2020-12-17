package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	ranges, myTicket, nearbyTickets := Input(2020, 16)

	var sum int
	sum += ErrorRate(myTicket, ranges)
	for _, ticket := range nearbyTickets {
		sum += ErrorRate(ticket, ranges)
	}

	fmt.Println(sum)
}

func ErrorRate(ticket Ticket, ranges Ranges) int {
	var sum int
	for _, n := range ticket {
		var ok bool
		for _, r := range ranges {
			if r.Contains(n) {
				ok = true
			}
		}

		if !ok {
			sum += n
		}
	}

	return sum
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
