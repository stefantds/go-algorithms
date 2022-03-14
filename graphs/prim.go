package graphs

type priorityQueue interface {
	IsEmpty() bool
	Push(x edge)
	Pop() edge
	Init([]edge)
}

func Prim(n int, adj map[int][]edge, pq priorityQueue) []edge {
	marked := make(map[int]bool)
	result := make([]edge, 0)

	// start with any vertex
	visit(adj, 0, marked, pq)

	for !pq.IsEmpty() && len(result) < n-1 {
		curr := pq.Pop() // take the edge with the min weight

		// ignore edges that have no new vertex (lazy implementation)
		if marked[curr.from] && marked[curr.to] {
			continue
		}

		// if not ignored, it is in the result
		result = append(result, curr)

		// visit the vertex that was not visited before
		if !marked[curr.from] {
			visit(adj, curr.from, marked, pq)
		} else {
			visit(adj, curr.to, marked, pq)
		}
	}

	return result
}

func visit(adj map[int][]edge, v int, marked map[int]bool, pq priorityQueue) {
	marked[v] = true

	// go through all neighbors of v and add them to the heap
	for _, e := range adj[v] {
		// since one end of the edge is v (which is visited), we ignore any edge that
		// has thre other side visited as well because it would create a cycle
		if !marked[e.to] {
			pq.Push(e)
		}
	}
}
