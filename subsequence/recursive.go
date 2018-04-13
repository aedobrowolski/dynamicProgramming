package subsequence


// RecursiveCount the number of subsequences of one string in another.
func RecursiveCount(zs, xs string) int {
	return Recursive([]rune(zs), []rune(xs))
}

// Recursive returns the number of ways in which the first sequence can
// appear in the second as a subsequence using recursive programming.
func Recursive(z, x []rune) int {
	switch {
	case len(z) == 0:
		return 1
	case len(x) == 0:
		return 0
	default:
		// Add the ways in which the first rune can be satisfied
		var count int
		for i := 0; i < len(x); i++ {
			if x[i] == z[0] {
				count += Recursive(z[1:], x[i+1:])
			}
		}
		return count
	}
}
