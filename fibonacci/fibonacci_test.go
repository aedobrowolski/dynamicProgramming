package fibonacci_test

import (
	"testing"

	"github.com/adobrowolski/problemSolvers/dynamicProrgramming/fibonacci"
)

var testFunctions = []struct {
	desc string
	f    func(int) int
}{
	{"Dynamic", fibonacci.Dynamic}, 
	{"Memoized", fibonacci.Memoized}, 
	{"Recursive", fibonacci.Recursive}, 
	{"Selfref", fibonacci.Fibonacci},
}

var testCases = []struct {
	desc    string
	in, out int
}{
	{desc: "1",  in: 1,  out: 1},
	{desc: "10", in: 10, out: 55},
	{desc: "20", in: 20, out: 6765},
	{desc: "40", in: 40, out: 102334155},
}

func TestFibonacci(t *testing.T) {
	for _, tC := range testCases {
		for _, tF := range testFunctions {
			t.Run(tF.desc+tC.desc, func(t *testing.T) {
				if got, exp := tF.f(tC.in), tC.out; got != exp {
					t.Errorf("Fib(%v). Got %v but expected %v.", tC.in, got, exp)
				}
			})
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for _, tC := range testCases {
		b.Logf("Fib(%v)\n", tC.in)
		for _, tF := range testFunctions {
			b.Run(tF.desc+tC.desc, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tF.f(tC.in)
				}
			})
		}
	}
}
