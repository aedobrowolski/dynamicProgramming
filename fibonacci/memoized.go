package fibonacci

// recursive defines a type of function that takes a reference to itself
type recursive func(int, recursive) int

// Fibonacci is a recursive implementation of the fibonacci function
func Fibonacci(n int, self recursive) int {
	if n <= 2 {
		return 1
	}
	return self(n-2, self) + self(n-1, self)
}

// Memoized uses a bottom up approach to computing the nth fibonacci number
func Memoized(n int) int {
	f := memoize(Fibonacci)
	return f(n, f)
}

func memoize(f recursive) recursive {
	cache := map[int]int{}
	var g recursive
	g = func(n int, self recursive) int {
		if stored, ok := cache[n]; ok {
			return stored
		}
		cache[n] = f(n, g) // invoke f and store the result
		return cache[n]
	}
	return g
}
