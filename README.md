# Dynamic Programming

I am glad that you are interested in learning about dynamic programming. The
first thing you need to learn is when to recognize that a problem can be solved
using DP techniques. DP applies to problem domains where the problem can be
partitioned into sub-problems and the ideal solution can be constructed from
ideal solutions of the sub-problems. This is referred to as the **optimal
substructure** property. DP is computationally more efficient than simple
recursion when there is overlap between these sub-problems.

Many problems can be solved by either DP or memoization; knowing which to apply
is also part of the art. The answer depends on the structure of the solution:
if all sub-problems need to be computed then DP with a bottom up solution will
have better performance since it uses no recursion and no indirect lookups
through a hash map. Other considerations to keep in mind are the need for a
hash function (needed by memoization but not by DP) and the lifetime of the
cache (e.g. can the table/cache be used for more than one problem). See the
Wikipedia articles on [Memoization][1] and [Dynamic Programming][2] for more
details.

An important sub-class of DP problems uses recurrence relationships to count
solutions to a problem, without finding an optimal solution. The recurrence
relationship has the property that at each step the number of solutions is a
mathematical combination of the number of solutions to the set of next simplest
sub-problems.

Probably the best way to learn dynamic programming is to read through an
example until your feel comfortable with it. Then go off to implement your own.
The list in the wiki is pretty long but some of these problems require a lot of
additional domain knowledge to set up. The egg dropping puzzle or the matrix
chain multiplication problem might be good starting points.

In this repository I have implemented solutions to some classical dynamic
programming problems. I hope you find them helpful in your studies.

[1]: http://en.wikipedia.org/wiki/Memoization
[2]: http://en.wikipedia.org/wiki/Dynamic_programming
