package heap

import (
	"testing"

	"github.com/stefantds/go-algorithms/heap/testdata"
)

func TestHeap(t *testing.T) {
	data := testdata.TinyHeapData
	pq := NewHeap()
	for i, step := range data.Ops {
		switch step.Op {
		case testdata.Insert:
			pq.Insert(step.Arg)
		case testdata.Max:
			max := pq.Max()
			if max != step.ExpectedResult {
				t.Errorf("wrong result at step %v (%v): expected result %v, got %v", i, step.Op, step.ExpectedResult, max)
			}
		case testdata.DelMax:
			max := pq.DelMax()
			if max != step.ExpectedResult {
				t.Errorf("wrong result at step %v (%v): expected result %v, got %v", i, step.Op, step.ExpectedResult, max)
			}
		}
	}
}
