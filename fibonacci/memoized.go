package fibonacci

// Memoized computes the nth fibonacci number using memoization.
func Memoized(n int) int {
	f := memoize(SelfRef)
	return f(n, f)
}

// memoize takes a pure self-referential function and returns one that uses a
// cache to avoid repeated computation of the same input.
func memoize(f selfreferential) selfreferential {
	cache := map[int]int{}
	var g selfreferential
	g = func(n int, self selfreferential) int {
		if stored, ok := cache[n]; ok {
			return stored
		}
		cache[n] = f(n, g) // invoke f and store the result
		return cache[n]
	}
	return g
}
