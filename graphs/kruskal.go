package graphs

import (
	"container/heap"

	"github.com/stefantds/go-algorithms/union_find"
)

type edge struct {
	from, to int
	weight   int
}

func Kruskal(n int, edges []edge) []edge {
	edgesPQ := edgeHeap(edges)
	heap.Init(&edgesPQ)

	uf := union_find.NewWeightedQuickUnion(n)

	result := make([]edge, 0)

	// stop when having enough edges in the MST or when there are no more edges
	for len(edgesPQ) > 0 && len(result) < n-1 {
		curr := heap.Pop(&edgesPQ).(edge)
		if !uf.Connected(curr.from, curr.to) {
			result = append(result, curr)
			uf.Union(curr.from, curr.to)
		}
	}

	return result
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
