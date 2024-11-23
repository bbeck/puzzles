package main

import (
	"container/list"
	"fmt"
	. "github.com/bbeck/puzzles/lib"
	"strings"
)

func main() {
	cols := InputToColumns()
	N := len(cols)

	for round := 1; round <= 10; round++ {
		clapper := cols[(round-1)%N].Remove(cols[(round-1)%N].Front()).(int)

		currentCol := round % N
		column := cols[round%N]

		// Remove complete cycles remembering that the clapper goes all the way
		// down and then back up a column.
		steps := (clapper - 1) % (2 * column.Len())

		// Determine which end we should start clapping from and remove any need to
		// turn around.
		var down bool
		var cursor *list.Element
		if steps < column.Len() {
			down = true
			cursor = column.Front()
		} else {
			down = false
			cursor = column.Back()
			steps -= column.Len()
		}

		// Clap the right number of times, we'll never have to worry about switching
		// directions since we've removed cycles and turns.
		for steps > 0 {
			if down {
				cursor = cursor.Next()
			} else {
				cursor = cursor.Prev()
			}
			steps--
		}

		// If the clapper was moving down on the last clap then they are absorbed
		// before the current person.  Otherwise, they are absorbed after.
		if down {
			cols[currentCol].InsertBefore(clapper, cursor)
		} else {
			cols[currentCol].InsertAfter(clapper, cursor)
		}
	}
	fmt.Println(Value(cols))
}

func InputToColumns() []list.List {
	cols := make([]list.List, 4)

	for _, line := range InputToLines() {
		for i, s := range strings.Fields(line) {
			cols[i].PushBack(ParseInt(s))
		}
	}

	return cols
}

func Value(columns []list.List) int {
	var sb strings.Builder
	for _, c := range columns {
		sb.WriteString(fmt.Sprintf("%d", c.Front().Value.(int)))
	}

	return ParseInt(sb.String())
}

func PrintCols(header string, columns []list.List) {
	fmt.Println(header)

	var elems []*list.Element
	for _, column := range columns {
		elems = append(elems, column.Front())
	}

	for {
		var ok bool
		for i := range elems {
			if elems[i] == nil {
				fmt.Print("  ")
				continue
			}

			fmt.Print(elems[i].Value.(int), " ")
			elems[i] = elems[i].Next()
			ok = true
		}
		fmt.Println()

		if !ok {
			break
		}
	}
}
