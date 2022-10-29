package aoc

// GetMapKeys returns the keys from the provided map.
func GetMapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GetMapValues returns the values from the provided map.
func GetMapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Make2D creates a two-dimensional slice with the passed in dimensions.
func Make2D[T any](width, height int) [][]T {
	a := make([][]T, height)
	for y := 0; y < height; y++ {
		a[y] = make([]T, width)
	}
	return a
}
