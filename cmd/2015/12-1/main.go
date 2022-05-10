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
	fmt.Println(sum)
}

func Sum(root interface{}) int {
	switch elem := root.(type) {
	case float64:
		return int(elem)

	case []interface{}:
		var sum int
		for i := 0; i < len(elem); i++ {
			sum += Sum(elem[i])
		}
		return sum

	case map[string]interface{}:
		var sum int
		for _, v := range elem {
			sum += Sum(v)
		}
		return sum
	}

	return 0
}
