package graphs

import (
	"container/heap"
	"math"
)

type indexedPQ interface {
	DelMin() (k, idx int)
	Insert(idx, key int)
	Contains(idx int) bool
	ChangeKey(idx, key int)
	IsEmpty() bool
}

func Dijkstra(n int, adj map[int][]edge, source int, pq indexedPQ) (map[int]int, map[int]int) {
	distTo := make(map[int]int)
	pathTo := make(map[int]int)

	relax := func(v int) {
		for _, e := range adj[v] {
			if distTo[v]+e.weight < distTo[e.to] {
				if pq.Contains(e.to) {
					pq.ChangeKey(e.to, distTo[v]+e.weight)
				} else {
					pq.Insert(e.to, distTo[v]+e.weight)
				}
				distTo[e.to] = distTo[v] + e.weight
				pathTo[e.to] = v
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

	for !pq.IsEmpty() {
		curr, _ := pq.DelMin()
		relax(curr)
	}

	return distTo, pathTo
}

// DijkstraLazy is a lazy implementation of Dijkstra, using a
// normal PQ, not an indexed one. As a result, multiple edges leading to
// the same vertex can be possible in the PQ at the same time.
// The solution is to keep a set of visited vertexes and ignore any edge
// coming from the PW that has already been visited
func DijkstraLazy(n int, adj map[int][]edge, source int) (map[int]int, map[int]int) {
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
