package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/bbeck/advent-of-code/puz"
)

func main() {
	time, distance := InputToTimeAndDistance()

	// We want distance - hold*(time - hold) < 0, thus we can solve for hold
	// using the quadratic formula. hold^2 - time*hold + distance < 0
	D := math.Sqrt(time*time - 4*distance)
	h1 := int(math.Ceil(time/2 - D/2))
	h2 := int(math.Floor(time/2 + D/2))

	// When the discriminant is an integer then we have a case where we tied
	// the record.  We need to adjust the roots in order to avoid the tie.
	if D == float64(int(D)) {
		h1 += 1
		h2 -= 1
	}

	fmt.Println(h2 - h1 + 1)
}

func InputToTimeAndDistance() (float64, float64) {
	var time, distance float64
	for _, line := range puz.InputToLines() {
		var sb strings.Builder
		for _, field := range strings.Fields(line)[1:] {
			sb.WriteString(field)
		}

		num := float64(puz.ParseInt(sb.String()))
		if strings.HasPrefix(line, "Time") {
			time = num
		} else {
			distance = num
		}
	}

	return time, distance
}
