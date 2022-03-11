package graphs

type PathsFromSource struct {
	parentOf map[*Node]*Node
	visited  map[*Node]bool
}

func NewPaths(g Graph, source int) PathsFromSource {
	if source >= len(g) {
		panic("invalid source vertex")
	}
	p := PathsFromSource{
		parentOf: make(map[*Node]*Node),
		visited:  make(map[*Node]bool),
	}

	p.dfs(g[source])

	return p
}

func (p *PathsFromSource) dfs(v *Node) {
	if v == nil {
		return
	}
	if !p.visited[v] {
		p.visited[v] = true
		for n := range v.Edges {
			p.dfs(n)
			p.parentOf[n] = v
		}
	}
}

func DFSWithCollectionOrder(n int, source int, edges [][]edge) (pre, post []int) {
	var doDFS func(int)
	marked := make([]bool, n)
	preorder := make([]int, 0, n)
	postorder := make([]int, 0, n)

	doDFS = func(v int) {
		marked[v] = true
		preorder = append(preorder, v) // preorder: add before continuing to neighbors

		for _, w := range edges[v] {
			if !marked[w.to] {
				doDFS(w.to)
			}
		}

		postorder = append(postorder, v) // postorder: add after continuing to neighbors
	}

	doDFS(source)

	return preorder, postorder
}
