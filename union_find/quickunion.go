package union_find

func NewQuickUnion(n int) UF {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	return quickUnion{
		parents: parents,
	}
}

type quickUnion struct {
	parents []int
}

func (q quickUnion) Union(x, y int) {
	// set the root component of y as the parent of the root component of x
	q.parents[q.root(x)] = q.root(y)
}

func (q quickUnion) Find(x int) int {
	// find the root ancestor of x
	return q.root(x)
}

func (q quickUnion) Connected(x, y int) bool {
	// x and y belong to the same component if they have the same ancestor
	return q.root(x) == q.root(y)
}

func (q quickUnion) root(x int) int {
	rootX := x
	for q.parents[rootX] != rootX {
		rootX = q.parents[rootX]
	}
	return rootX
}

func (q quickUnion) Count() int {
	return len(q.parents)
}
