package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	for i, tc := range []struct {
		n             int
		adj           [][]int
		s             int
		t             int
		expectedPath  []int
		expectedFound bool
	}{
		{
			n: 3,
			adj: [][]int{
				{1, 2},
				{0, 2},
				{1, 0},
			},
			s:             0,
			t:             2,
			expectedPath:  []int{0, 2},
			expectedFound: true,
		},
		{
			n: 3,
			adj: [][]int{
				{1, 2},
				{0},
				{0},
			},
			s:             1,
			t:             2,
			expectedPath:  []int{1, 0, 2},
			expectedFound: true,
		},
		{
			n: 5,
			adj: [][]int{
				{3, 4},
				{0},
				{0, 1},
				{4},
				{2},
			},
			s:             2,
			t:             4,
			expectedPath:  []int{2, 0, 4},
			expectedFound: true,
		},
		{
			n: 5,
			adj: [][]int{
				{3, 4},
				{0},
				{0, 1},
				{4},
				{2},
			},
			s:             0,
			t:             1,
			expectedPath:  []int{0, 4, 2, 1},
			expectedFound: true,
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			found, path := BFS(tc.n, tc.adj, tc.s, tc.t)
			if found != tc.expectedFound {
				t.Errorf("expected result %v, have %v", tc.expectedFound, found)
			}
			if !reflect.DeepEqual(path, tc.expectedPath) {
				t.Errorf("expected path %v, have %v", tc.expectedPath, path)
			}
		})
	}
}
