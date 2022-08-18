package trees

import (
	"math"
)

type segmentTree struct {
	vals []int
	n    int
}

func NewSegmentTree(values []int) segmentTree {
	n := len(values)
	tree := make([]int, 2*len(values)) // we need size 2*n to store n values; for simplicity the array is 1-indexed
	copy(tree[n:], values)

	for i := n - 1; i > 0; i-- {
		tree[i] = max(tree[2*i], tree[2*i+1])
	}

	return segmentTree{
		vals: tree,
		n:    n,
	}
}

// RangeQuery returns the max value in [from:to] (to is exclusive)
func (s segmentTree) RangeQuery(from, to int) int {
	from += s.n
	to += s.n
	maxFound := math.MinInt32

	for from < to {
		if from%2 == 1 { // from is the right child. Consider it but don't move up directly, as it would include the previous element
			maxFound = max(maxFound, s.vals[from])
			from++
		}

		if to%2 == 1 {
			to--
			maxFound = max(maxFound, s.vals[to]) // to is exclusive. If it's the right child, consider the left child's value
		}
		from /= 2
		to /= 2
	}

	return maxFound
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
