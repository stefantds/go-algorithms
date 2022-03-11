package graphs

// BidirectionalBFS uses the bidirectional BFS algorithm to check if it is possible to
// reach the sink from the source.
// it returns a bool and the path from source to sink if such path exist.
func BidirectionalBFS(n int, edges [][]int, s, t int) (bool, []int) {
	sPathTo, tPathTo := make(map[int]int), make(map[int]int)
	sPathTo[s] = s
	sPathTo[t] = t

	sVisited, tVisited := make(map[int]bool), make(map[int]bool)
	sVisited[s] = true
	tVisited[t] = true

	sQueue, tQueue := make([]int, 0, n), make([]int, 0, n)
	sQueue = append(sQueue, s)
	tQueue = append(tQueue, t)

	targetReached := false
	var meetPoint int
	for len(sQueue) > 0 && len(tQueue) > 0 && !targetReached {
		newSQueue, newTQueue := make([]int, 0, n), make([]int, 0, n)
		for len(sQueue) > 0 {
			var next int
			next, sQueue = (sQueue)[0], (sQueue)[1:]
			for _, neighbor := range edges[next] {
				if !sVisited[neighbor] {
					sVisited[neighbor] = true
					sPathTo[neighbor] = next
					newSQueue = append(newSQueue, neighbor)
					if tVisited[neighbor] {
						targetReached = true
						meetPoint = neighbor
					}
				}
			}
		}

		for len(tQueue) > 0 {
			var next int
			next, tQueue = (tQueue)[0], (tQueue)[1:]
			for _, neighbor := range edges[next] {
				if !tVisited[neighbor] {
					tVisited[neighbor] = true
					tPathTo[neighbor] = next
					newTQueue = append(newTQueue, neighbor)
					if sVisited[neighbor] {
						targetReached = true
						meetPoint = neighbor
					}
				}
			}
		}

		sQueue, tQueue = newSQueue, newTQueue
	}

	var pathFromSToT []int
	if targetReached {
		firstPart := make([]int, 0)
		curr := meetPoint
		for curr != s {
			firstPart = append(firstPart, curr)
			curr = sPathTo[curr]
		}

		secondPart := make([]int, 0)
		curr = meetPoint
		for curr != t {
			secondPart = append(secondPart, curr)
			curr = tPathTo[curr]
		}

		pathFromSToT = append(
			append([]int{s}, reverse(firstPart)...),
			append(secondPart, t)...,
		)
	}

	return targetReached, pathFromSToT
}

func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	return s
}
