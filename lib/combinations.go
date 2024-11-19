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
