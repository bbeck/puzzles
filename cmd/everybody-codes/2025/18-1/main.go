package main

import (
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	plants, root := InputToPlants()
	fmt.Println(Energy(root, plants))
}

func Energy(id int, plants map[int]Plant) int {
	if plants[id].Kind == "free" {
		return 1
	}

	var sum int
	for cid, thickness := range plants[id].Children {
		sum += Energy(cid, plants) * thickness
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

func InputToPlants() (map[int]Plant, int) {
	var plants = make(map[int]Plant)

	var lastID int
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
		lastID = plant.ID
	}

	return plants, lastID
}
