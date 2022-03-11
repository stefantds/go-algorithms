package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

type BFTestCase struct {
	GraphSize                int
	Edges                    []edge
	Source                   int
	ExpectedDistTo           map[int]int
	ExpectedPathTo           map[int]int
	ExpectedHasNegativeCycle bool
}

func TestBellmanFord(t *testing.T) {
	for i, tc := range []BFTestCase{
		{
			GraphSize: 3,
			Edges: []edge{
				{from: 0, to: 1, weight: 2},
				{from: 0, to: 2, weight: 5},
				{from: 1, to: 0, weight: 3},
				{from: 1, to: 2, weight: 2},
				{from: 2, to: 1, weight: 2},
				{from: 2, to: 0, weight: 5},
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
			ExpectedHasNegativeCycle: false,
		},
		{
			GraphSize: 5,
			Edges: []edge{

				{from: 0, to: 3, weight: 1},
				{from: 0, to: 4, weight: 2},
				{from: 1, to: 0, weight: 1},
				{from: 2, to: 0, weight: 1},
				{from: 2, to: 1, weight: -1},
				{from: 2, to: 4, weight: 4},
				{from: 3, to: 4, weight: 2},
			},
			Source: 2,
			ExpectedDistTo: map[int]int{
				0: 0,
				1: -1,
				2: 0,
				3: 1,
				4: 2,
			},
			ExpectedPathTo: map[int]int{
				0: 1,
				1: 2,
				2: 2,
				3: 0,
				4: 0,
			},
			ExpectedHasNegativeCycle: false,
		},
		{
			GraphSize: 5,
			Edges: []edge{

				{from: 0, to: 3, weight: 1},
				{from: 0, to: 4, weight: 2},
				{from: 1, to: 0, weight: 1},
				{from: 2, to: 0, weight: 1},
				{from: 2, to: 1, weight: -1},
				{from: 3, to: 4, weight: 2},
				{from: 4, to: 2, weight: -4},
			},
			Source:                   2,
			ExpectedHasNegativeCycle: true,
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			distTo, pathTo, hasNegCycle := BellmanFord(tc.GraphSize, tc.Edges, tc.Source)
			if tc.ExpectedDistTo != nil && !reflect.DeepEqual(distTo, tc.ExpectedDistTo) {
				t.Errorf("expected distances %v, have %v", tc.ExpectedDistTo, distTo)
			}
			if tc.ExpectedPathTo != nil && !reflect.DeepEqual(pathTo, tc.ExpectedPathTo) {
				t.Errorf("expected path %v, have %v", tc.ExpectedPathTo, pathTo)
			}
			if !reflect.DeepEqual(hasNegCycle, tc.ExpectedHasNegativeCycle) {
				t.Errorf("expected negative cycles: %v, have: %v", tc.ExpectedHasNegativeCycle, hasNegCycle)
			}
		})
	}
}
