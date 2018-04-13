package subsequence

// selfreferential defines a type of function that takes a reference to itself
type selfreferential func(string, string, selfreferential) int

// SelfRefCount returns the number of ways in which the first sequence can
// appear in the second as a subsequence using recursive programming.
func SelfRefCount(zs, xs string) int {
	return SelfRef(zs, xs, SelfRef)
}

// SelfRef is a self-referential implementation of the subsequence function
func SelfRef(z, x string, self selfreferential) int {
	switch {
	case len(z) == 0:
		return 1
	case len(x) == 0:
		return 0
	default:
		// Add the ways in which the first letter can be satisfied
		var count int
		for i := 0; i < len(x); i++ {
			if x[i] == z[0] {
				count += self(z[1:], x[i+1:], self)
			}
		}
		return count
	}
}
