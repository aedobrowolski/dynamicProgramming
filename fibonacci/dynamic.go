package fibonacci

// Dynamic uses a bottom up approach to computing the nth fibonacci number
func Dynamic(n int) int {
	a, b := 1, 1
	for i := 2; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
