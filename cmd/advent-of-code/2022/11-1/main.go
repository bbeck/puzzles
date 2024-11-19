package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/lib"
	"sort"
	"strings"
)

func main() {
	monkeys := InputToMonkeys()

	for n := 0; n < 20; n++ {
		for _, monkey := range monkeys {
			for !monkey.Items.Empty() {
				monkey.Inspections++

				worry := monkey.Op(monkey.Items.PopFront()) / 3
				if worry%monkey.Mod == 0 {
					monkeys[monkey.TrueID].Items.PushBack(worry)
				} else {
					monkeys[monkey.FalseID].Items.PushBack(worry)
				}
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})
	fmt.Println(monkeys[0].Inspections * monkeys[1].Inspections)
}

type Monkey struct {
	Items       *lib.Deque[int]
	Op          func(int) int
	Mod         int
	TrueID      int
	FalseID     int
	Inspections int
}

func InputToMonkeys() []*Monkey {
	lines := lib.InputToLines()

	var monkeys []*Monkey
	for i := 0; i < len(lines); i += 7 {
		var id int
		fmt.Sscanf(lines[i+0], "Monkey %d:", &id)

		var monkey Monkey
		monkey.Items = new(lib.Deque[int])
		for _, s := range strings.Split(lines[i+1][18:], ", ") {
			monkey.Items.PushBack(lib.ParseInt(s))
		}

		monkey.Op = ParseOp(lines[i+2][19:])
		monkey.Mod = lib.ParseInt(lines[i+3][21:])
		monkey.TrueID = lib.ParseInt(lines[i+4][29:])
		monkey.FalseID = lib.ParseInt(lines[i+5][30:])

		monkeys = append(monkeys, &monkey)
	}

	return monkeys
}

func ParseOp(s string) func(int) int {
	fields := strings.Fields(s)
	_, op, rhs := fields[0], fields[1], fields[2]

	switch {
	case op == "+" && rhs != "old":
		n := lib.ParseInt(rhs)
		return func(old int) int { return old + n }

	case op == "+" && rhs == "old":
		return func(old int) int { return old + old }

	case op == "*" && rhs != "old":
		n := lib.ParseInt(rhs)
		return func(old int) int { return old * n }

	case op == "*" && rhs == "old":
		return func(old int) int { return old * old }
	}

	return nil
}
