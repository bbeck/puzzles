package lib

// EnumerateCombinations calls the fn callback with each possible combination
// of k items from a set of n items.  Fn will be called n choose k times.  If
// fn returns true then the enumeration will stop.
func EnumerateCombinations(n, k int, fn func([]int) bool) {
	var done bool
	current := make([]int, k)

	var helper func(int, int) bool
	helper = func(i, next int) bool {
		for j := next; j < n && !done; j++ {
			current[i] = j
			if i == k-1 {
				if fn(current) {
					done = true
					return true
				}
				continue
			}

			if helper(i+1, j+1) {
				done = true
				return true
			}
		}

		return false
	}

	helper(0, 0)
}

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

// EnumerateChoices calls the fn callback with each possible choice of k items
// from a set of n items.
func EnumerateChoices(n, k int, fn func([]int) bool) {
	var done bool
	choices := make([]int, k)

	var helper func(int)
	helper = func(d int) {
		if d == k {
			done = fn(choices)
			return
		}

		for v := 0; !done && v < n; v++ {
			choices[d] = v
			helper(d + 1)
		}
	}

	helper(0)
}
