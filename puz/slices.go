package puz

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

// Make3D creates a three-dimensional slice with the passed in dimensions.
func Make3D[T any](width, height, depth int) [][][]T {
	a := make([][][]T, width)
	for x := 0; x < width; x++ {
		a[x] = make([][]T, height)
		for y := 0; y < height; y++ {
			a[x][y] = make([]T, depth)
		}
	}
	return a
}

// Reverse will reverse the order of the elements of a slice in place.
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Reversed returns a copy of the slice with the order of the elements reversed.
func Reversed[T any](s []T) []T {
	t := make([]T, len(s))
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		t[i] = s[j]
	}
	return t
}

// Repeat repeats a slice multiple times.
func Repeat[T any](s []T, n int) []T {
	t := make([]T, 0, n*len(s))
	for n > 0 {
		t = append(t, s...)
		n--
	}
	return t
}

// Split partitions a slice into chunks using a partition function.  Elements of
// the slice are passed into the partition function and runs of true return
// values are grouped together into a chunk.
func Split[T any](ts []T, fn func(T) bool) [][]T {
	var partitions [][]T

	var current []T
	for _, t := range ts {
		if !fn(t) {
			if current != nil {
				partitions = append(partitions, current)
				current = nil
			}
			continue
		}

		current = append(current, t)
	}

	if current != nil {
		partitions = append(partitions, current)
	}

	return partitions
}

// Chunk breaks a slice into chunks with the specified size.  If the number of
// elements in the slice isn't evenly divisible by n, then the last chunk will
// container fewer than n elements.
func Chunk[T any](ts []T, n int) [][]T {
	var chunks [][]T
	for start := 0; start < len(ts); start += n {
		end := Min(len(ts), start+n)
		chunks = append(chunks, ts[start:end])
	}
	return chunks
}

// Identity is a predicate that returns its argument unmodified.
func Identity[T any](t T) T { return t }

// All determines if every element in the specified slice meets a predicate.
func All[T any](elems []T, fn func(T) bool) bool {
	for _, elem := range elems {
		if !fn(elem) {
			return false
		}
	}
	return true
}

// Any determines if any element in the specified slice meets a predicate.
func Any[T any](elems []T, fn func(T) bool) bool {
	for _, elem := range elems {
		if fn(elem) {
			return true
		}
	}
	return false
}
