package heap

type intHeap []int

func NewHeap() *intHeap {
	h := make(intHeap, 1)
	return &h
}

func (h *intHeap) Max() int {
	if len(*h) <= 1 {
		panic("called Max on an empty heap")
	}

	return (*h)[1]
}

func (h *intHeap) Len() int {
	return len(*h) - 1
}

func (h *intHeap) DelMax() int {
	if len(*h) <= 1 {
		panic("called DelMax on an empty heap")
	}

	max := (*h)[1]

	// move the current max to the end of the heap
	h.swap(1, len(*h)-1)

	// exclude the current max from the heap by decreasing the heap size
	(*h) = (*h)[:len(*h)-1]

	// sink the element from pos 1, as it is not at its place
	h.sink(1)

	return max
}

func (h *intHeap) Insert(val int) {
	(*h) = append(*h, val)

	// move the new element to its correct place
	h.swim(len(*h) - 1)
}

// swim takes the element at pos and moves it up the heap until
// its place (until smaller than its parent)
func (h *intHeap) swim(pos int) {
	for i := pos; i > 1 && !h.less(i, i/2); i = i / 2 {
		h.swap(i, i/2)
	}
}

// swim takes the element at pos and moves it down the heap until
// its place (until greater than its children)
func (h *intHeap) sink(pos int) {
	i := pos
	for 2*i < len(*h) {
		j := 2 * i
		if j < len(*h)-1 && h.less(j, j+1) {
			j += 1
		}
		if !h.less(i, j) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

func (h *intHeap) less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *intHeap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
