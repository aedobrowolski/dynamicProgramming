# Sub-sequence counting

The problem proposed here is from an old ACM programming competition. It is
meant as a learning tool to study dynamic programming.

> A subsequence of a given sequence $S$ consists of $S$ with zero or more elements
> deleted. Formally a sequence $Z=z_1 z_2...z_k$ is a subsequence of $X=x_1 x_2...x_m$ if
> there exists a strictly increasing sequence $<i_1,i_2,...,i_k>$ of indices of $X$
> such that for all $j=1,2,...,k$ we have $x_{i_j}=z_j$. The problem is to write a
> program that counts the number of occurrences of $Z$ in $X$ as a subsequence such
> that each has a distinct index sequence.

For example, the word `bag` can appear as a subsequence of `babgbag` in five
different ways, as illustrated below. The asterisks indicate matches of the
symbols in the first string within the second.

```text
babgbag
** *
**    *
*    **
  *  **
    ***
```

## Simplification

Observe that if the symbols from the sequence $X$ which do not occur in the
sequence $Z$ are removed from $X$ to form a new sequence $X'$ the count will
remain the same. Preprocessing $X$ to remove unused symbols can improve
performance (searching for a match) without affecting the result.

## Recursion

We are being asked to count how many sequences $<i_1,i_2,...,i_k>$ of indices
of $X$ match the constraints:

* the $i_j$ are strictly increasing, and
* for all $j=1,2,...,k$ we have $x_{i_j}=z_j$.

Let's call this count $C(Z,X)$ and define $C(\emptyset,X)=1$ for all $X$.

We will use the notation $Z_j=z_1 z_2...z_j$ for $j<=k$ and likewise for
$X_i=x_1 x_2...x_i$ for $i<=m$ to represent the possible initial sequences of
$X$ and $Z$ respectively. We define $X_0=Z_0=\emptyset$, the empty sequence.
With this notation we can write $C(Z,X)=C(Z_k, X_m)$ assuming $k=|Z|, m=|X|$.

We see that our counting problem can be solved by recursion by observing that
if $x_{i_k}=z_k$, then the number of sequences of indices ending in $i_k$ is
equal to $C(Z_{k-1},X_{i_k-1})$. This is a simpler sub-problem of the original
problem. The number of ways in which $i_k$ can be chosen so that $x_{i_k}=z_k$
is precisely the number of occurrences of the symbol $z_k$ in $X$.  Thus the
solution is expressed in the recurrence:

$$C(Z_k,X_m) = \sum_{i_k \in\{i_k | x_{i_k}=z_k\}} C(Z_{k-1},X_{i_k-1})$$

The number of matches $x_{i_k}=z_k$ may be zero in which case the sum is zero.
The recurrence also ends if $k=0$ since $Z_0=\emptyset$ so $C(Z_k,X_m) = 1$.

Using our example from above where Z=`bag` and X=`babgbag` we can work out a
tree of recurrences. The notation `C(bag,babgbag)` signifies the number of ways
in which the string `bag` can be embedded in `babgbag` as a subsequence.
Expanding this expression is done by matching the last letter in the first
argument with the occurrences of that letter in the right, dropping the
matching letters and all following letters. The nodes `C(∅,...)` are leaf nodes
and so are not expanded.

```text
C(bag,babgbag)  = C(ba,babgba) + C(ba,bab)
C(ba,babgba)    = C(b,babgb) + C(b,b)
C(b, babgb)     = C(∅,babg) + C(∅,ba) + C(∅,∅)
C(b,b)          = C(∅,∅)
C(ba,bab)       = C(b,b)
C(b,b)          = C(∅,∅)
```

Substituting values computed at the leaf nodes of the tree (all 1) up the
recursion stack we arrive at the final answer.

```text
C(bag,babgbag)  = C(ba,babgba) + C(ba,bab)      = 5
C(ba,babgba)    = C(b,babgb) + C(b,b)           = 4
C(b,babgb)      = C(∅,babg) + C(∅,ba) + C(∅,∅)  = 3
C(b,b)          = C(∅,∅)                        = 1
C(ba,bab)       = C(b,b)                        = 1
C(b,b)          = C(∅,∅)                        = 1
```

## Memoization

The implementation of this recursive solution can be made more efficient by
memoization since the counts $C(Z_k,X_m)$ may occur in multiple recurrences. In
our example, the value `C(b,b)` was computed twice. In a more complex example
the repeated calculations would dominate the performance. Since we are building
a tree the performance is exponential in the worst case.

The number of cache entries will be O($|Z||X|$). It should be much less in
practice since only indexes of $X$ prior to a matching letter will be
referenced. As a result, memoization will improve the performance.

## Dynamic Programming

Fortunately this problem has the *optimal sub-problem* property. This allows us
to use dynamic programming to build the solution from the bottom up. We do this
by iterating on the indexes in $Z$ and matching them against the indexes in
$X$.

To store the partial counts we create a $k+1$ by $m+1$ table. Entry $(i,j)$ is
equal to $C(Z_i,X_j)$. All entries in the first row, corresponding to
`C(∅,...)`, are initialized to 1. The remaining entries in the first column,
corresponding to `C(...,∅)`, are initialized to 0. All other entries get the
value to their left, unless the symbols match, in which case they get the sum
of the value on the left and the value in the cell north-east of the entry.

![Subsequence Tableau](./DynamicTableau.png "Tableau for subsequence matches")

The relationship is expressed in this formula:

$$
T[i,j] = \begin{cases}
   T[i-1,j] &\text{if } Z[i] \ne X[j] \\
   T[i-1,j] + T[i-1,j-1] &\text{if } Z[i] = X[j]
\end{cases}
$$

The table can be filled in either by row or by column. An interesting
observation is that only two rows (or two columns) need to be kept around. In
fact, only one row/column is needed if the evaluations are done from high to
low index. Since $|Z|\leq|X|$ for a non-zero result, storing a column only
leads to a very elegant solution with O($|Z|$) storage overhead.  This is
much better than the memoized solution.