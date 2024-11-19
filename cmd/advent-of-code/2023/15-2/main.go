package main

import (
	"fmt"
	"github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	var boxes [256][]Lens

	find := func(box int, label string) (int, bool) {
		for i, lens := range boxes[box] {
			if lens.Label == label {
				return i, true
			}
		}

		return 0, false
	}

	for _, op := range InputToOperations() {
		box := op.Hash()
		index, found := find(box, op.Label)

		switch op.Kind {
		case "-":
			if found {
				boxes[box] = append(boxes[box][:index], boxes[box][index+1:]...)
			}

		case "=":
			lens := Lens{Label: op.Label, FocalLength: op.FocalLength}

			if found {
				boxes[box][index] = lens
			} else {
				boxes[box] = append(boxes[box], lens)
			}
		}
	}

	var power int
	for b, box := range boxes {
		for l, lens := range box {
			power += (b + 1) * (l + 1) * lens.FocalLength
		}
	}
	fmt.Println(power)
}

type Lens struct {
	Label       string
	FocalLength int
}

type Operation struct {
	Kind        string
	Label       string
	FocalLength int
}

func (o Operation) Hash() int {
	var hash int32
	for _, c := range o.Label {
		hash = 17 * (hash + c) % 256
	}
	return int(hash)
}

func InputToOperations() []Operation {
	line := lib.InputToString()
	parts := strings.Split(line, ",")

	var ops []Operation
	for _, part := range parts {
		var op Operation
		if label, _, ok := strings.Cut(part, "-"); ok {
			op.Kind = "-"
			op.Label = label
		}

		if label, length, ok := strings.Cut(part, "="); ok {
			op.Kind = "="
			op.Label = label
			op.FocalLength = lib.ParseInt(length)
		}

		ops = append(ops, op)
	}

	return ops
}
