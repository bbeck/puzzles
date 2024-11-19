package main

import (
	"fmt"
	"strings"

	"github.com/bbeck/puzzles/lib"
)

// 167409079868000 expected
func main() {
	workflows := InputToWorkflows()

	part := Part{
		X: Interval{Start: 1, End: 4000},
		M: Interval{Start: 1, End: 4000},
		A: Interval{Start: 1, End: 4000},
		S: Interval{Start: 1, End: 4000},
	}

	var count int
	Eval("in", part, workflows, func(p Part) {
		count += p.Size()
	})
	fmt.Println(count)
}

func Eval(current string, part Part, workflows map[string][]Rule, fn func(Part)) {
	if current == "A" {
		fn(part)
		return
	}

	if current == "R" {
		return
	}

	var done bool
	for i := 0; i < len(workflows[current]) && !done; i++ {
		rule := workflows[current][i]
		switch rule.Op {
		case "<":
			success := part.Intersect(rule.Attribute, Interval{Start: 1, End: rule.Value - 1})
			Eval(rule.Goto, success, workflows, fn)
			part = part.Intersect(rule.Attribute, Interval{Start: rule.Value, End: 4000})

		case ">":
			success := part.Intersect(rule.Attribute, Interval{Start: rule.Value + 1, End: 4000})
			Eval(rule.Goto, success, workflows, fn)
			part = part.Intersect(rule.Attribute, Interval{Start: 1, End: rule.Value})

		case "":
			Eval(rule.Goto, part, workflows, fn)
		}
	}
}

type Part struct {
	X, M, A, S Interval
}

func (p Part) Intersect(s string, i Interval) Part {
	switch s {
	case "x":
		p.X = p.X.Intersect(i)
	case "m":
		p.M = p.M.Intersect(i)
	case "a":
		p.A = p.A.Intersect(i)
	case "s":
		p.S = p.S.Intersect(i)
	}
	return p
}

func (p Part) Size() int {
	return 1 *
		(p.X.End - p.X.Start + 1) *
		(p.M.End - p.M.Start + 1) *
		(p.A.End - p.A.Start + 1) *
		(p.S.End - p.S.Start + 1)
}

type Interval struct {
	Start int
	End   int
}

func (i Interval) Intersect(j Interval) Interval {
	return Interval{
		Start: lib.Max(i.Start, j.Start),
		End:   lib.Min(i.End, j.End),
	}
}

func InputToWorkflows() map[string][]Rule {
	workflows := make(map[string][]Rule)
	for _, line := range lib.InputToLines() {
		if line == "" {
			continue
		}

		if line[0] != '{' {
			id, workflow := ParseWorkflow(line)
			workflows[id] = workflow
		}
	}

	return workflows
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
			rule.Value = lib.ParseInt(parts[2])
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
