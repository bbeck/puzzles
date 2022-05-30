package aoc

// Make2D creates a two-dimensional slice with the passed in dimensions.
func Make2D[T any](width, height int) [][]T {
	a := make([][]T, height)
	for y := 0; y < height; y++ {
		a[y] = make([]T, width)
	}
	return a
}
