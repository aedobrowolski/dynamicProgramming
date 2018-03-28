package fibonacci

// Recursive calculates the nth fibonacci number via simple recursion.
func Recursive(n int) int {
	if n <= 2 {
		return 1
	}
	return Recursive(n-2) + Recursive(n-1)
}
