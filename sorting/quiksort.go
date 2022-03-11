package sorting

import "math/rand"

func Quicksort(s []int) {
	shuffle(s)
	sort(s, 0, len(s)-1)
}

func sort(s []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotS, pivotE := partition(s, lo, hi)
	sort(s, lo, pivotS-1)
	sort(s, pivotE+1, hi)
}

// partition uses Dijkstra 3-way partitioning
func partition(s []int, lo, hi int) (pivotStart, pivotEnd int) {
	pivot := s[lo]
	lt, gt := lo, hi
	i := lo
	for i <= gt {
		if s[i] < pivot {
			swap(s, lt, i)
			lt++
			i++
		} else if s[i] > pivot {
			swap(s, gt, i)
			gt--
		} else {
			i++
		}
	}

	return lt, gt
}

func shuffle(s []int) {
	for i := 0; i < len(s); i++ {
		pos := i + rand.Intn(len(s)-i)
		swap(s, i, pos)
	}
}

func swap(s []int, i, j int) {
	s[i], s[j] = s[j], s[i]
}
