package heap

type PriorityQueue interface {
	Max() int
	Insert(element int)
	DelMax() int
	Len() int
}
