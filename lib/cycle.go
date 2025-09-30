package lib

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

// FindCycle steps through a chain of functions looking for a cycle.
// When a cycle is detected, the states leading up to the cycle are returned
// along with the states within the cycle.
func FindCycle[T comparable](start T, next func(T) T) ([]T, []T) {
	return FindCycleWithIdentity(start, next, Identity[T])
}

// FindCycleWithIdentity behaves identically to FindCycle but uses the provided
// ID function allowing the node type to not be comparable.
func FindCycleWithIdentity[T any, I comparable](start T, next func(T) T, ID IDFunc[T, I]) ([]T, []T) {
	current := start

	var visited []T
	seen := make(map[I]int)

	var id I
	for {
		id = ID(current)
		if _, found := seen[id]; found {
			break
		}

		visited = append(visited, current)
		seen[id] = len(seen)

		current = next(current)
	}

	begin := seen[id]
	return visited[:begin], visited[begin:]
}
