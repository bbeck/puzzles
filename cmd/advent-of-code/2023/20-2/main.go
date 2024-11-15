package main

import (
	"fmt"
	"github.com/bbeck/advent-of-code/puz"
	"strings"
)

func main() {
	modules := InputToModules()

	// This solution only works because the input is crafted such that the rx node
	// is fed by 2 layers of conjunctions.  So determining when it transitions is
	// as simple as watching to see what frequency the grandparents turn on/off
	// at, and then extrapolating it to figure out the button push that causes all
	// grandparents to synchronize.
	grandparents := GetGrandparents("rx", modules)
	pushes := GetNumPushesToTurnOn(grandparents, modules)

	fmt.Println(puz.LCM(pushes...))
}

func GetGrandparents(id string, modules map[string]Module) []string {
	var gp puz.Set[string]
	for _, parent := range modules[id].Parents {
		gp.Add(modules[parent].Parents...)
	}

	return gp.Entries()
}

func GetNumPushesToTurnOn(ids []string, modules map[string]Module) []int {
	needs := puz.SetFrom(ids...)
	rounds := make([]int, 0)

	states := make(map[string]bool)
	pulses := puz.Deque[Pulse]{}

	for n := 1; n < 10000 && len(needs) > 0; n++ {
		pulses.PushBack(Pulse{Source: "button", Target: "broadcaster", Value: false})
		for !pulses.Empty() {
			pulse := pulses.PopFront()

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

			if needs.Contains(pulse.Target) && state {
				needs.Remove(pulse.Target)
				rounds = append(rounds, n)
			}
		}
	}

	return rounds
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
	ids := puz.Set[string]{}
	kinds := make(map[string]string)
	targets := make(map[string][]string)
	parents := make(map[string][]string)

	for _, line := range puz.InputToLines() {
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
