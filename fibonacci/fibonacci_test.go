package fibonacci_test

import (
	"testing"

	"github.com/adobrowolski/problemSolvers/dynamicProrgramming/fibonacci"
)

func BenchmarkRecursive10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Recursive(10)
	}
}

func BenchmarkRecursive20(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Recursive(20)
	}
}

func BenchmarkRecursive40(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Recursive(40)
	}
}

func BenchmarkDynamic10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Dynamic(10)
	}
}

func BenchmarkDynamic20(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Dynamic(20)
	}
}

func BenchmarkDynamic40(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Dynamic(40)
	}
}

func BenchmarkMemoized10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Memoized(10)
	}
}

func BenchmarkMemoized20(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Memoized(20)
	}
}

func BenchmarkMemoized40(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fibonacci.Memoized(40)
	}
}
