package aoc

// EnumeratePermutations calls the fn callback with each permutation of size
// items.  Fn will be called size factorial times.
func EnumeratePermutations(size int, fn func(perm []int)) {
	perm := make([]int, size)
	for i := 0; i < size; i++ {
		perm[i] = i
	}

	var generate func(k int)
	generate = func(k int) {
		if k == 1 {
			fn(perm)
			return
		}

		generate(k - 1)
		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				perm[i], perm[k-1] = perm[k-1], perm[i]
			} else {
				perm[0], perm[k-1] = perm[k-1], perm[0]
			}

			generate(k - 1)
		}
	}

	generate(size)
}
