package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

type DFSTestCase struct {
	GraphSize         int
	Edges             [][]edge
	Source            int
	ExpectedPostorder []int
	ExpectedPreorder  []int
}

func TestDFSWithCollectionOrder(t *testing.T) {
	for i, tc := range []DFSTestCase{
		{
			GraphSize: 3,
			Edges: [][]edge{
				{
					{from: 0, to: 1},
					{from: 0, to: 2},
				},
				{
					{from: 1, to: 0},
					{from: 1, to: 2},
				},
				{
					{from: 2, to: 1},
					{from: 2, to: 0},
				},
			},
			Source:            0,
			ExpectedPreorder:  []int{0, 1, 2},
			ExpectedPostorder: []int{2, 1, 0},
		},
		{
			GraphSize: 3,
			Edges: [][]edge{
				{
					{from: 0, to: 1},
					{from: 0, to: 2},
				},
				{},
				{},
			},
			Source:            0,
			ExpectedPreorder:  []int{0, 1, 2},
			ExpectedPostorder: []int{1, 2, 0},
		},
		{
			GraphSize: 5,
			Edges: [][]edge{
				{{from: 0, to: 3}, {from: 0, to: 4}},
				{{from: 1, to: 0}},
				{{from: 2, to: 0}, {from: 2, to: 1}},
				{{from: 3, to: 4}},
				{{from: 4, to: 2}},
			},
			Source:            2,
			ExpectedPreorder:  []int{2, 0, 3, 4, 1},
			ExpectedPostorder: []int{4, 3, 0, 1, 2},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			pre, post := DFSWithCollectionOrder(tc.GraphSize, tc.Source, tc.Edges)
			if !reflect.DeepEqual(pre, tc.ExpectedPreorder) {
				t.Errorf("expected preorder %v, have %v", tc.ExpectedPreorder, pre)
			}
			if !reflect.DeepEqual(post, tc.ExpectedPostorder) {
				t.Errorf("expected postorder %v, have %v", tc.ExpectedPostorder, post)
			}
		})
	}
}
