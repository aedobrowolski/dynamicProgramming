package subsequence_test

import (
	"testing"

	"github.com/adobrowolski/problemSolvers/dynamicProrgramming/subsequence"
)

func TestDynamicCount(t *testing.T) {
	if got, exp := subsequence.DynamicCount("bag", "babgbag"), 5; got != exp {
		t.Errorf("DynamicCount(%q,%q): got %v, expected %v.", "bag", "babgbag", got, exp)
	}
}

var testFunctions = []struct {
	desc string
	f    func(z, x string) int
}{
	{"Dynamic", subsequence.DynamicCount},
	{"Recursive", subsequence.RecursiveCount},
	{"Selfref", subsequence.SelfRefCount},
	{"Memoized", subsequence.Memoized},
}

var testCases = []struct {
	desc  string
	z, x  string
	count int
}{
	{desc: "small", z: "bag", x: "babgbag", count: 5},
	{desc: "med-1", z: "bag", x: "babgbagbabgbag", count: 27},
	{desc: "med-2", z: "baag", x: "babgbagbabgbag", count: 19},
	{desc: "med-3", z: "tag", x: "taggcttgaagaaagg", count: 42},
	{desc: "big-1", z: "aaaaaaaa", x: "aaaaaaaaaaaaaaaa", count: 12870},
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		for _, tF := range testFunctions {
			t.Run(tF.desc+"-"+tC.desc, func(t *testing.T) {
				if got, exp := tF.f(tC.z, tC.x), tC.count; got != exp {
					t.Errorf("%s(%q,%q): got %v, expected %v.", tF.desc, tC.z, tC.x, got, exp)
				}
			})
		}
	}
}

func Benchmark(b *testing.B) {
	for _, tC := range testCases {
		for _, tF := range testFunctions {
			b.Run(tF.desc+"-"+tC.desc, func(b *testing.B) {
				for n := 0; n < b.N; n++ {
					tF.f(tC.z, tC.x)
				}
			})
		}
	}
}
