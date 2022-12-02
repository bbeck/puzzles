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
	a := make([][]T, width)
	for x := 0; x < width; x++ {
		a[x] = make([]T, height)
	}
	return a
}

// Split partitions a slice into chunks using a partition function.  Elements of
// the slice are passed into the partition function and runs of true return
// values are grouped together into a chunk.
func Split[T any](ts []T, fn func(T) bool) [][]T {
	var partitions [][]T

	var current []T
	for _, t := range ts {
		if !fn(t) {
			partitions = append(partitions, current)
			current = nil
			continue
		}

		current = append(current, t)
	}

	if current != nil {
		partitions = append(partitions, current)
	}

	return partitions
}
