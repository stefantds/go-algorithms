package sorting

func Heapsort(a []int) {
	n := len(a)

	// heapify phase
	for k := n / 2; k >= 1; k-- {
		sink(a, k, n)
	}

	// repeatedly find the max and move it to the end
	k := n
	for k > 1 {
		swap1Indexed(a, 1, k) // move the max from a[1:k+1] to pos k
		k--
		sink(a, 1, k) // fix heap but with size decreased by 1
	}
}

// sink moves the element at index k in a
// to its correct heap position.
// k is 1-indexed
func sink(a []int, k int, n int) {
	for 2*k <= n {
		j := 2 * k // the first child of a[k]
		if j < n && less1Indexed(a, j, j+1) {
			j = j + 1 // the second child of a[k] is greater
		}
		if !less1Indexed(a, k, j) { // a[k] is in the right place, as both the children are greater or equal
			break
		}
		swap1Indexed(a, k, j) // swap a[k] with its greater child
		k = j
	}
}

func less1Indexed(a []int, i, j int) bool {
	return a[i-1] < a[j-1]
}

func swap1Indexed(a []int, i, j int) {
	a[i-1], a[j-1] = a[j-1], a[i-1]
}
