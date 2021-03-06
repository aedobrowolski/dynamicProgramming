# Fibonacci performance tests

The tests in this folder compare four different implementations of a function
that computes the nth fibonacci number. Dynamic uses a bottom up approach to
compute the value. Memoized is a top down recursive approach, but all results
are stored in a map and looked up if subsequently needed. Recursive does not
cache any results. Self-referential uses the same algorithm as recursive but
has the function to call passed in, allowing it to be memoized.

The tests are in `fibonacci_test.go` while `Results.txt` shows the results of the
experiment.

```text
$ go test -bench .

BenchmarkFibonacci/Dynamic1-8         	300000000	         4.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Memoized1-8        	 5000000	       373 ns/op	     232 B/op	       4 allocs/op
BenchmarkFibonacci/Recursive1-8       	300000000	         4.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Selfref1-8         	200000000	         6.54 ns/op	       0 B/op	       0 allocs/op

BenchmarkFibonacci/Dynamic10-8        	100000000	        12.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Memoized10-8       	 1000000	      1849 ns/op	     524 B/op	       5 allocs/op
BenchmarkFibonacci/Recursive10-8      	 5000000	       329 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Selfref10-8        	 5000000	       385 ns/op	       0 B/op	       0 allocs/op

BenchmarkFibonacci/Dynamic20-8        	50000000	        24.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Memoized20-8       	  300000	      4560 ns/op	    1181 B/op	       7 allocs/op
BenchmarkFibonacci/Recursive20-8      	   30000	     40833 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Selfref20-8        	   30000	     47666 ns/op	       0 B/op	       0 allocs/op

BenchmarkFibonacci/Dynamic40-8        	30000000	        53.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Memoized40-8       	  200000	     10125 ns/op	    2522 B/op	      11 allocs/op
BenchmarkFibonacci/Recursive40-8      	       2	 618500000 ns/op	       0 B/op	       0 allocs/op
BenchmarkFibonacci/Selfref40-8        	       2	 720500000 ns/op	       0 B/op	       0 allocs/op
PASS
```

These represent calls to compute Fib(1), Fib(10), Fib(20) and Fib(40)
respectively, repeated a large number of times. No global data is stored
between invocations. The two pure recursive implementations show exponential
growth, as expected. The dynamic solution beats the memoized by a large margin
(about 200 times faster) even though both do exactly the same number of additions.
This is due to the elimination of the recursive function calls still needed by
the memoized function, the computation of a hash for the inputs, and a cache
lookup.