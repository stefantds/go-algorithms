package heap

import "fmt"

type indexedHeap struct {
	keys []int
	pq   []int
	qp   []int
	n    int
	maxN int
}

// NewIndexedHeap returns a new PQ initialized an empty indexed priority
// queue with indices between 0 and maxN-1
func NewIndexedHeap(maxN int) *indexedHeap {
	h := &indexedHeap{
		keys: make([]int, maxN+1),
		pq:   make([]int, maxN+1),
		qp:   make([]int, maxN+1),
		maxN: maxN,
	}

	for i := 0; i <= maxN; i++ {
		h.qp[i] = -1
	}

	return h
}

func (h *indexedHeap) Len() int {
	return h.n
}

func (h *indexedHeap) IsEmpty() bool {
	return h.n == 0
}

// Min returns the key and the value corresponding to the
// maximum value currently in the PQ
func (h *indexedHeap) Min() (idx, key int) {
	if h.n == 0 {
		panic("called Min on an empty heap")
	}

	return h.pq[1], h.keys[h.pq[1]]
}

func (h *indexedHeap) DelMin() (idx, key int) {
	if h.n == 0 {
		panic("called DelMin on an empty heap")
	}

	min := h.pq[1]
	minKey := h.keys[min]

	// move the current max to the end of the heap
	h.swap(1, h.n-1)
	h.n--

	// sink the element from pos 1, as it is not at its place
	h.sink(1)

	h.qp[min] = -1

	return min, minKey
}

func (h *indexedHeap) Insert(idx, key int) {
	h.validateIdx(idx)

	if h.Contains(idx) {
		panic(fmt.Sprintf("index %v already exists", idx))
	}

	h.n++
	h.qp[idx] = h.n
	h.pq[h.n] = idx
	h.keys[idx] = key

	// move the new element to its correct place
	h.swim(h.n)
}

func (h *indexedHeap) Contains(idx int) bool {
	h.validateIdx(idx)

	return h.qp[idx] != -1
}

func (h *indexedHeap) ChangeKey(idx, key int) {
	h.validateIdx(idx)
	if !h.Contains(idx) {
		panic(fmt.Sprintf("index %v not in the PQ", idx))
	}

	h.keys[idx] = key
	h.swim(h.qp[idx])
	h.sink(h.qp[idx])
}

// swim takes the element at pos and moves it up the heap until
// its place (until smaller than its parent)
func (h *indexedHeap) swim(pos int) {
	for pos > 1 && h.greater(pos/2, pos) {
		h.swap(pos, pos/2)
		pos = pos / 2
	}
}

// swim takes the element at pos and moves it down the heap until
// its place (until greater than its children)
func (h *indexedHeap) sink(pos int) {
	for 2*pos <= h.n {
		j := 2 * pos
		if j < h.n && h.greater(j, j+1) {
			j += 1
		}
		if !h.greater(pos, j) {
			break
		}

		h.swap(pos, j)
		pos = j
	}
}

func (h *indexedHeap) greater(i, j int) bool {
	return h.keys[h.pq[i]] > h.keys[h.pq[j]]
}

func (h *indexedHeap) swap(i, j int) {
	h.pq[i], h.pq[j] = h.pq[j], h.pq[i]
	h.qp[h.pq[i]] = i
	h.qp[h.pq[j]] = j
}

func (h *indexedHeap) validateIdx(idx int) {
	if idx < 0 || idx >= h.maxN {
		panic(fmt.Sprintf("invalid index %v", idx))
	}
}
