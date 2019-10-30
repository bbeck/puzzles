package aoc

// EnumerateCombinations calls the fn callback with each possible combination
// of k items from a set of n items.  Fn will be called n choose k times.
func EnumerateCombinations(n, k int, fn func([]int)) {
	current := make([]int, k)

	var helper func(int, int)
	helper = func(i, next int) {
		for j := next; j < n; j++ {
			current[i] = j
			if i == k-1 {
				fn(current)
				continue
			}

			helper(i+1, j+1)
		}
	}

	helper(0, 0)
}
