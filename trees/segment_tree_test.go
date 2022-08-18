package trees

import (
	"fmt"
	"testing"
)

type SementTreeTestCase struct {
	values          []int
	queries         [][]int
	expectedResults []int
}

func TestSementTree(t *testing.T) {
	for i, tc := range []SementTreeTestCase{
		{
			values:          []int{6, 10, 5, 2, 7, 1, 0, 9},
			queries:         [][]int{{0, 5}, {0, 8}, {2, 8}, {5, 6}},
			expectedResults: []int{10, 10, 9, 1},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			st := NewSegmentTree(tc.values)
			for i, q := range tc.queries {
				res := st.RangeQuery(q[0], q[1])
				if res != tc.expectedResults[i] {
					t.Errorf("expected %v, got %v", tc.expectedResults[i], res)
				}
			}
		})
	}
}
