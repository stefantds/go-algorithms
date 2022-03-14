package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrim(t *testing.T) {
	for i, tc := range []struct {
		graphSize      int
		edges          map[int][]edge
		expectedResult []edge
	}{
		{
			graphSize: 3,
			edges: map[int][]edge{
				0: {
					{from: 0, to: 1, weight: 3},
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
			expectedResult: []edge{
				{from: 0, to: 1, weight: 3},
				{from: 1, to: 2, weight: 2},
			},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			result := Prim(tc.graphSize, tc.edges, NewEdgesPQ())
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("expected %v, have %v", tc.expectedResult, result)
			}
		})
	}
}
