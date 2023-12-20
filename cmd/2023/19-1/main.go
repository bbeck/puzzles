package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	workflows, parts := InputToWorkflowsAndParts()

	var sum int
	for _, p := range parts {
		if Eval("in", p, workflows) == "A" {
			sum += aoc.Sum(aoc.GetMapValues(p)...)
		}
	}
	fmt.Println(sum)
}

func Eval(current string, part Part, workflows map[string][]Rule) string {
	if current == "A" || current == "R" {
		return current
	}

	var done bool
	for i := 0; i < len(workflows[current]) && !done; i++ {
		rule := workflows[current][i]
		switch rule.Op {
		case "<":
			if value := part[rule.Attribute]; value < rule.Value {
				current = rule.Goto
				done = true
			}
		case ">":
			if value := part[rule.Attribute]; value > rule.Value {
				current = rule.Goto
				done = true
			}
		case "":
			current = rule.Goto
			done = true
		}
	}

	return Eval(current, part, workflows)
}

func InputToWorkflowsAndParts() (map[string][]Rule, []Part) {
	workflows := make(map[string][]Rule)
	parts := make([]Part, 0)

	for _, line := range aoc.InputToLines(2023, 19) {
		if line == "" {
			continue
		}

		if line[0] != '{' {
			id, workflow := ParseWorkflow(line)
			workflows[id] = workflow
		} else {
			parts = append(parts, ParsePart(line))
		}
	}

	return workflows, parts
}

func ParseWorkflow(line string) (string, []Rule) {
	line = strings.ReplaceAll(line, "{", " ")
	line = strings.ReplaceAll(line, "}", " ")
	line = strings.ReplaceAll(line, ",", " ")
	fields := strings.Fields(line)

	var rules []Rule
	for _, s := range fields[1:] {
		s = strings.ReplaceAll(s, "<", " < ")
		s = strings.ReplaceAll(s, ">", " > ")
		s = strings.ReplaceAll(s, ":", " ")
		parts := strings.Fields(s)

		var rule Rule
		switch len(parts) {
		case 1:
			rule.Goto = parts[0]
		case 4:
			rule.Attribute = parts[0]
			rule.Op = parts[1]
			rule.Value = aoc.ParseInt(parts[2])
			rule.Goto = parts[3]
		}

		rules = append(rules, rule)
	}

	return fields[0], rules
}

type Rule struct {
	Attribute string
	Op        string
	Value     int
	Goto      string
}

type Part map[string]int

func ParsePart(s string) Part {
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	s = strings.ReplaceAll(s, "=", " ")
	s = strings.ReplaceAll(s, ",", " ")
	fields := strings.Fields(s)

	part := make(map[string]int)
	for i := 0; i < len(fields); i += 2 {
		part[fields[i]] = aoc.ParseInt(fields[i+1])
	}
	return part
}
