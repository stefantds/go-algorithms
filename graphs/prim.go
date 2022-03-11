package graphs

import (
	"container/heap"
)

func Prim(n int, adj map[int][]edge) []edge {
	pq := make(edgeHeap, 0)
	heap.Init(&pq)
	marked := make(map[int]bool)
	result := make([]edge, 0)

	// start with any vertex
	visit(adj, 0, marked, &pq)

	for len(pq) > 0 && len(result) < n-1 {
		curr := heap.Pop(&pq).(edge) // take the edge with the min weight

		// ignore edges that have no new vertex (lazy implementation)
		if marked[curr.from] && marked[curr.to] {
			continue
		}

		// if not ignored, it is in the result
		result = append(result, curr)

		// visit the vertex that was not visited before
		if !marked[curr.from] {
			visit(adj, curr.from, marked, &pq)
		} else {
			visit(adj, curr.to, marked, &pq)
		}
	}

	return result
}

func visit(adj map[int][]edge, v int, marked map[int]bool, pq *edgeHeap) {
	marked[v] = true

	// go through all neighbors of v and add them to the heap
	for _, e := range adj[v] {
		// since one end of the edge is v (which is visited), we ignore any edge that
		// has thre other side visited as well because it would create a cycle
		if !marked[e.to] {
			heap.Push(pq, e)
		}
	}
}
