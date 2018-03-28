package fibonacci

// selfreferential defines a type of function that takes a reference to itself
type selfreferential func(int, selfreferential) int

// SelfRef is a self-referential implementation of the fibonacci function
func SelfRef(n int, self selfreferential) int {
	if n <= 2 {
		return 1
	}
	return self(n-2, self) + self(n-1, self)
}

// Fibonacci returns the nth fibonacci number.
// It uses a self-referential implementation internally.
func Fibonacci(n int) int {
	return SelfRef(n, SelfRef)
}
