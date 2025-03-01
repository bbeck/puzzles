package main

import (
	"encoding/json"
	"fmt"

	"github.com/bbeck/puzzles/lib/in"
)

func main() {
	var tree any
	json.Unmarshal(in.Bytes(), &tree)
	sum := Walk(tree)
	fmt.Println(sum)
}

func Walk(root any) int {
	switch elem := root.(type) {
	case float64:
		return int(elem)

	case []any:
		var sum int
		for i := range elem {
			sum += Walk(elem[i])
		}
		return sum

	case map[string]any:
		var sum int
		for _, v := range elem {
			// If this object has a key with a value of "red" then it should be ignored.
			if v == "red" {
				return 0
			}

			sum += Walk(v)
		}
		return sum
	}

	return 0
}
