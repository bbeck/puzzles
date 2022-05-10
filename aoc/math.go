package aoc

// Abs returns the absolute value of the provided integer or float.
func Abs[T Integer | Float](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

// Sum adds together each of the provided elements.
func Sum[T Number](elems ...T) T {
	var sum T
	for _, elem := range elems {
		sum += elem
	}
	return sum
}

// Product multiplies together each of the provided elements.
// If no elements are passed in as arguments then 1 is returned.
func Product[T Number](elems ...T) T {
	var product T = 1
	for _, elem := range elems {
		product *= elem
	}
	return product
}

// Min returns the smallest element of the provided elements.
func Min[T Ordered](elems ...T) T {
	min := elems[0]
	for _, elem := range elems[1:] {
		if elem < min {
			min = elem
		}
	}
	return min
}

// Max returns the largest element of the provided elements.
func Max[T Ordered](elems ...T) T {
	max := elems[0]
	for _, elem := range elems[1:] {
		if elem > max {
			max = elem
		}
	}
	return max
}
