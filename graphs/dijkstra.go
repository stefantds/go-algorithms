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

	relax := func(e edge) {
		if distTo[e.from]+e.weight < distTo[e.to] {
			distTo[e.to] = distTo[e.from] + e.weight
			pathTo[e.to] = e.from
			if pq.Contains(e.to) {
				pq.ChangeKey(e.to, distTo[e.to])
			} else {
				pq.Insert(e.to, distTo[e.to])
			}
		}
	}

	// initialize
	for v := 0; v < n; v++ {
		distTo[v] = math.MaxInt32
	}
	distTo[source] = 0
	pathTo[source] = source
	pq.Insert(source, 0)

	for !pq.IsEmpty() {
		curr, _ := pq.DelMin()
		for _, e := range adj[curr] {
			relax(e)
		}
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

type edgeHeap []edge

func (h edgeHeap) Len() int            { return len(h) }
func (h edgeHeap) Less(i, j int) bool  { return h[i].weight < h[j].weight }
func (h *edgeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *edgeHeap) Push(v interface{}) { (*h) = append(*h, v.(edge)) }
func (h *edgeHeap) Pop() interface{} {
	var e edge
	*h, e = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return e
}

type edgesPQ struct {
	h edgeHeap
}

func NewEdgesPQ() *edgesPQ {
	return &edgesPQ{
		h: make(edgeHeap, 0),
	}
}

func (h *edgesPQ) Init(edges []edge) {
	*h = edgesPQ{
		h: edges,
	}
	heap.Init(&h.h)
}

func (h *edgesPQ) Push(v edge) {
	heap.Push(&h.h, v)
}

func (h *edgesPQ) Pop() edge {
	return heap.Pop(&h.h).(edge)
}

func (h *edgesPQ) IsEmpty() bool {
	return len(h.h) == 0
}
