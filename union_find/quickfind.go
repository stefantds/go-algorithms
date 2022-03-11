package union_find

func NewQuickFind(n int) UF {
	components := make([]int, n)
	for i := 0; i < n; i++ {
		components[i] = i
	}
	return quickFind{
		components: components,
	}
}

type quickFind struct {
	components []int
}

func (q quickFind) Union(x, y int) {
	if q.Connected(x, y) {
		return
	}

	// set all the elements having the same component as x
	// to belong to the component of y
	cx := q.components[x]
	for i, c := range q.components {
		if c == cx {
			q.components[i] = q.components[y]
		}
	}
}

func (q quickFind) Find(x int) int {
	return q.components[x]
}

func (q quickFind) Connected(x, y int) bool {
	return q.components[x] == q.components[y]
}

func (q quickFind) Count() int {
	return len(q.components)
}
