package aoc

// WalkCycle steps through a list a large number of times.  When a cycle is
// detected it is used to skip a large number of iterations.
func WalkCycle[T comparable](root T, N int, next func(T) T) T {
	return WalkCycleWithIdentity(root, N, next, Identity[T])
}

// WalkCycleWithIdentity behaves identically to WalkCycle but uses the provided
// ID function allowing the node type to not be comparable.
func WalkCycleWithIdentity[T any, I comparable](start T, N int, next func(T) T, ID IDFunc[T, I]) T {
	current := start

	seen := make(map[I]int)

	var id I
	var n int
	for n = 0; n < N; n++ {
		id = ID(current)
		if _, ok := seen[id]; ok {
			break
		}

		next := next(current)
		seen[id] = n
		current = next
	}

	length := n - seen[id]
	for remaining := (N - n) % length; remaining > 0; remaining-- {
		current = next(current)
	}

	return current
}
