package sorting

func Mergesort(a []int) {
	msort(a, make([]int, len(a)), 0, len(a)-1)
}

func msort(a, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}

	mid := lo + (hi-lo)/2
	msort(a, aux, lo, mid)
	msort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func merge(a []int, aux []int, lo, mid, hi int) {
	for i := lo; i <= hi; i++ {
		aux[i] = a[i]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}
