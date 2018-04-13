package subsequence

// Memoized computes the nth fibonacci number using memoization.
func Memoized(zs, xs string) int {
	f := memoize(SelfRef)
	return f(zs, xs, f)
}

type key struct{ z, x string }

// memoize takes a pure self-referential function and returns one that uses a
// cache to avoid repeated computation of the same input.
func memoize(f selfreferential) selfreferential {
	cache := map[key]int{}
	var g selfreferential
	g = func(z, x string, self selfreferential) int {
		k := key{z, x}
		if stored, ok := cache[k]; ok {
			return stored
		}
		cache[k] = f(z, x, g) // invoke f and store the result
		return cache[k]
	}
	return g
}
