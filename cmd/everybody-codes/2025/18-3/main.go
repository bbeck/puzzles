package main

import (
	"bytes"
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	plants, root, tests := InputToPlantsAndTestCases()
	best := Energy(root, plants, BuildMaximumTestCase(plants, len(tests[0])))

	var sum int
	for _, test := range tests {
		if energy := Energy(root, plants, test); energy > 0 {
			sum += best - energy
		}
	}
	fmt.Println(sum)
}

func BuildMaximumTestCase(plants map[int]Plant, size int) []int {
	// Analyzing the input reveals that branches are always used with either a
	// negative or positive thickness, but not a mix of both.  Therefore, we can
	// just cancel out the negative thickness ones by using a 0 in the test case
	// and keep the positive thickness ones by using a 1 in the test case.
	var test = make([]int, size)

	for _, p := range plants {
		for id, thickness := range p.Children {
			if id <= size && thickness > 0 {
				test[id-1] = 1
			}
		}
	}

	return test
}

func Energy(id int, plants map[int]Plant, test []int) int {
	if plants[id].Kind == "free" {
		return test[id-1]
	}

	var sum int
	for cid, thickness := range plants[id].Children {
		sum += Energy(cid, plants, test) * thickness
	}

	if sum >= plants[id].Thickness {
		return sum
	}
	return 0
}

type Plant struct {
	ID        int
	Thickness int
	Kind      string
	Children  map[int]int
}

func InputToPlantsAndTestCases() (map[int]Plant, int, [][]int) {
	lhs, rhs, _ := bytes.Cut(in.Bytes(), []byte{'\n', '\n', '\n'})

	plants, lastPlantID := InputToPlants(lhs)
	cases := InputToTestCases(rhs)
	return plants, lastPlantID, cases
}

func InputToPlants(in in.Scanner[any]) (map[int]Plant, int) {
	var plants = make(map[int]Plant)
	var lastPlantID int

	for in.HasNext() {
		chunk := in.ChunkS()

		var plant = Plant{Children: make(map[int]int)}
		plant.ID, plant.Thickness = chunk.Int(), chunk.Int()
		chunk.Expect(":\n")

		for chunk.HasNext() {
			switch {
			case chunk.HasPrefix("- free"):
				plant.Kind = "free"
				chunk.Expect("- free branch with thickness 1")

			case chunk.HasPrefix("- branch"):
				plant.Kind = "branch"
				id, thickness := chunk.Int(), chunk.Int()
				plant.Children[id] = thickness
				if chunk.HasNext() {
					chunk.Expect("\n")
				}
			}
		}

		plants[plant.ID] = plant
		lastPlantID = plant.ID
	}

	return plants, lastPlantID
}

func InputToTestCases(s in.Scanner[[]int]) [][]int {
	return s.LinesToS(func(s in.Scanner[[]int]) []int {
		return s.Ints()
	})
}
