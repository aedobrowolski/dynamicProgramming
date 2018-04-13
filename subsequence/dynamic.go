package subsequence

// DynamicCount the number of subsequences of one string in another.
func DynamicCount(zs, xs string) int {
	return Dynamic([]rune(zs), []rune(xs))
}

// Dynamic returns the number of ways in which the first sequence can
// appear in the second as a subsequence using dynamic programming.
func Dynamic(z, x []rune) int {
	count := make([]int, len(z)+1)
	count[0] = 1
	for i := 0; i < len(x); i++ {
		for j := len(z); j > 0; j-- {
			if x[i] == z[j-1] {
				count[j] += count[j-1]
			}
		}
	}
	return count[len(z)]
}
