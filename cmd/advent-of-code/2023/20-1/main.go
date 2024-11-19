package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	modules := InputToModules()
	states := make(map[string]bool)
	pulses := lib.Deque[Pulse]{}
	counts := make(map[bool]int)

	for n := 0; n < 1000; n++ {
		pulses.PushBack(Pulse{Source: "button", Target: "broadcaster", Value: false})

		for !pulses.Empty() {
			pulse := pulses.PopFront()
			counts[pulse.Value]++

			module, state := modules[pulse.Target], states[pulse.Target]
			switch module.Kind {
			case "%":
				if pulse.Value {
					continue
				}
				state = !state

			case "&":
				state = true
				for _, parent := range module.Parents {
					state = state && states[parent]
				}
				state = !state

			case "":
				state = pulse.Value
			}

			states[pulse.Target] = state
			for _, target := range module.Targets {
				pulses.PushBack(Pulse{Source: module.ID, Target: target, Value: state})
			}
		}
	}

	fmt.Println(lib.Product(lib.GetMapValues(counts)...))
}

type Pulse struct {
	Source, Target string
	Value          bool
}

type Module struct {
	Kind    string
	ID      string
	Targets []string
	Parents []string
}

func InputToModules() map[string]Module {
	ids := lib.Set[string]{}
	kinds := make(map[string]string)
	targets := make(map[string][]string)
	parents := make(map[string][]string)

	for _, line := range lib.InputToLines() {
		kind, id, children := ParseModule(line)

		ids.Add(id)
		kinds[id] = kind
		targets[id] = children
		for _, child := range children {
			ids.Add(child)
			parents[child] = append(parents[child], id)
		}
	}

	modules := make(map[string]Module)
	for id := range ids {
		modules[id] = Module{
			Kind:    kinds[id],
			ID:      id,
			Targets: targets[id],
			Parents: parents[id],
		}
	}
	return modules
}

func ParseModule(s string) (string, string, []string) {
	lhs, rhs, _ := strings.Cut(s, " -> ")
	targets := strings.Split(rhs, ", ")

	if lhs[0] == '%' || lhs[0] == '&' {
		return string(lhs[0]), lhs[1:], targets
	}
	return "", lhs, targets
}
