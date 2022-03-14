package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

type MSTTestCase struct {
	graphSize      int
	edges          []edge
	expectedResult []edge
}

func TestKruskal(t *testing.T) {
	for i, tc := range []MSTTestCase{
		{
			graphSize: 3,
			edges: []edge{
				{from: 0, to: 1, weight: 3},
				{from: 0, to: 2, weight: 5},
				{from: 1, to: 2, weight: 2},
			},
			expectedResult: []edge{
				{from: 1, to: 2, weight: 2},
				{from: 0, to: 1, weight: 3},
			},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			result := Kruskal(tc.graphSize, tc.edges, NewEdgesPQ())
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("expected %v, have %v", tc.expectedResult, result)
			}
		})
	}
}
