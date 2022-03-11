package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

type ShortestPathTestCase struct {
	GraphSize      int
	Edges          map[int][]edge
	Source         int
	ExpectedDistTo map[int]int
	ExpectedPathTo map[int]int
}

func TestDijkstra(t *testing.T) {
	for i, tc := range []ShortestPathTestCase{
		{
			GraphSize: 3,
			Edges: map[int][]edge{
				0: {
					{from: 0, to: 1, weight: 2},
					{from: 0, to: 2, weight: 5},
				},
				1: {
					{from: 1, to: 0, weight: 3},
					{from: 1, to: 2, weight: 2},
				},
				2: {
					{from: 2, to: 1, weight: 2},
					{from: 2, to: 0, weight: 5},
				},
			},
			Source: 0,
			ExpectedDistTo: map[int]int{
				0: 0,
				1: 2,
				2: 4,
			},
			ExpectedPathTo: map[int]int{
				0: 0,
				1: 0,
				2: 1,
			},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			distTo, pathTo := Dijkstra(tc.GraphSize, tc.Edges, tc.Source)
			if !reflect.DeepEqual(distTo, tc.ExpectedDistTo) {
				t.Errorf("expected distances %v, have %v", tc.ExpectedDistTo, distTo)
			}
			if !reflect.DeepEqual(pathTo, tc.ExpectedPathTo) {
				t.Errorf("expected path %v, have %v", tc.ExpectedPathTo, pathTo)
			}
		})
	}
}
