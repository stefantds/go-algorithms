package union_find

func NewWeightedQuickUnion(n int) weightedQuickUnion {
	parents := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		size[i] = 1
	}
	return weightedQuickUnion{
		parents: parents,
		size:    size,
	}
}

type weightedQuickUnion struct {
	parents []int
	size    []int
}

func (q weightedQuickUnion) Union(x, y int) {
	rootX := q.root(x)
	rootY := q.root(y)
	if q.size[rootX] < q.size[rootY] {
		q.parents[rootX] = rootY
	} else if q.size[rootX] > q.size[rootY] {
		q.parents[rootY] = rootX
	} else {
		q.parents[rootY] = rootX
		q.size[rootX] += 1
	}
}

func (q weightedQuickUnion) Find(x int) int {
	// find the root ancestor of x
	return q.root(x)
}

func (q weightedQuickUnion) Connected(x, y int) bool {
	// x and y belong to the same component if they have the same ancestor
	return q.root(x) == q.root(y)
}

func (q weightedQuickUnion) root(x int) int {
	if x == q.parents[x] {
		return x
	}
	q.parents[x] = q.root(q.parents[x])
	return q.parents[x]
}

func (q weightedQuickUnion) Count() int {
	return len(q.parents)
}
