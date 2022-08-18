package graphs

import (
	"sort"

	"github.com/stefantds/go-algorithms/union_find"
)

type edge struct {
	from, to int
	weight   int
}

func Kruskal(n int, edges []edge) []edge {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	uf := union_find.NewWeightedQuickUnion(n)

	result := make([]edge, 0)

	i := 0
	// stop when having enough edges in the MST or when there are no more edges
	for i < len(edges) && len(result) < n-1 {
		curr := edges[i]
		if !uf.Connected(curr.from, curr.to) {
			result = append(result, curr)
			uf.Union(curr.from, curr.to)
		}
		i++
	}

	return result
}
