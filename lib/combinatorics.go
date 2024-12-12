package lib

import (
	"cmp"
	"slices"
)

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

// UniquePermutations returns a generator of all unique permutations in
// lexiographic order.  If the inputted slice contains duplicate values then
// only a single permutation is returned.
//
// See also:
// https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order
func UniquePermutations[S ~[]E, E cmp.Ordered](items S) func(func(S) bool) {
	slices.Sort(items)
	size := len(items)

	return func(yield func(S) bool) {
		var found bool
		var i, j int

		for {
			// Yield the permutation we have
			if !yield(items) {
				return
			}

			// Find the largest index i such that items[i] < items[i+1]
			found = false
			for i = size - 2; i >= 0; i-- {
				if items[i] < items[i+1] {
					found = true
					break
				}
			}

			// If no such index exists, the permutation is the last one
			if !found {
				return
			}

			// Find the largest index j greater than i such that items[i] < items[j]
			for j = size - 1; j > i; j-- {
				if items[i] < items[j] {
					break
				}
			}

			// Swap the value of items[i] with that of items[j], then reverse the
			// sequence from items[i+1] to the end to form the new permutation
			items[i], items[j] = items[j], items[i]
			Reverse(items[i+1:])
		}
	}
}
