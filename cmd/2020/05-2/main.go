package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	var ids []int
	for _, line := range puz.InputToLines(2020, 5) {
		ids = append(ids, puz.ParseIntWithBase(replacer.Replace(line), 2))
	}

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i]+1 != ids[i+1] {
			fmt.Println(ids[i] + 1)
		}
	}
}
