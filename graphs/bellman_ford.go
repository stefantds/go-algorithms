package graphs

import (
	"math"
)

func BellmanFord(n int, edges []edge, source int) (map[int]int, map[int]int, bool) {
	distTo := make(map[int]int)
	pathTo := make(map[int]int)

	relax := func(e edge) bool {
		if distTo[e.from] != math.MaxInt32 &&
			distTo[e.from]+e.weight < distTo[e.to] {
			distTo[e.to] = distTo[e.from] + e.weight
			pathTo[e.to] = e.from
			return true
		}
		return false
	}

	// initialize
	for v := 0; v < n; v++ {
		distTo[v] = math.MaxInt32
	}
	distTo[source] = 0
	pathTo[source] = source

	// iterate V-1 times;
	// alternatively iterate V times if we need t detect negative cycles. Check
	// if in the last iteration any relaxation was done
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			relax(e)
		}
	}

	// alternatively one can do another pass over all edges to detect negative cycle. If any further updates
	// is found, it means there is a negative-weight cycle
	hasNegativeCycles := false
	for _, e := range edges {
		if relax(e) {
			hasNegativeCycles = true
		}
	}

	return distTo, pathTo, hasNegativeCycles
}
