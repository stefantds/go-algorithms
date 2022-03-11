package union_find

import (
	"testing"

	"github.com/stefantds/go-algorithms/union_find/testdata"
)

func TestQuickFind(t *testing.T) {
	data := testdata.TinyUF
	uf := NewQuickFind(data.N)
	for i, step := range data.Ops {
		switch step.Op {
		case testdata.Union:
			uf.Union(step.Args[0], step.Args[1])
		case testdata.Connected:
			c := uf.Connected(step.Args[0], step.Args[1])
			if c != step.ExpectedResult {
				t.Errorf("wrong result at step %v (%v): expected result %v, got %v", i, step.Op, step.ExpectedResult, c)
			}
		}
	}
}
