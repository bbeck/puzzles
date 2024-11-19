package lib

// EnumeratePermutations calls the fn callback with each permutation of size
// items.  Fn will be called size factorial times.
func EnumeratePermutations(size int, fn func([]int) bool) {
	perm := make([]int, size)
	for i := 0; i < size; i++ {
		perm[i] = i
	}

	var generate func(k int) bool
	generate = func(k int) bool {
		if k == 1 {
			return fn(perm)
		}

		if generate(k - 1) {
			return true
		}

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}

			if generate(k - 1) {
				return true
			}
		}

		return false
	}

	generate(size)
}
