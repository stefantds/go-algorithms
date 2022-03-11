package union_find

type UFCreator interface {
	// New connects returns an UF item with n integers (values from 0 to n-1)
	New(n int) UF
}

type UF interface {
	// Union connects p and q
	Union(p, q int)

	// Find returns the component to which p belongs
	Find(p int) int

	// Connected returns true if p and q are in the same component
	Connected(p, q int) bool

	// Count returns the number of components
	Count() int
}
