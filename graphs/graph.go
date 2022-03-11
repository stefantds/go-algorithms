package graphs

type Graph []*Node

type Node struct {
	Value int
	Edges map[*Node]bool
}

func NewGraph(size int) Graph {
	g := make([]*Node, size)
	for i := 0; i < len(g); i++ {
		g[i] = &Node{}
	}

	return g
}

func (g Graph) AddEdge(from, to int) {
	if from >= len(g) || to >= len(g) {
		panic("invalid index")
	}
	g[from].Edges[g[to]] = true
	g[to].Edges[g[from]] = true
}
