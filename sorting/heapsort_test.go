package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapsort(t *testing.T) {
	for i, tc := range []struct {
		input          []int
		expectedOutput []int
	}{
		{
			input:          []int{},
			expectedOutput: []int{},
		},
		{
			input:          []int{1},
			expectedOutput: []int{1},
		},
		{
			input:          []int{3, 2, 1},
			expectedOutput: []int{1, 2, 3},
		},
		{
			input:          []int{1, 2, 3},
			expectedOutput: []int{1, 2, 3},
		},
		{
			input:          []int{1, 2, 2, 2, 1},
			expectedOutput: []int{1, 1, 2, 2, 2},
		},
		{
			input:          []int{1, 1, 1, 1, 1},
			expectedOutput: []int{1, 1, 1, 1, 1},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			Heapsort(tc.input)
			if !reflect.DeepEqual(tc.input, tc.expectedOutput) {
				t.Errorf("expected %v, have %v", tc.expectedOutput, tc.input)
			}
		})
	}
}
