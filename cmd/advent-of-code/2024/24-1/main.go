package main

import (
	"fmt"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	inputs, rules := InputToInputsAndRules()

	var z int
	for _, rule := range rules {
		var bit int
		bit, inputs = Eval(rule.Output, inputs, rules)

		if strings.HasPrefix(rule.Output, "z") {
			z |= bit << ParseInt(rule.Output[1:])
		}
	}
	fmt.Println(z)
}

func Eval(output string, inputs map[string]int, rules []Rule) (int, map[string]int) {
	if v, ok := inputs[output]; ok {
		return v, inputs
	}

	var rule Rule
	for _, rule = range rules {
		if rule.Output == output {
			break
		}
	}

	var lhs, rhs int
	lhs, inputs = Eval(rule.Arg1, inputs, rules)
	rhs, inputs = Eval(rule.Arg2, inputs, rules)

	switch rule.Op {
	case "AND":
		inputs[output] = lhs & rhs
	case "OR":
		inputs[output] = lhs | rhs
	case "XOR":
		inputs[output] = lhs ^ rhs
	}

	return inputs[output], inputs
}

type Rule struct {
	Output         string
	Arg1, Op, Arg2 string
}

func InputToInputsAndRules() (map[string]int, []Rule) {
	var inputs = make(map[string]int)
	var rules []Rule

	for _, line := range InputToLines() {
		switch {
		case strings.Contains(line, ":"):
			lhs, rhs, _ := strings.Cut(line, ": ")
			inputs[lhs] = ParseInt(rhs)

		case strings.Contains(line, "->"):
			fields := strings.Fields(line)
			rules = append(rules, Rule{
				Output: fields[4],
				Arg1:   fields[0],
				Op:     fields[1],
				Arg2:   fields[2],
			})
		}
	}

	return inputs, rules
}
