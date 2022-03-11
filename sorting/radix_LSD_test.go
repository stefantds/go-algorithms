package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	input          []string
	expectedOutput []string
}

func TestSortRadixLSD(t *testing.T) {
	for i, tc := range []testCase{
		{
			input:          []string{},
			expectedOutput: []string{},
		},
		{
			input:          []string{"a"},
			expectedOutput: []string{"a"},
		},
		{
			input:          []string{"a", "b"},
			expectedOutput: []string{"a", "b"},
		},
		{
			input:          []string{"b", "a"},
			expectedOutput: []string{"a", "b"},
		},
		{
			input:          []string{"baa", "aba", "aab"},
			expectedOutput: []string{"aab", "aba", "baa"},
		},
	} {
		tc := tc
		t.Run(fmt.Sprintf("test case %v", i), func(t *testing.T) {
			result := SortRadixLSD(tc.input)
			if !reflect.DeepEqual(result, tc.expectedOutput) {
				t.Errorf("expected %v, have %v", tc.expectedOutput, result)
			}
		})
	}
}
