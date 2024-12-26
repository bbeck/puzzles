package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	. "github.com/bbeck/puzzles/lib"
)

func main() {
	inputs, rules := InputToInputsAndRules()

	var swaps []string
	var cin string

	// The equations of a half adder are:
	//   zn = xn XOR yn
	//   cout = xn AND yn
	{
		rS := FindRule(rules, "XOR", "x00", "y00")
		if rS.Output != "z00" {
			swaps = append(swaps, rS.Output, "z00")
			Swap(rules, rS.Output, "z00")
		}
		rC := FindRule(rules, "AND", "x00", "y00")
		cin = rC.Output
	}

	// The equations of a full adder are:
	//   S = (x XOR y) XOR cin
	//   C = (x AND y) OR (cin AND (x XOR y))
	//
	// Thus for each stage with a full adder we expect to find the following
	// intermediate terms.  Note, for the S rule we associate (x XOR y) together
	// since that lets there be reuse of that term in the C rule.
	//   r1 = xn XOR yn
	//   rS = r1 XOR cin
	//
	//   r2 = xn AND yn
	//   r3 = cin AND r1
	//   rC = r2 OR r3
	for n := 1; n <= 44; n++ {
		xn := fmt.Sprintf("x%02d", n)
		yn := fmt.Sprintf("y%02d", n)
		zn := fmt.Sprintf("z%02d", n)

		rS := FindRule(rules, "XOR", cin)
		if rS.Output != zn {
			// rS's output is wrong because it should be zn
			swaps = append(swaps, rS.Output, zn)
			Swap(rules, rS.Output, zn)
		}

		r1 := FindRule(rules, "XOR", xn, yn)
		if rS.Arg1 != r1.Output && rS.Arg2 != r1.Output {
			// r1's output is wrong because it's not used by rS
			if rS.Arg1 != cin {
				swaps = append(swaps, r1.Output, rS.Arg1)
				Swap(rules, r1.Output, rS.Arg1)
			} else {
				swaps = append(swaps, r1.Output, rS.Arg2)
				Swap(rules, r1.Output, rS.Arg2)
			}
		}

		r2 := FindRule(rules, "AND", xn, yn)
		r3 := FindRule(rules, "AND", cin)
		rC := FindRule(rules, "OR", r2.Output, r3.Output)

		// These remaining rules don't have any validations because there's really
		// not anything that can go wrong with them.  If r2 and r3 were swapped we
		// wouldn't notice because they're only used in rC in a commutative fashion.
		// rC can't really be swapped because it's an unnamed output.

		// Propagate our carry to the next iteration of the loop.
		cin = rC.Output
	}

	// Make sure the adder emits the correct number.
	for _, rule := range rules {
		_, inputs = Eval(rule.Output, inputs, rules)
	}

	var x, y, z int
	for input, bit := range inputs {
		switch input[0] {
		case 'x':
			x |= bit << ParseInt(input[1:])
		case 'y':
			y |= bit << ParseInt(input[1:])
		case 'z':
			z |= bit << ParseInt(input[1:])
		}
	}
	if x+y != z {
		log.Fatalf("%d + %d != %d", x, y, z)
	}

	// Output our solution
	sort.Strings(swaps)
	fmt.Println(strings.Join(swaps, ","))
}

func FindRule(rules []*Rule, op string, args ...string) *Rule {
	for _, rule := range rules {
		if rule.Op != op {
			continue
		}

		s := SetFrom(args...).DifferenceElems(rule.Arg1, rule.Arg2)
		if len(s) == 0 {
			return rule
		}
	}

	return nil
}

func Swap(rules []*Rule, out1, out2 string) {
	var idx1, idx2 int

	for i := 0; i < len(rules); i++ {
		switch rules[i].Output {
		case out1:
			idx1 = i
		case out2:
			idx2 = i
		}
	}

	rules[idx1].Output = out2
	rules[idx2].Output = out1
}

func Eval(output string, inputs map[string]int, rules []*Rule) (int, map[string]int) {
	if v, ok := inputs[output]; ok {
		return v, inputs
	}

	var rule *Rule
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

func InputToInputsAndRules() (map[string]int, []*Rule) {
	var inputs = make(map[string]int)
	var rules []*Rule

	for _, line := range InputToLines() {
		switch {
		case strings.Contains(line, ":"):
			lhs, rhs, _ := strings.Cut(line, ": ")
			inputs[lhs] = ParseInt(rhs)

		case strings.Contains(line, "->"):
			fields := strings.Fields(line)
			rules = append(rules, &Rule{
				Output: fields[4],
				Arg1:   fields[0],
				Op:     fields[1],
				Arg2:   fields[2],
			})
		}
	}

	return inputs, rules
}
