package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bbeck/advent-of-code/lib"
)

func main() {
	s := lib.InputToString()

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
			// If this object has a key with a value of "red" then it should be ignored.
			if v == "red" {
				return 0
			}

			sum += Sum(v)
		}
		return sum
	}

	return 0
}
