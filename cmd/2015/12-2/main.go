package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/aoc"
)

func main() {
	s := aoc.InputToString(2015, 12)

	var tree interface{}
	if err := json.Unmarshal([]byte(s), &tree); err != nil {
		log.Fatalf("error parsing json: %+v", err)
	}

	sum := Sum(tree)
	fmt.Printf("sum: %d\n", sum)
}

func Sum(root interface{}) int {
	switch elem := root.(type) {
	case float64:
		return int(elem)

	case string:
		return 0

	case []interface{}:
		var sum int
		for i := 0; i < len(elem); i++ {
			sum += Sum(elem[i])
		}
		return sum

	case map[string]interface{}:
		// If this object has a key with a value of "red" then it should be ignored.
		for _, v := range elem {
			if v == "red" {
				return 0
			}
		}

		var sum int
		for _, v := range elem {
			sum += Sum(v)
		}
		return sum

	default:
		log.Fatalf("unsupported type: %T", elem)
	}

	return 0
}
