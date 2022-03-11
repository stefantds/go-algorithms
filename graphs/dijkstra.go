package graphs

import (
	"container/heap"
	"math"
)

func Dijkstra(n int, adj map[int][]edge, source int) (map[int]int, map[int]int) {
	pq := make(edgeHeap, 0)
	heap.Init(&pq)

	// marked is needed for the lazy implementation since the
	// sam vertex will be added multiple times to the pq
	marked := make(map[int]bool)
	distTo := make(map[int]int)
	pathTo := make(map[int]int)

	relax := func(v int) {
		marked[v] = true
		for _, n := range adj[v] {
			if distTo[v]+n.weight < distTo[n.to] {
				distTo[n.to] = distTo[v] + n.weight
				pathTo[n.to] = v
				heap.Push(&pq, n)
			}
		}
	}

	// initialize
	for v := 0; v < n; v++ {
		distTo[v] = math.MaxInt32
	}
	distTo[source] = 0
	pathTo[source] = source

	relax(source)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(edge)
		if marked[curr.to] {
			continue
		}
		relax(curr.to)
	}

	return distTo, pathTo
}
