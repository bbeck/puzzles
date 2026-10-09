package main

import (
	"fmt"
	"math/bits"

	. "github.com/bbeck/puzzles/lib"
	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	lights, masks, buttons := InputToLightsMasksButtons()

	var sum int
	for i := range lights {
		sum += Shortest(lights[i], masks[i], buttons[i])
	}
	fmt.Println(sum)
}

func Shortest(target Lights, mask Mask, buttons Buttons) int {
	// We'll only ever push a button a single time, so just try all combinations
	// of an increasing number of button pushes until we find the solution.
	n := len(buttons)
	for k := 1; k < n; k++ {
		var found bool
		EnumerateCombinations(n, k, func(presses []int) bool {
			var state uint64 = 0
			for _, b := range presses {
				state = state ^ buttons[b]
			}
			found = (state & mask) == target
			return found
		})

		if found {
			return k
		}
	}

	panic("no solution")
}

type Lights = uint64
type Buttons = []uint64
type Mask = uint64

func InputToLightsMasksButtons() ([]Lights, []Mask, []Buttons) {
	in.Remove("[", "]", "(", ")", "{", "}")

	toLightsAndMask := func(in in.Scanner[any]) (uint64, uint64) {
		var mask, lights uint64
		for in.HasNext() {
			mask = (mask << 1) | 1

			lights <<= 1
			if ch := in.Byte(); ch == '#' {
				lights |= 1
			}
		}
		return lights, mask
	}

	toButton := func(in in.Scanner[any], N int) uint64 {
		var button uint64
		for _, n := range in.Ints() {
			button |= 1 << (N - 1 - n)
		}
		return button
	}

	var allLights []Lights
	var allMasks []Mask
	var allButtons []Buttons
	for in.HasNext() {
		fields := in.FieldsS[any]()

		var lights, mask = toLightsAndMask(fields[0])
		allLights = append(allLights, lights)
		allMasks = append(allMasks, mask)
		n := bits.OnesCount64(mask)

		var buttons Buttons
		for i := 1; i < len(fields)-1; i++ {
			buttons = append(buttons, toButton(fields[i], n))
		}
		allButtons = append(allButtons, buttons)
	}

	return allLights, allMasks, allButtons
}
